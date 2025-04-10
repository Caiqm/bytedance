package bytedance

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
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

// GetPhoneNumberByCode getPhoneNumber组件code换取手机号 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/basic-abilities/log-in/get-phone-number
// POST https://open.douyin.com/api/apps/v1/get_phonenumber_info/
func (c *Client) GetPhoneNumberByCode(param GetUserPhoneNumber) (result *GetPhoneNumberRsp, err error) {
	var phoneRsp *GetUserPhoneNumberRsp
	err = c.doRequest("POST", param, &phoneRsp)
	if err != nil {
		return
	}
	// encryptData rsa加密后的数据
	// privateKey 应用私钥 需要把RSA的私钥头部和尾部的标识去掉，并且整合成一行字符串
	// decryptData 解密后的数据
	decryptData, err := c.rsaDecryptByPrivateKeyStr(phoneRsp.Data, c.privateKey)
	if err != nil {
		err = fmt.Errorf("[Error] RsaDecryptByPrivateKeyStr err=%v\n", err)
		return
	}
	fmt.Printf("[Info] decryptData=%v", decryptData)
	if err = json.Unmarshal([]byte(decryptData), &result); err != nil {
		err = fmt.Errorf("[Error] json.Unmarshal err=%v", err)
		return
	}
	return
}

// rsaDecryptByPrivateKeyStr rsa解密
func (c *Client) rsaDecryptByPrivateKeyStr(ciperData, privateKey string) (originText string, err error) {
	//读取私钥
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		err = fmt.Errorf("[Error] base64 decode privateKey=%s err=%v", privateKey, err)
		return "", err
	}
	priRSA, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		err = fmt.Errorf("[Error] x509.ParsePKCS1PrivateKey err=%v", err)
		return "", err
	}
	ciperDateBytes, err := base64.StdEncoding.DecodeString(ciperData)
	if err != nil {
		err = fmt.Errorf("[Error] base64 decode ciperData=%s, err=%v", ciperData, err)
		return "", err
	}
	originTextBytes, err := rsa.DecryptPKCS1v15(rand.Reader, priRSA, ciperDateBytes)
	if err != nil {
		err = fmt.Errorf("[Error] rsa.DecryptPKCS1v15 err=%v", err)
		return "", err
	}
	return string(originTextBytes), nil
}
