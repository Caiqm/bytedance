package bytedance

// QrcodeCreate 生成QRCodeV2 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/url-and-qrcode/qrcode/create-qr-code-v2
type QrcodeCreate struct {
	AuxParam
	AccessToken  string     `json:"access_token"`         // 调用/oauth/client_token/生成的 token，此 token 不需要用户授权。
	AppName      string     `json:"app_name"`             // 是打开二维码的字节系 app 名称，默认为今日头条toutiao，今日头条极速版toutiao_lite，抖音douyin，抖音极速版douyin_lite，抖音火山版huoshan，全宿主，即在哪个宿主扫码就在哪个宿主打开小程序，若某个宿主没有该小程序，兜底往抖音app跳转all
	Path         string     `json:"path"`                 // 小程序/小游戏启动参数，小程序则格式为 encode({path}?{query})，小游戏则格式为 JSON 字符串，默认为空
	Width        int        `json:"width"`                // 二维码宽度，单位 px，最小 280px，最大 1280px，默认为 430px
	LineColor    LineColor  `json:"line_color,omitempty"` // 二维码线条颜色，默认为黑色
	Background   Background `json:"background,omitempty"` // 二维码背景颜色，默认为白色
	SetIcon      bool       `json:"set_icon"`             // 是否展示小程序/小游戏 icon，默认不展示
	IsCircleCode bool       `json:"is_circle_code"`       // 默认是false，是否生成抖音码，默认不生成（抖音码不支持自定义颜色）
}

func (q QrcodeCreate) NeedSign() bool {
	return false
}

func (q QrcodeCreate) NeedAccessToken() bool {
	return true
}

type LineColor struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type Background struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

// QrcodeCreateRsp 生成QRCodeV2响应参数
type QrcodeCreateRsp struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
	Data   struct {
		Img string `json:"img"` // 返回的图片数据以base64进行编码
	} `json:"data"`
}
