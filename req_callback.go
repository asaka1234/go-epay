package go_epay

import (
	"errors"
	"github.com/asaka1234/go-epay/utils"
	"github.com/mitchellh/mapstructure"
)

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCancelCallback(req EPayDepositCancelBackReq, processor func(EPayDepositCancelBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	verifyResult := utils.VerifySignDeposit(params, cli.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}
	if req.SysNo != cli.MerchantID {
		return errors.New("merchanID is wrong!")
	}

	//开始处理
	return processor(req)
}

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositSucceedCallBack(req EPayDepositSucceedBackReq, processor func(EPayDepositSucceedBackReq) error) error {
	//验证签名
	params := map[string]interface{}{
		"bill_no": req.BillNo, //只是value的拼接
	}

	verifyResult := utils.VerifySignWithdraw(params, cli.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}
	if req.SysNo != cli.MerchantID {
		return errors.New("merchanID is wrong!")
	}

	//开始处理
	return processor(req)
}

//==========================================

// 充值的回调处理(传入一个处理函数)
func (cli *Client) WithdrawCancelCallBack(req EPayWithdrawCancelBackReq, processor func(EPayWithdrawCancelBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	verifyResult := utils.VerifySignDeposit(params, cli.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}
	if req.SysNo != cli.MerchantID {
		return errors.New("merchanID is wrong!")
	}

	//开始处理
	return processor(req)
}

// 充值的回调处理(传入一个处理函数)
func (cli *Client) WithdrawSucceedCallBack(req EPayWithdrawSucceedBackReq, processor func(EPayWithdrawSucceedBackReq) error) error {
	//验证签名
	params := map[string]interface{}{
		"bill_no": req.BillNo, //只是value的拼接
	}

	verifyResult := utils.VerifySignWithdraw(params, cli.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}
	if req.SysNo != cli.MerchantID {
		return errors.New("merchanID is wrong!")
	}

	//开始处理
	return processor(req)
}
