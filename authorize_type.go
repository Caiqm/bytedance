package bytedance

type Applet struct {
	AuxParam
}

func (a Applet) NeedSign() bool {
	return false
}

func (a Applet) NeedSecret() bool {
	return true
}

// ClientToken 获取应用授权调用凭证 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/interface-request-credential/non-user-authorization/get-client_token
type ClientToken struct {
	AuxParam
	ClientKey    string `json:"client_key"`    // 应用唯一标识，对应小程序id
	ClientSecret string `json:"client_secret"` // 应用唯一标识对应的密钥，对应小程序的app secret，可以在开发者后台获取
	GrantType    string `json:"grant_type"`    // 固定值“client_credential”
}

func (a ClientToken) NeedSign() bool {
	return false
}

func (a ClientToken) NeedAppId() bool {
	return false
}

func (a ClientToken) NeedSecret() bool {
	return false
}

// ClientTokenRsp 获取应用授权调用凭证响应参数
type ClientTokenRsp struct {
	Data struct {
		AccessToken string `json:"access_token,omitempty"` // client_token 接口调用凭证
		Description string `json:"description,omitempty"`  // 错误码描述
		ErrorCode   int    `json:"error_code,omitempty"`   // 错误码
		ExpiresIn   int    `json:"expires_in,omitempty"`   // client_token 接口调用凭证超时时间，单位（秒）
	} `json:"data"` // client_token信息
	Message string `json:"message,omitempty"` // 请求响应
	Extra   struct {
		Logid string `json:"logid"` // 日志记录ID
		Now   int64  `json:"now"`   // 当前时间戳
	} `json:"extra"`
}

// GetAccessToken 小程序的全局唯一调用凭据 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/interface-request-credential/non-user-authorization/get-access-token
type GetAccessToken struct {
	Applet
	GrantType string `json:"grant_type"` // 获取 access_token 时值为 client_credential
}

// GetAccessTokenRsp 小程序的全局唯一调用凭据响应参数
type GetAccessTokenRsp struct {
	Error
	Data struct {
		AccessToken string `json:"access_token"` // 获取的 access_token
		ExpiresIn   int64  `json:"expires_in"`   // access_token 有效时间，单位：秒
	} `json:"data"`
}

// Code2Session 小程序登录 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/log-in/code-2-session
type Code2Session struct {
	Applet
	AnonymousCode string `json:"anonymous_code,omitempty"` // login 接口返回的匿名登录凭证
	Code          string `json:"code,omitempty"`           // login 接口返回的登录凭证
}

// Code2SessionRsp 小程序登录响应参数
type Code2SessionRsp struct {
	Error
	Data struct {
		SessionKey      string `json:"session_key"`      // 会话密钥，如果请求时有 code 参数才会返回
		Openid          string `json:"openid"`           // 用户在当前小程序的 ID，如果请求时有 code 参数才会返回
		AnonymousOpenid string `json:"anonymous_openid"` // 匿名用户在当前小程序的 ID，如果请求时有 anonymous_code 参数才会返回
		Unionid         string `json:"unionid"`          // 用户在小程序平台的唯一标识符，请求时有 code 参数才会返回。如果开发者拥有多个小程序，可通过 unionid 来区分用户的唯一性。
	} `json:"data"`
}

// GetPhoneNumber 获取手机号 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/guide/open-capabilities/acquire-phone-number-acquire/
type GetPhoneNumber struct {
	EncryptedData string `json:"encrypted_data"`
	SessionKey    string `json:"session_key"`
	Iv            string `json:"iv"`
}

// GetPhoneNumberRsp 获取手机号响应参数
type GetPhoneNumberRsp struct {
	PhoneNumber     string        `json:"phoneNumber"`     // 用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string        `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string        `json:"countryCode"`     // 区号
	Watermark       WatermarkData `json:"watermark"`
}

type WatermarkData struct {
	Appid     string `json:"appid"`
	Timestamp int    `json:"timestamp"`
}

// GetUserPhoneNumber 获取用户手机号 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/basic-abilities/log-in/get-phone-number
type GetUserPhoneNumber struct {
	Applet
	AccessToken string `json:"access_token"` // 调用https://open.douyin.com/oauth/client_token/生成的token
	Code        string `json:"code"`         // 获取手机号的凭证code，通过前端getPhoneNumber组件获取
}

func (a GetUserPhoneNumber) NeedAccessToken() bool {
	return true
}

// GetUserPhoneNumberRsp 获取用户手机号响应参数
type GetUserPhoneNumberRsp struct {
	Error
	LogId string `json:"log_id"` // 抖音开放平台统一日志id
	Data  string `json:"data"`   // 通过开发者在「抖音开放平台-控制台-找到对应的小程序-开发-开发配置」配置的应用公钥 进行RSA非对称加密后的密文数据，需要开发者通过对应的应用私钥进行解密。应用公私钥由开发者生成，并将公钥录入在「应用公钥」上，私钥由开发者保管。
}
