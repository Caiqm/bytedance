package bytedance

// OrderPush 订单同步 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/order/order-sync
// POST https://developer.toutiao.com/api/apps/order/v2/push
func (c *Client) OrderPush(param OrderPush) (result OrderPushRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}
