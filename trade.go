package bytedance

// 担保交易 ========================================================

// TradeEcPay 发起支付 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/pay-list/pay/
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

// 通用交易系统 ========================================================

// TradeTransactionPay 生成下单参数与签名 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/payment/trade-system/general/order/request-order-data-sign
func (c *Client) TradeTransactionPay(param TradeTransaction) (result *TradeTransactionResponse, err error) {
	err = c.GenerateOrderParametersAndSignatures(param, &result)
	return
}

// TradeTransactionQuery 查询订单 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/payment/trade-system/general/order/query-order
// POST https://open.douyin.com/api/trade_basic/v1/developer/order_query/
func (c *Client) TradeTransactionQuery(param TradeTransactionQuery) (result *TradeTransactionQueryResponse, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// TradeRefundCreate 发起退款 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/payment/trade-system/general/refund/create-refund
// POST https://open.douyin.com/api/trade_basic/v1/developer/refund_create/
func (c *Client) TradeRefundCreate(param TradeRefundCreate) (result *TradeRefundCreateResponse, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// TradeRefundCreateQuery 查询退款 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/payment/trade-system/general/refund/query-refund
// POST https://open.douyin.com/api/trade_basic/v1/developer/refund_query/
func (c *Client) TradeRefundCreateQuery(param TradeRefundCreateQuery) (result *TradeRefundCreateQueryResponse, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// RefundAuditCallback 同步退款审核结果 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/payment/trade-system/general/refund/refund-audit
// POST https://open.douyin.com/api/trade_basic/v1/developer/refund_audit_callback/
func (c *Client) RefundAuditCallback(param RefundAuditCallback) (result *RefundAuditCallbackResponse, err error) {
	err = c.doRequest("POST", param, &result)
	return
}
