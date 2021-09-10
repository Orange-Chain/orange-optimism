import React from 'react';
import { connect } from 'react-redux';
import { isEqual } from 'lodash';
import BN from 'bignumber.js';

import { closeModal, openAlert, openError } from 'actions/uiAction';
import { getFarmInfo } from 'actions/farmAction';

import Button from 'components/button/Button';
import Modal from 'components/modal/Modal';
import Input from 'components/input/Input';
import { logAmount, powAmount } from 'util/amountConvert';

import networkService from 'services/networkService';

import * as S from './FarmModal.styles';
import { Typography } from '@material-ui/core';

class FarmDepositModal extends React.Component {

  constructor(props) {
    super(props);

    const { open } = this.props;
    const { stakeToken } = this.props.farm;

    this.state = {
      open,
      stakeToken,
      stakeValue: '',
      // allowance
      approvedAllowance: 0,
      // loading
      loading: false,
    }
  }

  async componentDidUpdate(prevState) {

    const { open } = this.props
    const { stakeToken } = this.props.farm

    if (prevState.open !== open) {
      this.setState({ open })
    }

    if (!isEqual(prevState.farm.stakeToken, stakeToken)) {
      let approvedAllowance = powAmount(10, 50)
      // Set to some very big number
      // There is no need to check allowance for depositing ETH
      if (stakeToken.currency !== networkService.L1_ETH_Address) {
        approvedAllowance = await networkService.checkAllowance(
          stakeToken.currency,
          stakeToken.LPAddress
        )
      }
      this.setState({ approvedAllowance, stakeToken })
    }

  }

  getMaxTransferValue () {
    const { stakeToken } = this.state;
    // const transferingBalanceObject = (stakeToken.L1orL2Pool === 'L1LP' ? layer1Balance : layer2Balance)
    //   .find(i => i.currency === stakeToken.currency);
    // if (!transferingBalanceObject) {
    //   return;
    // }
    return logAmount(stakeToken.balance, stakeToken.decimals);
  }

  handleClose() {
    this.props.dispatch(closeModal("farmDepositModal"))
  }

  async handleApprove() {

    const { stakeToken, stakeValue } = this.state

    this.setState({ loading: true })

    let approveTX

    if (stakeToken.L1orL2Pool === 'L2LP') {
      approveTX = await networkService.approveERC20_L2LP(
        powAmount(stakeValue, stakeToken.decimals),
        stakeToken.currency,
      )
    } else if (stakeToken.L1orL2Pool === 'L1LP') {
      approveTX = await networkService.approveERC20_L1LP(
        powAmount(stakeValue, stakeToken.decimals),
        stakeToken.currency,
      )
    }

    if (approveTX) {
      this.props.dispatch(openAlert("Amount was approved"))
      let approvedAllowance = powAmount(10, 50)
      // There is no need to check allowance for depositing ETH
      if (stakeToken.currency !== networkService.L1_ETH_Address) {
        approvedAllowance = await networkService.checkAllowance(
          stakeToken.currency,
          stakeToken.LPAddress
        )
      }

      this.setState({ approvedAllowance, loading: false })
    } else {
      this.props.dispatch(openError("Failed to approve amount"))
      this.setState({ loading: false })
    }
  }

  async handleConfirm() {

    const { stakeToken, stakeValue } = this.state

    this.setState({ loading: true })

    const addLiquidityTX = await networkService.addLiquidity(
      stakeToken.currency,
      stakeValue,
      stakeToken.L1orL2Pool,
      stakeToken.decimals
    )

    if (addLiquidityTX) {
      this.props.dispatch(openAlert("Your liquidity was added"))
      this.props.dispatch(getFarmInfo())
      this.setState({ loading: false, stakeValue: '' })
      this.props.dispatch(closeModal("farmDepositModal"))
    } else {
      this.props.dispatch(openError("Failed to add liquidity"))
      this.setState({ loading: false, stakeValue: '' })
    }
  }

  render() {

    const {
      open,
      stakeToken,
      stakeValue,
      approvedAllowance,
      loading,
    } = this.state;

    const valueIsValid = Number(stakeValue) > 0

    return (

      <Modal
        open={open}
        maxWidth="md"
        onClose={()=>{this.handleClose()}}
      >

        <Typography variant="h2" sx={{fontWeight: 700, mb: 3}}>
          Stake {`${stakeToken.symbol}`}
        </Typography>

        <Input
          placeholder={`Amount to stake`}
          value={stakeValue}
          type="number"
          onChange={i=>{this.setState({stakeValue: i.target.value})}}
          unit={stakeToken.symbol}
          maxValue={this.getMaxTransferValue()}
          newStyle
          variant="standard"
        />

        {!valueIsValid && stakeValue !== '' ?
          <Typography variant="body2" sx={{mt: 2}}>
            You can't add 0 to the pool
          </Typography>
          : null
        }

        {Number(stakeValue) > Number(this.getMaxTransferValue()) &&
          <Typography variant="body2" sx={{mt: 2}}>
            You don't have enough {stakeToken.symbol} to stake.
          </Typography>
        }

        {(new BN(approvedAllowance).gte(powAmount(stakeValue, 18)) || stakeValue === '') &&
          <S.WrapperActions>
            <Button
              onClick={()=>{this.handleClose()}}
              color="neutral"
              size="large"
            >
              CANCEL
            </Button>
            <Button
              onClick={()=>{this.handleConfirm()}}
              disabled={Number(this.getMaxTransferValue()) < Number(stakeValue) || stakeValue === '' || !stakeValue || !valueIsValid}
              loading={loading}
              color='primary'
              size="large"
              variant="contained"
              // fullWidth={isMobile}
            >
              STAKE!
            </Button>
          </S.WrapperActions>
        }

        {new BN(approvedAllowance).lt(new BN(powAmount(stakeValue, 18))) &&
          <>
            <Typography variant="body2" sx={{mt: 2}}>
              To stake {stakeValue} {stakeToken.symbol},
              you first need to approve this amount.
            </Typography>

            <S.WrapperActions>
              <Button
                onClick={()=>{this.handleClose()}}
                color="neutral"
                size="large"
              >
                CANCEL
              </Button>
              <Button
                onClick={()=>{this.handleApprove()}}
                loading={loading}
                disabled={Number(this.getMaxTransferValue()) < Number(stakeValue)}
                color='primary'
                size="large"
                variant="contained"
              >
                APPROVE AMOUNT
              </Button>
            </S.WrapperActions>
          </>
        }
      </Modal>
    )
  }
}

const mapStateToProps = state => ({
  ui: state.ui,
  farm: state.farm,
  balance: state.balance,
})

export default connect(mapStateToProps)(FarmDepositModal)
