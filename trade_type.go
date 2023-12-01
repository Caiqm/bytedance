package bytedance

type Trade struct {
	AuxParam
	OutOrderNo string `json:"out_order_no"` // 开发者侧的订单号，商户分配支付单号，标识进行退款的订单
}

// TradeEcPay 担保交易 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/pay-list/pay/
type TradeEcPay struct {
	Trade
	TotalAmount     int64  `json:"total_amount"` // 支付价格。 单位为[分]
	Subject         string `json:"subject"`      // 商品描述。 长度限制不超过 128 字节且不超过 42 字符
	Body            string `json:"body"`         // 商品详情 长度限制不超过 128 字节且不超过 42 字符
	ValidTime       int64  `json:"valid_time"`   // 订单过期时间(秒)。最小5分钟，最大2天，小于5分钟会被置为5分钟，大于2天会被置为2天，取值范围：[300,172800]
	CpExtra         string `json:"cp_extra"`     // 开发者自定义字段，回调原样回传。 超过最大长度会被截断(2048)
	NotifyUrl       string `json:"notify_url"`
	ThirdPartyId    string `json:"thirdparty_id"`               // 第三方平台服务商 id，非服务商模式留空，服务商模式接入必传
	StoreUid        string `json:"store_uid,omitempty"`         // 可用此字段指定本单使用的收款商户号（目前为灰度功能，需要联系平台运营添加白名单，白名单添加1小时后生效；未在白名单的小程序，该字段不生效）多门店模式下可传
	DisableMsg      int64  `json:"disable_msg,omitempty"`       // 是否屏蔽支付完成后推送用户抖音消息，1-屏蔽 0-非屏蔽，默认为0。 特别注意： 若接入POI, 请传1。因为POI订单体系会发消息，所以不用再接收一次担保支付推送消息
	MsgPage         string `json:"msg_page,omitempty"`          // 支付完成后推送给用户的抖音消息跳转页面，开发者需要传入在app.json中定义的链接，如果不传则跳转首页。
	ExpandOrderInfo string `json:"expand_order_info,omitempty"` // 订单拓展信息，{"original_delivery_fee":10,"actual_delivery_fee":10}
	LimitPayWay     string `json:"limit_pay_way,omitempty"`     // 屏蔽指定支付方式，屏蔽多个支付方式，请使用逗号","分割，枚举值：屏蔽微信支付：LIMIT_WX，屏蔽支付宝支付：LIMIT_ALI
}

// ExpandOrderInfo 参数
type ExpandOrderInfo struct {
	OriginalDeliveryFee int `json:"original_delivery_fee,omitempty"` // 配送费原价，单位为[分]，仅外卖小程序需要传对应信息
	ActualDeliveryFee   int `json:"actual_delivery_fee,omitempty"`   // 实付配送费，单位为[分]，仅外卖小程序需要传对应信息
}

// TradeEcPayRsp 担保交易响应参数
type TradeEcPayRsp struct {
	Error
	Data struct {
		OrderId    string `json:"order_id"`    // 抖音侧唯一订单号
		OrderToken string `json:"order_token"` // 签名后的订单信息
	} `json:"data"`
}

// TradeOrderQuery 支付结果查询 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/pay-list/query
type TradeOrderQuery struct {
	Trade
	ThirdPartyId string `json:"thirdparty_id"` // 第三方平台服务商 id，非服务商模式留空，服务商模式接入必传
}

// TradeOrderQueryRsp 支付结果查询响应参数
type TradeOrderQueryRsp struct {
	Error
	OutOrderNo  string      `json:"out_order_no"`       // 开发者侧的订单号
	OrderId     string      `json:"order_id"`           // 抖音侧唯一订单号
	PaymentInfo PaymentInfo `json:"payment_info"`       // 支付信息
	CpsInfo     interface{} `json:"cps_info,omitempty"` // 若该订单为CPS订单，该字段会返回该笔订单的CPS相关信息
}

// PaymentInfo 支付信息
type PaymentInfo struct {
	TotalFee    int    `json:"total_fee"`    // 支付金额，单位为分
	OrderStatus string `json:"order_status"` // 支付状态枚举值： SUCCESS：成功 TIMEOUT：超时未支付 PROCESSING：处理中 FAIL：失败
	PayTime     string `json:"pay_time"`     //  支付完成时间，order_status不为SUCCESS时会返回默认值空字符串，order_status为SUCCESS时返回非空字符串，格式为"yyyy-MM-dd HH:mm:ss"
	Way         int    `json:"way"`          // 支付渠道，order_status不为SUCCESS时会返回默认值0，order_status为SUCCESS时会返回以下枚举：1-微信支付，2-支付宝支付，10-抖音支付
	ChannelNo   string `json:"channel_no"`   // 支付渠道侧的支付单号
	SellerUid   string `json:"seller_uid"`   // 该笔交易卖家商户号
	ItemId      string `json:"item_id"`      // 订单来源视频对应视频id
	CpExtra     string `json:"cp_extra"`     // 开发者自定义字段
}

// CpsInfo CPS相关信息
type CpsInfo struct {
	ShareAmount string `json:"share_amount"` // 达人分佣金额，单位为分。后续商户在进行分账时需要注意可分账金额应扣除达人分佣金额。
	DouyinId    string `json:"douyin_id"`    // 达人抖音号
	Nickname    string `json:"nickname"`     // 达人昵称
}

// TradeRefund 发起退款 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/refund-list/refund
type TradeRefund struct {
	Trade
	OutRefundNo  string `json:"out_refund_no"` // 商户分配退款号，保证在商户中唯一
	Reason       string `json:"reason"`        // 退款原因
	RefundAmount int    `json:"refund_amount"` // 退款金额，单位分
	CpExtra      string `json:"cp_extra"`      // 开发者自定义字段，回调原样回传
	NotifyUrl    string `json:"notify_url"`    // 商户自定义回调地址，必须以 https 开头，支持 443 端口
	ThirdPartyId string `json:"thirdparty_id"` // 第三方平台服务商id，服务商模式接入必传，非服务商模式留空
	DisableMsg   int64  `json:"disable_msg"`   // 是否屏蔽支付完成后推送用户抖音消息，1-屏蔽 0-非屏蔽，默认为0。 特别注意： 若接入POI, 请传1。因为POI订单体系会发消息，所以不用再接收一次担保支付推送消息
	MsgPage      string `json:"msg_page"`      // 支付完成后推送给用户的抖音消息跳转页面，开发者需要传入在app.json中定义的链接，如果不传则跳转首页。
}

// TradeRefundRsp 发起退款响应参数
type TradeRefundRsp struct {
	Error
	RefundNo string `json:"refund_no"` // 担保交易服务端退款单号
}

// TradeRefundQuery 退款结果查询 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/refund-list/query
type TradeRefundQuery struct {
	AuxParam
	OutRefundNo  string `json:"out_refund_no"` // 商户分配退款号，保证在商户中唯一
	ThirdPartyId string `json:"thirdparty_id"` // 第三方平台服务商id，服务商模式接入必传，非服务商模式留空
}

// TradeRefundQueryRsp 退款结果查询响应参数
type TradeRefundQueryRsp struct {
	Error
	RefundInfo struct {
		RefundNo     string `json:"refund_no"`      // 抖音退款单号
		RefundAmount int    `json:"refund_amount"`  // 退款金额，单位为分
		RefundStatus string `json:"refund_status"`  // 退款状态枚举 SUCCESS：成功 FAIL：失败 PROCESSING：处理中
		RefundedAt   int    `json:"refunded_at"`    // 退款时间，Unix 时间戳，10 位，整型数，秒级
		IsAllSettled bool   `json:"is_all_settled"` // 退款账户枚举： TRUE：分账后退款，现金户出款 FALSE：分账前退款，在途户出款
		CpExtra      string `json:"cp_extra"`       // 开发者自定义字段，回调原样回传
		Msg          string `json:"msg"`            // 退款错误描述
	} `json:"refundInfo"`
}
