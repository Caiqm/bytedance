package bytedance

import "fmt"

const (
	kContentTypeJson = "application/json"
	kContentTypeForm = "multipart/form-data;charset=utf-8"
	kTimeFormat      = "2006-01-02 15:04:05"
)

const (
	kFieldAppId        = "app_id"
	kFieldSecret       = "secret"
	kFieldAccessToken  = "access_token"
	kFieldSign         = "sign"
	kFieldMsgSignature = "msg_signature"
	kFieldTimestamp    = "timestamp"
	kFieldNonce        = "nonce"
	kFieldMsg          = "msg"
	kFieldErrNo        = "err_no"
)

type Param interface {
	// NeedSign 是否需要签名，有的接口不需要签名，比如：小程序登录与获取手机号接口
	NeedSign() bool

	// NeedAppId 是否需要APPID，有的接口不需要APPID，比如：获取应用授权调用凭证
	NeedAppId() bool

	// NeedSecret 是否需要密钥
	NeedSecret() bool

	// NeedAccessToken 是否需要token
	NeedAccessToken() bool

	// ContentType 请求头，有的接口不是JSON的请求头
	ContentType() string
}

type AuxParam struct {
}

func (aux AuxParam) NeedSign() bool {
	return true
}

func (aux AuxParam) NeedAppId() bool {
	return true
}

func (aux AuxParam) NeedSecret() bool {
	return false
}

func (aux AuxParam) NeedAccessToken() bool {
	return false
}

func (aux AuxParam) ContentType() string {
	return kContentTypeJson
}

// Error 支付错误类
type Error struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
	ErrMsg  string `json:"err_msg"`
}

// 错误日志
func (e Error) Error() string {
	var errMsg string
	if e.ErrTips != "" {
		errMsg = e.ErrTips
	} else {
		errMsg = e.ErrMsg
	}
	return fmt.Sprintf("%d - %s", e.ErrNo, errMsg)
}
