package bytedance

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/goccy/go-json"
	"strconv"
	"strings"
	"time"
)

// 通用交易系统支付
func (c *Client) GenerateOrderParametersAndSignatures(param TradeTransaction, result interface{}) (err error) {
	reqOrderDataJson, err := json.Marshal(param)
	if err != nil {
		return
	}
	var (
		// 请求时间戳
		timestamp = strconv.FormatInt(time.Now().Unix(), 10)
		// 开发者填入自己的小程序app_id
		appId = c.appId
		// 随机字符串
		nonceStr = c.randStr(10)
		// 应用公钥版本,每次重新上传公钥后需要更新,可通过「开发管理-开发设置-密钥设置」处获取
		keyVersion = c.keyVersion
		// 应用私钥,用于加签 重要：1.测试时请修改为开发者自行生成的私钥;2.请勿将示例密钥用于生产环境;3.建议开发者不要将私钥文本写在代码中
		privateKeyStr = c.privateKey
		// 生成好的data
		data = string(reqOrderDataJson)
	)
	byteAuthorization, err := c.getByteAuthorization(privateKeyStr, data, appId, nonceStr, timestamp, keyVersion)
	if err != nil {
		err = fmt.Errorf("getByteAuthorization err: %s", err)
		return
	}
	requestData := map[string]string{
		"data":              data,
		"byteAuthorization": byteAuthorization,
	}
	requestDataByte, err := json.Marshal(requestData)
	if err != nil {
		return
	}
	if err = json.Unmarshal(requestDataByte, result); err != nil {
		return
	}
	return
}

// 生成byteAuthorization
func (c *Client) getByteAuthorization(privateKeyStr, data, appId, nonceStr, timestamp, keyVersion string) (string, error) {
	var byteAuthorization string
	// 读取私钥
	key, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(privateKeyStr, "\n", ""))
	if err != nil {
		return "", err
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		privateKey, err = x509.ParsePKCS1PrivateKey(key)
		if err != nil {
			err = fmt.Errorf("解析密钥失败，err: %v", err)
			return "", err
		}
	}
	// 生成签名
	signature, err := c.getSignature("POST", "/requestOrder", timestamp, nonceStr, data, privateKey.(*rsa.PrivateKey))
	if err != nil {
		return "", err
	}
	// 构造byteAuthorization
	byteAuthorization = fmt.Sprintf("SHA256-RSA2048 appid=%s,nonce_str=%s,timestamp=%s,key_version=%s,signature=%s", appId, nonceStr, timestamp, keyVersion, signature)
	return byteAuthorization, nil
}

// 生成签名
func (c *Client) getSignature(method, url, timestamp, nonce, data string, privateKey *rsa.PrivateKey) (string, error) {
	targetStr := method + "\n" + url + "\n" + timestamp + "\n" + nonce + "\n" + data + "\n"
	h := sha256.New()
	h.Write([]byte(targetStr))
	digestBytes := h.Sum(nil)

	signBytes, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digestBytes)
	if err != nil {
		return "", err
	}
	sign := base64.StdEncoding.EncodeToString(signBytes)

	return sign, nil
}

// 生成随机字符串
func (c *Client) randStr(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
