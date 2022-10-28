import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deployAndVerifyAndThen,
  getDeploymentAddress,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const proxyAdmin = await getDeploymentAddress(hre, 'ProxyAdmin')
  await deployAndVerifyAndThen({
    hre,
    name: 'L1CrossDomainMessengerProxy',
    contract: 'Proxy',
    args: [proxyAdmin],
    postDeployAction: async (contract) => {
      await assertContractVariable(contract, 'admin', proxyAdmin)
    },
  })
}

deployFn.tags = ['L1CrossDomainMessengerProxy', 'fresh']

export default deployFn
