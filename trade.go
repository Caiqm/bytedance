package bytedance

// TradeEcPay 担保交易 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/pay-list/pay/
// POST https://developer.toutiao.com/api/apps/ecpay/v1/create_order https://open-sandbox.douyin.com/api/apps/ecpay/v1/create_order
func (c *Client) TradeEcPay(param TradeEcPay) (result *TradeEcPayRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// TradeOrderQuery 支付结果查询 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/pay-list/query
// POST https://developer.toutiao.com/api/apps/ecpay/v1/query_order https://open-sandbox.douyin.com/api/apps/ecpay/v1/query_order
func (c *Client) TradeOrderQuery(param TradeOrderQuery) (result *TradeOrderQueryRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// TradeRefund 发起退款 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/refund-list/refund
// POST https://developer.toutiao.com/api/apps/ecpay/v1/create_refund https://open-sandbox.douyin.com/api/apps/ecpay/v1/create_refund
func (c *Client) TradeRefund(param TradeRefund) (result *TradeRefundRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// TradeRefundQuery 退款结果查询 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/refund-list/query
// POST https://developer.toutiao.com/api/apps/ecpay/v1/query_refund https://open-sandbox.douyin.com/api/apps/ecpay/v1/query_refund
func (c *Client) TradeRefundQuery(param TradeRefundQuery) (result *TradeRefundQueryRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}
