package bytedance

// QrcodeCreate 生成QRCodeV2 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/url-and-qrcode/qrcode/create-qr-code-v2
// POST https://open.douyin.com/api/apps/v1/qrcode/create/ 需要小程序通过试运营期
func (c *Client) QrcodeCreate(param QrcodeCreate) (result *QrcodeCreateRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}
