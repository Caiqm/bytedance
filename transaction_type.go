package bytedance

type Transaction struct {
	AuxParam
}

// TradeTransaction 通用交易系统生成下单参数与签名 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/payment/trade-system/general/order/request-order-data-sign
type TradeTransaction struct {
	Transaction
	SkuList          []SkuList        `json:"skuList"` // 下单商品信息
	OutOrderNo       string           `json:"outOrderNo"`
	TotalAmount      int              `json:"totalAmount"`                // 订单总金额, 单位：分
	PayExpireSeconds int              `json:"payExpireSeconds,omitempty"` // 支付超时时间，单位秒，例如 300 表示 300 秒后过期；不传或传 0 会使用默认值 300，不能超过48小时。
	PayNotifyUrl     string           `json:"payNotifyUrl,omitempty"`     // 支付结果通知地址，必须是 HTTPS 类型，传入后该笔订单将通知到此地址。
	OrderEntrySchema OrderEntrySchema `json:"orderEntrySchema"`           // 订单详情页
	MerchantUid      string           `json:"merchantUid,omitempty"`      // 开发者自定义收款商户号
	LimitPayWayList  []int            `json:"limitPayWayList,omitempty"`  // 屏蔽的支付方式，当开发者没有进件某个支付渠道，可在下单时屏蔽对应的支付方式。如[1, 2]表示屏蔽微信和支付宝 枚举说明： 1-微信 2-支付宝
}

type SkuList struct {
	SkuId       string           `json:"skuId"`                 // 外部商品id，如：号卡商品id、会员充值套餐id、某类服务id、付费工具id等
	Price       int              `json:"price"`                 // 价格 单位：分
	Quantity    int              `json:"quantity"`              // 购买数量，0 < quantity <= 100
	Title       string           `json:"title"`                 // 商品标题，长度 <= 256字节
	ImageList   []string         `json:"imageList"`             // 商品图片链接，长度 <= 512 字节 注意：目前只支持传入一项
	Type        int              `json:"type"`                  // 商品类型 示例：号卡商品：传101、剧集： 传404；详见商品类型枚举值（https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/open-capacity/trade-system/guide/general/basicrules#6a1682c4）
	TagGroupId  string           `json:"tagGroupId"`            // 交易规则标签组，查看对应商品类型的标签组ID（https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/trade-system/general/tag/tag_group_query#2b56d127）
	EntrySchema OrderEntrySchema `json:"entrySchema,omitempty"` // 商品详情页链接
	SkuAttr     string           `json:"skuAttr,omitempty"`     // 商品信息：需要将不同商品类型定义的具体结构，转换成json string，号卡类商品必填，即当前商品类型 type in [101、102、103、104、105、106、107]的商品必填；内容消费：当前商品类型 type in [402、403、404、405、406]的商品必填
}

type SkuAttr struct {
	PackageCost         PackageCost `json:"package_cost"`          // 号卡商品套餐详情
	CallDuration        int         `json:"call_duration"`         // 套餐包含的通话时长，单位：分钟，
	TrafficBundle       int         `json:"traffic_bundle"`        // 套餐包含的流量包大小，单位：G
	TelecomOperatorType string      `json:"telecom_operator_type"` // 提供套餐的运营方性质，枚举值 official：官方，private：私营
}

type PackageCost struct {
	Amount  int    `json:"amount"`   // 套餐售价，单位：分
	TimeLen int    `json:"time_len"` // 套餐时长
	Unit    string `json:"unit"`     // 时长单位，枚举值 year：年 month：月 day：天
}

type OrderEntrySchema struct {
	Path   string `json:"path"`   // 小程序xxx详情页跳转路径，没有前导的“/”，路径后不可携带query参数，路径中不可携带『？: & *』等特殊字符，路径只可以是『英文字符、数字、_、/ 』等组成，长度<=512byte，示例：page/path/index
	Params string `json:"params"` // xx情页路径参数，自定义的json结构，内部为k-v结构，序列化成字符串存入该字段，平台不限制，但是写入的内容需要能够保证生成访问xx详情的schema能正确跳转到小程序内部的xx详情页，长度须<=512byte，params内key不可重复。示例：'{"id":1234, "name":"hello"}'
}

// 通用交易系统返回
type TradeTransactionResponse struct {
	Data              string `json:"data"`              // 结果为string类型，且必须符合json格式
	ByteAuthorization string `json:"byteAuthorization"` // 构造byteAuthorization，其结果为string类型
}

type TransactionOrder struct {
	AuxParam
	OrderId     string `json:"order_id,omitempty"` // 抖音开放平台交易订单号，查询订单接口时 order_id 与 out_order_no 二选一，发起退款接口时必填
	AccessToken string `json:"access_token"`       // 调用https://open.douyin.com/oauth/client_token/生成的token
}

func (aux TransactionOrder) NeedSign() bool {
	return false
}

func (aux TransactionOrder) NeedSecret() bool {
	return false
}

func (aux TransactionOrder) NeedAppId() bool {
	return false
}

func (aux TransactionOrder) NeedAccessToken() bool { return true }

// TradeTransactionQuery 查询订单信息 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/payment/trade-system/general/order/query-order
type TradeTransactionQuery struct {
	TransactionOrder
	OutOrderNo string `json:"out_order_no,omitempty"` // 开发者的单号，order_id 与 out_order_no 二选一
}

// TradeTransactionQueryResponse 查询订单信息返回
type TradeTransactionQueryResponse struct {
	Error
	Data struct {
		OrderId        string `json:"order_id"`        // 抖音开平侧订单号
		OutOrderNo     string `json:"out_order_no"`    // 开发者侧订单号，与 order_id 一一对应
		AppId          string `json:"app_id"`          // 小程序id
		PayStatus      string `json:"pay_status"`      // 订单支付状态，PROCESS：订单处理中 支付处理中；SUCCESS：成功 支付成功；FAIL：失败 支付失败 暂无该情况会支付失败；TIMEOUT：用户超时未支付
		PayTime        int64  `json:"pay_time"`        // 支付成功时间，精度：毫秒。只有在支付成功时才会有值。
		PayChannel     int    `json:"pay_channel"`     // 支付渠道枚举，1：微信；2：支付宝；10：抖音支付；只有在支付成功时才会有值。
		Currency       string `json:"currency"`        // 币种
		ChannelPayId   string `json:"channel_pay_id"`  // 渠道支付单号，如：微信的支付单号、支付宝支付单号。只有在支付成功时才会有值。
		TradeTime      int64  `json:"trade_time"`      // 交易下单时间，精度：毫秒
		TotalAmount    int    `json:"total_amount"`    // 订单总金额，单位：分，支付金额 = total_amount - discount_amount
		DiscountAmount int    `json:"discount_amount"` // 订单优惠金额，单位：分，接入营销时请关注这个字段
		MerchantUid    string `json:"merchant_uid"`    // 开发者自定义收款商户号，小程序在抖音开平商户进件时会绑定一个收款账号，当用户交易时，资金会默认收款到此账号，如果某笔交易开发者希望收款到其他账号则需指定希望收款的账号id，此账号必须与默认账号一样属于同一开发者
		ItemOrderList  []struct {
			ItemOrderId             string `json:"item_order_id"`              // 交易系统商品单号
			SkuId                   string `json:"sku_id"`                     // 用户下单时传入的商品sku_id
			ItemOrderAmount         int    `json:"item_order_amount"`          // item单订单金额，单位分
			ItemOrderCurrencyAmount int64  `json:"item_order_currency_amount"` // item单订单币种金额，单位分
		} `json:"item_order_list"`
		OrderServiceStatus  string `json:"order_service_status"`  // 仅服务单会有该状态值，"INIT":服务单已创建；"CONFIRMED":已提交信息；"DONE":已完成；"CLOSED":已取消
		ServiceCancelReason string `json:"service_cancel_reason"` // 仅服务单会有该值；服务单取消原因
		SubmisssionTime     int64  `json:"submisssion_time"`      // 仅服务单会有该值；用户提交留资信息的时间，秒级时间戳
		TotalCurrencyAmount int64  `json:"total_currency_amount"` // 订单币种金额，单位分
	} `json:"data"`
	LogId string `json:"log_id"` // 日志id，排查问题时使用
}

// TradeRefundCreate 发起退款 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/payment/trade-system/general/refund/create-refund
type TradeRefundCreate struct {
	TransactionOrder
	OrderEntrySchema  OrderEntrySchema  `json:"order_entry_schema"`
	OutRefundNo       string            `json:"out_refund_no"`
	RefundReason      []ReasonItem      `json:"refund_reason"`
	RefundTotalAmount int64             `json:"refund_total_amount"` // 退款总金额，单位分
	CpExtra           string            `json:"cp_extra"`
	NotifyUrl         string            `json:"notify_url"`
	RefundAll         bool              `json:"refund_all"`
	ItemOrderDetail   []ItemOrderDetail `json:"item_order_detail"`
}

type ReasonItem struct {
	Code int64  `json:"code"` // 退款原因code，必须从以下code中选择:[{"code":101,"text":"不想要了"},{"code":102,"text":"商家服务原因"},{"code":103,"text":"商品质量问题"},{"code":999,"text":"其他"}]
	Text string `json:"text"` // 退款原因描述，开发者可自定义，长度<50
}

type ItemOrderDetail struct {
	ItemOrderId  string `json:"item_order_id"` // 商品单号，参见通用参数-重要 ID 字段说明
	RefundAmount int64  `json:"refund_amount"` // 该item_order需要退款的金额，单位[分]，必须>0且不能大于该 item_order 实付金额
}

// TradeRefundCreateResponse 发起退款响应参数
type TradeRefundCreateResponse struct {
	Error
	LogId string `json:"log_id"` // 日志id，排查问题时使用
	Data  struct {
		RefundAuditDeadline int64  `json:"refund_audit_deadline"` // 退款审核的最后期限，13位unix时间戳，精度：毫秒
		RefundId            string `json:"refund_id"`             // 抖音开放平台交易系统内部退款单号
	} `json:"data"`
}

// TradeRefundCreateQuery 查询退款 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/payment/trade-system/general/refund/query-refund
type TradeRefundCreateQuery struct {
	TransactionOrder
	OutRefundNo string `json:"out_refund_no"` // 开发者系统生成的退款单号。注意：refund_id , out_refund_no , order_id 三选一，不能都不填。
	RefundId    string `json:"refund_id"`     // 抖音开平内部交易退款单号，长度<= 64byte。注意：refund_id , out_refund_no , order_id 三选一，不能都不填。
}

// TradeRefundCreateQueryResponse 查询退款响应参数
type TradeRefundCreateQueryResponse struct {
	Error
	LogId string `json:"log_id"` // 日志id，排查问题时使用
	Data  struct {
		RefundList []struct {
			MerchantAuditDetail struct {
				AuditStatus         string `json:"audit_status"`          // 退款审核状态；INIT：初始化、TOAUDIT：待审核、AGREE：同意、DENY：拒绝、OVERTIME：超时未审核自动同意
				NeedRefundAudit     int    `json:"need_refund_audit"`     // 是否需要退款审核，1-需要审核、2-不需要审核
				RefundAuditDeadline int64  `json:"refund_audit_deadline"` // 退款审核的最后期限，过期无需审核，自动退款，13 位 unix 时间戳，精度：毫秒
				DenyMessage         string `json:"deny_message"`          // 不同意退款信息，长度 <= 512 byte>
			} `json:"merchant_audit_detail"` // 退款审核信息
			CreateAt          int64  `json:"create_at"`           // 退款创建时间，13位毫秒时间戳
			RefundAt          int64  `json:"refund_at"`           // 退款时间，13位毫秒时间戳，只有已退款才有退款时间
			RefundStatus      string `json:"refund_status"`       // 退款状态；退款中-PROCESSING、已退款-SUCCESS、退款失败-FAIL
			RefundTotalAmount int    `json:"refund_total_amount"` // 退款金额，单位[分]
			ItemOrderDetail   []struct {
				ItemOrderId  string `json:"item_order_id"` // 抖音开平侧的商品单号
				RefundAmount int    `json:"refund_amount"` // 该商品单退款金额，单位[分]
			} `json:"item_order_detail"`
			Message      string `json:"message"`       // 退款结果信息，可以通过该字段了解退款失败原因
			OrderId      string `json:"order_id"`      // 系统订单信息，开放平台生成的订单号
			OutRefundNo  string `json:"out_refund_no"` // 开发者系统生成的退款单号，与抖音开平退款单号唯一关联
			RefundId     string `json:"refund_id"`     // 系统退款单号，开放平台生成的退款单号
			RefundSource int32  `json:"refund_source"` // 退款来源，老的担保交易/1.0订单可能没有记录来源；1: 用户发起退款、2: 开发者发起退款、3: 过期自动退款、4: 抖音客服退款、5: 预约失败自动发起退款、6: 开发者拒绝接单退款、7: 后约单触发先买单退款
		} `json:"refund_list"`
	} `json:"data"`
}

// RefundAuditCallback 同步退款审核结果 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/payment/trade-system/general/refund/refund-audit
type RefundAuditCallback struct {
	TransactionOrder
	RefundAuditStatus int32  `json:"refund_audit_status"` // 审核状态，1-同意退款，2-不同意退款
	RefundId          string `json:"refund_id"`           // 交易系统侧退款单号，长度 <= 64 byte
	DenyMessage       string `json:"deny_message"`        // 不同意退款信息(不同意退款时必填)，长度 <= 512 byte>
}

// RefundAuditCallbackResponse 同步退款审核结果响应参数
type RefundAuditCallbackResponse struct {
	Error
	LogId string `json:"log_id"` // 日志id，排查问题时使用
}
