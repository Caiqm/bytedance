package bytedance

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/goccy/go-json"
)

// ClientToken 获取应用授权调用凭证 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/interface-request-credential/non-user-authorization/get-client_token
// POST https://open.douyin.com/oauth/client_token/ https://open-sandbox.douyin.com/oauth/client_token/
func (c *Client) ClientToken(param ClientToken) (result *ClientTokenRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// GetAccessToken 小程序的全局唯一调用凭据 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/interface-request-credential/non-user-authorization/get-access-token
// POST https://developer.toutiao.com/api/apps/v2/token https://open-sandbox.douyin.com/api/apps/v2/token
func (c *Client) GetAccessToken(param GetAccessToken) (result *GetAccessTokenRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// Code2Session 小程序登录 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/log-in/code-2-session
// POST https://developer.toutiao.com/api/apps/v2/jscode2session https://open-sandbox.douyin.com/api/apps/v2/jscode2session
func (c *Client) Code2Session(param Code2Session) (result *Code2SessionRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// GetPhoneNumber 获取手机号 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/guide/open-capabilities/acquire-phone-number-acquire/
func (c *Client) GetPhoneNumber(param GetPhoneNumber) (result *GetPhoneNumberRsp, err error) {
	src, _ := base64.StdEncoding.DecodeString(param.EncryptedData)
	_key, _ := base64.StdEncoding.DecodeString(param.SessionKey)
	_iv, _ := base64.StdEncoding.DecodeString(param.Iv)
	// 解密
	block, _ := aes.NewCipher(_key)
	mode := cipher.NewCBCDecrypter(block, _iv)
	dst := make([]byte, len(src))
	mode.CryptBlocks(dst, src)
	err = json.Unmarshal(dst, &result)
	return
}
