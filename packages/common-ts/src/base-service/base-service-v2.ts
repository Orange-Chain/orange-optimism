import { Server } from 'net'

import Config from 'bcfg'
import * as dotenv from 'dotenv'
import { Command, Option } from 'commander'
import { ValidatorSpec, Spec, cleanEnv } from 'envalid'
import snakeCase from 'lodash/snakeCase'
import express, { Router } from 'express'
import prometheus, { Registry } from 'prom-client'
import promBundle from 'express-prom-bundle'
import bodyParser from 'body-parser'
import morgan from 'morgan'

import { Logger } from '../common/logger'
import { Metric, Gauge, Counter } from './metrics'
import { validators } from './validators'

export type Options = {
  [key: string]: any
}

export type StandardOptions = {
  loopIntervalMs?: number
  port?: number
  hostname?: string
}

export type OptionsSpec<TOptions extends Options> = {
  [P in keyof Required<TOptions>]: {
    validator: (spec?: Spec<TOptions[P]>) => ValidatorSpec<TOptions[P]>
    desc: string
    default?: TOptions[P]
    secret?: boolean
  }
}

export type MetricsV2 = Record<any, Metric>

export type StandardMetrics = {
  metadata: Gauge
  unhandledErrors: Counter
}

export type MetricsSpec<TMetrics extends MetricsV2> = {
  [P in keyof Required<TMetrics>]: {
    type: new (configuration: any) => TMetrics[P]
    desc: string
    labels?: string[]
  }
}

export type ExpressRouter = Router

/**
 * BaseServiceV2 is an advanced but simple base class for long-running TypeScript services.
 */
export abstract class BaseServiceV2<
  TOptions extends Options,
  TMetrics extends MetricsV2,
  TServiceState
> {
  /**
   * The timeout that controls the polling interval
   * If clearTimeout(this.pollingTimeout) is called the timeout will stop
   */
  private pollingTimeout: NodeJS.Timeout

  /**
   * The promise representing this.main
   */
  private mainPromise: ReturnType<typeof this.main>

  /**
   * Whether or not the service will loop.
   */
  protected loop: boolean

  /**
   * Waiting period in ms between loops, if the service will loop.
   */
  protected loopIntervalMs: number

  /**
   * Whether or not the service is currently running.
   */
  protected running: boolean

  /**
   * Whether or not the service is currently healthy.
   */
  protected healthy: boolean

  /**
   * Logger class for this service.
   */
  protected logger: Logger

  /**
   * Service state, persisted between loops.
   */
  protected state: TServiceState

  /**
   * Service options.
   */
  protected readonly options: TOptions & StandardOptions

  /**
   * Metrics.
   */
  protected readonly metrics: TMetrics & StandardMetrics

  /**
   * Registry for prometheus metrics.
   */
  protected readonly metricsRegistry: Registry

  /**
   * App server.
   */
  protected server: Server

  /**
   * Port for the app server.
   */
  protected readonly port: number

  /**
   * Hostname for the app server.
   */
  protected readonly hostname: string

  /**
   * @param params Options for the construction of the service.
   * @param params.name Name for the service. This name will determine the prefix used for logging,
   * metrics, and loading environment variables.
   * @param params.optionsSpec Settings for input options. You must specify at least a
   * description for each option.
   * @param params.metricsSpec Settings that define which metrics are collected. All metrics that
   * you plan to collect must be defined within this object.
   * @param params.options Options to pass to the service.
   * @param params.loops Whether or not the service should loop. Defaults to true.
   * @param params.loopIntervalMs Loop interval in milliseconds. Defaults to zero.
   * @param params.port Port for the app server. Defaults to 7300.
   * @param params.hostname Hostname for the app server. Defaults to 0.0.0.0.
   */
  constructor(
    private readonly params: {
      name: string
      version: string
      optionsSpec: OptionsSpec<TOptions>
      metricsSpec: MetricsSpec<TMetrics>
      options?: Partial<TOptions>
      loop?: boolean
      loopIntervalMs?: number
      port?: number
      hostname?: string
    }
  ) {
    this.loop = params.loop !== undefined ? params.loop : true
    this.state = {} as TServiceState

    const stdOptionsSpec: OptionsSpec<StandardOptions> = {
      loopIntervalMs: {
        validator: validators.num,
        desc: 'Loop interval in milliseconds',
        default: params.loopIntervalMs || 0,
      },
      port: {
        validator: validators.num,
        desc: 'Port for the app server',
        default: params.port || 7300,
      },
      hostname: {
        validator: validators.str,
        desc: 'Hostname for the app server',
        default: params.hostname || '0.0.0.0',
      },
    }

    // Add default options to options spec.
    ;(params.optionsSpec as any) = {
      ...(params.optionsSpec || {}),
      ...stdOptionsSpec,
    }

    // List of options that can safely be logged.
    const publicOptionNames = Object.entries(params.optionsSpec)
      .filter(([, spec]) => {
        return spec.secret !== true
      })
      .map(([key]) => {
        return key
      })

    // Add default metrics to metrics spec.
    ;(params.metricsSpec as any) = {
      ...(params.metricsSpec || {}),

      // Users cannot set these options.
      metadata: {
        type: Gauge,
        desc: 'Service metadata',
        labels: ['name', 'version'].concat(publicOptionNames),
      },
      unhandledErrors: {
        type: Counter,
        desc: 'Unhandled errors',
      },
    }

    /**
     * Special snake_case function which accounts for the common strings "L1" and "L2" which would
     * normally be split into "L_1" and "L_2" by the snake_case function.
     *
     * @param str String to convert to snake_case.
     * @returns snake_case string.
     */
    const opSnakeCase = (str: string) => {
      const reg = /l_1|l_2/g
      const repl = str.includes('l1') ? 'l1' : 'l2'
      return snakeCase(str).replace(reg, repl)
    }

    // Use commander as a way to communicate info about the service. We don't actually *use*
    // commander for anything besides the ability to run `ts-node ./service.ts --help`.
    const program = new Command()
    for (const [optionName, optionSpec] of Object.entries(params.optionsSpec)) {
      program.addOption(
        new Option(`--${optionName.toLowerCase()}`, `${optionSpec.desc}`).env(
          `${opSnakeCase(
            params.name.replace(/-/g, '_')
          ).toUpperCase()}__${opSnakeCase(optionName).toUpperCase()}`
        )
      )
    }

    const longestMetricNameLength = Object.keys(params.metricsSpec).reduce(
      (acc, key) => {
        const nameLength = snakeCase(key).length
        if (nameLength > acc) {
          return nameLength
        } else {
          return acc
        }
      },
      0
    )

    program.addHelpText(
      'after',
      `\nMetrics:\n${Object.entries(params.metricsSpec)
        .map(([metricName, metricSpec]) => {
          const parsedName = opSnakeCase(metricName)
          return `  ${parsedName}${' '.repeat(
            longestMetricNameLength - parsedName.length + 2
          )}${metricSpec.desc} (type: ${metricSpec.type.name})`
        })
        .join('\n')}
      `
    )

    // Load all configuration values from the environment and argv.
    program.parse()
    dotenv.config()
    const config = new Config(params.name)
    config.load({
      env: true,
      argv: true,
    })

    // Clean configuration values using the options spec.
    // Since BCFG turns everything into lower case, we're required to turn all of the input option
    // names into lower case for the validation step. We'll turn the names back into their original
    // names when we're done.
    const cleaned = cleanEnv<TOptions>(
      { ...config.env, ...config.args },
      Object.entries(params.optionsSpec || {}).reduce((acc, [key, val]) => {
        acc[key.toLowerCase()] = val.validator({
          desc: val.desc,
          default: val.default,
        })
        return acc
      }, {}) as any,
      Object.entries(params.options || {}).reduce((acc, [key, val]) => {
        acc[key.toLowerCase()] = val
        return acc
      }, {}) as any
    )

    // Turn the lowercased option names back into camelCase.
    this.options = Object.keys(params.optionsSpec || {}).reduce((acc, key) => {
      acc[key] = cleaned[key.toLowerCase()]
      return acc
    }, {}) as TOptions

    // Make sure all options are defined.
    for (const [optionName, optionSpec] of Object.entries(params.optionsSpec)) {
      if (
        optionSpec.default === undefined &&
        this.options[optionName] === undefined
      ) {
        throw new Error(`missing required option: ${optionName}`)
      }
    }

    // Create the metrics objects.
    this.metrics = Object.keys(params.metricsSpec || {}).reduce((acc, key) => {
      const spec = params.metricsSpec[key]
      acc[key] = new spec.type({
        name: `${opSnakeCase(params.name)}_${opSnakeCase(key)}`,
        help: spec.desc,
        labelNames: spec.labels || [],
      })
      return acc
    }, {}) as TMetrics & StandardMetrics

    // Create the metrics server.
    this.metricsRegistry = prometheus.register
    this.port = this.options.port
    this.hostname = this.options.hostname

    // Set up everything else.
    this.loopIntervalMs = this.options.loopIntervalMs
    this.logger = new Logger({ name: params.name })
    this.healthy = true

    // Gracefully handle stop signals.
    const maxSignalCount = 3
    let currSignalCount = 0
    const stop = async (signal: string) => {
      // Allow exiting fast if more signals are received.
      currSignalCount++
      if (currSignalCount === 1) {
        this.logger.info(`stopping service with signal`, { signal })
        await this.stop()
        process.exit(0)
      } else if (currSignalCount >= maxSignalCount) {
        this.logger.info(`performing hard stop`)
        process.exit(0)
      } else {
        this.logger.info(
          `send ${maxSignalCount - currSignalCount} more signal(s) to hard stop`
        )
      }
    }

    // Handle stop signals.
    process.on('SIGTERM', stop)
    process.on('SIGINT', stop)

    // Set metadata synthetic metric.
    this.metrics.metadata.set(
      {
        name: params.name,
        version: params.version,
        ...publicOptionNames.reduce((acc, key) => {
          if (key in stdOptionsSpec) {
            acc[key] = this.options[key].toString()
          } else {
            acc[key] = config.str(key)
          }
          return acc
        }, {}),
      },
      1
    )

    // Collect default node metrics.
    prometheus.collectDefaultMetrics({
      register: this.metricsRegistry,
      labels: { name: params.name, version: params.version },
    })
  }

  /**
   * Runs the main function. If this service is set up to loop, will repeatedly loop around the
   * main function. Will also catch unhandled errors.
   */
  public async run(): Promise<void> {
    // Start the app server if not yet running.
    if (!this.server) {
      this.logger.info('starting app server')

      // Start building the app.
      const app = express()

      // Body parsing.
      app.use(bodyParser.urlencoded({ extended: true }))

      // Keep the raw body around in case the application needs it.
      app.use(
        bodyParser.json({
          verify: (req, res, buf, encoding) => {
            ;(req as any).rawBody = buf?.toString(encoding || 'utf8') || ''
          },
        })
      )

      // Logging.
      app.use(
        morgan('short', {
          stream: {
            write: (str: string) => {
              this.logger.info(`server log`, {
                log: str,
              })
            },
          },
        })
      )

      // Health status.
      app.get('/healthz', async (req, res) => {
        return res.json({
          ok: this.healthy,
          version: this.params.version,
        })
      })

      // Register user routes.
      const router = express.Router()
      if (this.routes) {
        this.routes(router)
      }

      // Metrics.
      // Will expose a /metrics endpoint by default.
      app.use(
        promBundle({
          promRegistry: this.metricsRegistry,
          includeMethod: true,
          includePath: true,
          includeStatusCode: true,
          normalizePath: (req) => {
            for (const layer of router.stack) {
              if (layer.route && req.path.match(layer.regexp)) {
                return layer.route.path
              }
            }

            return '/invalid_path_not_a_real_route'
          },
        })
      )

      app.use('/api', router)

      // Wait for server to come up.
      await new Promise((resolve) => {
        this.server = app.listen(this.port, this.hostname, () => {
          resolve(null)
        })
      })

      this.logger.info(`app server started`, {
        port: this.port,
        hostname: this.hostname,
      })
    }

    if (this.init) {
      this.logger.info('initializing service')
      await this.init()
      this.logger.info('service initialized')
    }

    if (this.loop) {
      this.logger.info('starting main loop')
      this.running = true

      const doLoop = async () => {
        try {
          this.mainPromise = this.main()
          await this.mainPromise
        } catch (err) {
          this.metrics.unhandledErrors.inc()
          this.logger.error('caught an unhandled exception', {
            message: err.message,
            stack: err.stack,
            code: err.code,
          })
        }

        // Sleep between loops if we're still running (service not stopped).
        if (this.running) {
          this.pollingTimeout = setTimeout(doLoop, this.loopIntervalMs)
        }
      }
      doLoop()
    } else {
      this.logger.info('running main function')
      await this.main()
    }
  }

  /**
   * Tries to gracefully stop the service. Service will continue running until the current loop
   * iteration is finished and will then stop looping.
   */
  public async stop(): Promise<void> {
    this.logger.info('stopping main loop...')
    this.running = false
    clearTimeout(this.pollingTimeout)
    this.logger.info('waiting for main to complete')
    // if main is in the middle of running wait for it to complete
    await this.mainPromise
    this.logger.info('main loop stoped.')

    // Shut down the metrics server if it's running.
    if (this.server) {
      this.logger.info('stopping metrics server')
      await new Promise((resolve) => {
        this.server.close(() => {
          resolve(null)
        })
      })
      this.logger.info('metrics server stopped')
      this.server = undefined
    }
  }

  /**
   * Initialization function. Runs once before the main function.
   */
  protected init?(): Promise<void>

  /**
   * Initialization function for router.
   *
   * @param router Express router.
   */
  protected routes?(router: ExpressRouter): Promise<void>

  /**
   * Main function. Runs repeatedly when run() is called.
   */
  protected abstract main(): Promise<void>
}
