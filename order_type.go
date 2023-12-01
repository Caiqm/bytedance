package bytedance

// OrderPush 订单同步 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/order/order-sync
type OrderPush struct {
	AuxParam

	ClientKey   string `json:"client_key,omitempty"`  // 第三方在抖音开放平台申请的 ClientKey 注意：POI 订单必传
	AccessToken string `json:"access_token"`          // 服务端 API 调用标识，通过 getAccessToken 获取
	ExtShopId   string `json:"ext_shop_id,omitempty"` // POI 店铺同步时使用的开发者侧店铺 ID，购买店铺 ID，长度 < 256 byte 注意：POI 订单必传
	AppName     string `json:"app_name"`              // 做订单展示的字节系 app 名称，目前为固定值“douyin”
	OpenId      string `json:"open_id"`               // 小程序用户的 open_id，通过 code2Session 获取
	UpdateTime  int    `json:"update_time"`           // 订单信息变更时间，10 位秒级时间戳，update_time每次状态变更推送时需要比上次推送的值大，否则可能忽略该次状态推送。例如：某次推送订单时的update_time为1694761323，则下次推送该订单时，update_time至少为1694761324。
	OrderDetail string `json:"order_detail"`          // json string，根据不同订单类型有不同的结构体，请参见 order_detail 字段说明（json string）
	OrderType   int    `json:"order_type"`            // 订单类型，枚举值: 0：普通小程序订单（非POI订单） 9101：团购券订单（POI 订单） 9001：景区门票订单（POI订单）
	OrderStatus int    `json:"order_status"`          // 普通小程序订单订单状态，POI 订单可以忽略 0：待支付 1：已支付 2：已取消（用户主动取消或者超时未支付导致的关单） 4：已核销（核销状态是整单核销,即一笔订单买了 3 个券，核销是指 3 个券核销的整单） 5：退款中 6：已退款 8：退款失败 注意：普通小程序订单必传，担保支付分账依赖该状态
	Extra       string `json:"extra"`                 // 自定义字段，用于关联具体业务场景下的特殊参数，长度 < 2048byte
}

func (aux OrderPush) NeedSign() bool {
	return false
}

func (aux OrderPush) NeedAppId() bool {
	return false
}

// 普通小程序订单参数（order_detail 字段说明）
type OrderDetailPoiApplet struct {
	OrderId    string           `json:"order_id"`    // 开发者侧业务单号。用作幂等控制。该订单号是和担保支付的支付单号绑定的，也就是预下单时传入的 out_order_no 字段，长度 <= 64byte
	Status     string           `json:"status"`      // 订单状态，建议采用以下枚举值： 待支付 已支付 已取消 已超时 已核销 退款中 已退款 退款失败
	CreateTime int64            `json:"create_time"` // 订单创建的时间，13 位毫秒时间戳
	Amount     int64            `json:"amount"`      // 订单商品总数
	TotalPrice int64            `json:"total_price"` // 订单总价，单位为分
	DetailUrl  string           `json:"detail_url"`  // 小程序订单详情页 path，长度<=1024 byte (备注：该路径需要保证在小程序内配置过，相对路径即可）
	ItemList   []ItemListApplet `json:"item_list"`   // 子订单商品列表，不可为空
}

// 普通小程序订单item_list字段说明
type ItemListApplet struct {
	ItemCode string `json:"item_code"` // 开发者侧商品 ID，长度 <= 64 byte
	Img      string `json:"img"`       // 子订单商品图片 URL， 长度 <= 512 byte
	Title    string `json:"title"`     // 子订单商品介绍标题，长度 <= 256 byte
	SubTitle string `json:"sub_title"` // 子订单商品介绍副标题，长度 <= 256 byte
	Amount   int64  `json:"amount"`    // 单类商品的数目
	Price    int64  `json:"price"`     // 单类商品的总价，单位为分
}

// 团购券类型订单参数（order_detail 字段说明）
type OrderDetailPoiGroup struct {
	ExtOrderId        string `json:"ext_order_id"`                // 开发者系统侧业务单号。用作幂等控制。该订单号是和担保支付的支付单号绑定的，即预下单时传入的 out_order_no 字段，长度 <= 64byte
	Status            int64  `json:"status"`                      // 枚举值： 10：已取消（抖音订单中心可看到，状态为"已取消"） 110：待支付 310：未使用 340：已使用 410：退款中 420： 退款成功 430： 退款失败
	ShopName          string `json:"shop_name"`                   // 商铺名字，长度 <= 256 byte
	EntryType         int64  `json:"entry_type"`                  // 订单详情页的外链跳转类型，通过该接口上传的都为 2， 1：H5 2：抖音小程序
	EntrySchema       string `json:"entry_schema"`                // 订单详情页的外链跳转 schema 参数，格式为 json 字符串。长度 <= 512byte
	CreateOrderTime   int64  `json:"create_order_time"`           // 下单时间（13位毫秒时间戳）
	Description       string `json:"description,omitempty"`       // 订单描述，长度<=500 byte
	TotalPrice        int64  `json:"total_price"`                 // 订单总金额（单位：分）
	PayTime           int64  `json:"pay_time,omitempty"`          // 支付时间（13位毫秒时间戳），未付款时不用传。
	ExtValidShopId    string `json:"ext_valid_shop_id,omitempty"` // 开发者侧卡劵核销门店ID，未核销时不用传，长度 <= 256 byte
	ValidPoiIdStr     string `json:"valid_poi_id_str,omitempty"`  // 开发者侧卡劵核销门店对应的抖音poiId，ext_valid_shop_id未匹配抖音POI时不用传，长度<= 128 byte
	ExtGoodsId        string `json:"ext_goods_id"`                // 开发者侧商品ID，长度<= 64 byte 备注：如果该商品没有接入抖音商品库，该字段为空
	GoodsName         string `json:"goods_name"`                  // 商品名称，长度 <= 256 byte
	GoodsInfo         string `json:"goods_info,omitempty"`        // 商品描述信息。向用户介绍商品，长度 <= 120byte。
	GoodsCoverImage   string `json:"goods_cover_image"`           // 商品图片，完整的url地址 长度 <= 512 byte
	GoodsEntryType    string `json:"goods_entry_type"`            // 商品详情页的外链跳转类型, 通过该接口上传的都为2, 1: H5 2: 抖音小程序
	GoodsEntrySchema  string `json:"goods_entry_schema"`          // 商品详情页的外链跳转schema参数，格式为 JSON 字符串，长度 <= 512 byte
	StartValidTime    string `json:"start_valid_time"`            // 生效时间，yyyy-MM-dd HH:mm:ss 格式字符串，24 小时制
	EndValidTime      string `json:"end_valid_time"`              // 失效时间，yyyy-MM-dd HH:mm:ss格式字符串，24小时制
	TicketNum         int64  `json:"ticket_num"`                  // 用户购买团购券的数量
	ExtTicketIds      string `json:"ext_ticket_ids"`              // 开发者侧券 ID，该信息用于用户可以明确的感知是哪一张券。格式为 JSON 数组字符串，每个 ID 长度 <= 64byte，例如：["123", "abc"]
	TicketDescription string `json:"ticket_description"`          // 券的使用说明。JSON 数组字符串，最多可以有10条，每条长度 <= 50byte。必须写明券的使用条件、领取条件、退款规则。例如：["1、本券不可兑换现金，不可找零。","2、每个用户最多可以领取1张。","3、如果订单发生退款，优惠券无法退还。"]
}

// 门票类型订单参数（order_detail 字段说明）
type OrderDetailPoiTicket struct {
	ExtOrderId        string `json:"ext_order_id"`               // 开发者系统侧业务单号。用作幂等控制。该订单号是和担保支付的支付单号绑定的，即预下单时传入的 out_order_no 字段，长度 <= 64byte
	Status            int64  `json:"status"`                     // 枚举值，如下： 10：已取消（抖音订单中心可看到，状态为"已取消"） 110：待支付 210：待确认 340：预订成功 410：退款中 420：退款成功 430：退款失败
	ShopName          string `json:"shop_name"`                  // 商铺名字，长度 <= 256 byte
	EntryType         int64  `json:"entry_type"`                 // 订单详情页的外链跳转类型，通过该接口上传的都为 2， 1：H5 2：抖音小程序
	EntrySchema       string `json:"entry_schema"`               // 订单详情页的外链跳转 schema 参数，格式为 json 字符串。长度 <= 512byte
	CreateOrderTime   int64  `json:"create_order_time"`          // 下单时间（13位毫秒时间戳）
	Description       string `json:"description,omitempty"`      // 订单描述，长度<=500 byte
	TotalPrice        int64  `json:"total_price"`                // 订单总金额（单位：分）
	PayTime           int64  `json:"pay_time,omitempty"`         // 支付时间（13位毫秒时间戳），未付款时不用传。
	ExtGoodsId        string `json:"ext_goods_id,omitempty"`     // 开发者侧商品ID，长度<= 64 byte 备注：如果该商品没有接入抖音商品库，该字段为空
	GoodsName         string `json:"goods_name"`                 // 商品名称，长度 <= 256 byte
	GoodsInfo         string `json:"goods_info,omitempty"`       // 商品描述信息。向用户介绍商品，长度 <= 120byte。
	GoodsCoverImage   string `json:"goods_cover_image"`          // 商品图片，完整的url地址 长度 <= 512 byte
	GoodsEntryType    string `json:"goods_entry_type"`           // 商品详情页的外链跳转类型, 通过该接口上传的都为2, 1: H5 2: 抖音小程序
	ValidPoiIdStr     string `json:"valid_poi_id_str,omitempty"` // 开发者侧卡劵核销门店对应的抖音poiId，ext_valid_shop_id未匹配抖音POI时不用传，长度<= 128 byte
	GoodsEntrySchema  string `json:"goods_entry_schema"`         // 商品详情页的外链跳转schema参数，格式为 JSON 字符串，长度 <= 512 byte
	StartValidTime    string `json:"start_valid_time"`           // 生效时间，yyyy-MM-dd HH:mm:ss 格式字符串，24 小时制
	EndValidTime      string `json:"end_valid_time"`             // 失效时间，yyyy-MM-dd HH:mm:ss格式字符串，24小时制
	TicketNum         int64  `json:"ticket_num"`                 // 用户购买团购券的数量
	ExtTicketIds      string `json:"ext_ticket_ids"`             // 开发者侧券 ID，该信息用于用户可以明确的感知是哪一张券。格式为 JSON 数组字符串，每个 ID 长度 <= 64byte，例如：["123", "abc"]
	TicketDescription string `json:"ticket_description"`         // 券的使用说明。JSON 数组字符串，最多可以有10条，每条长度 <= 50byte。必须写明券的使用条件、领取条件、退款规则。例如：["1、本券不可兑换现金，不可找零。","2、每个用户最多可以领取1张。","3、如果订单发生退款，优惠券无法退还。"]
}

// OrderPushRsp 订单同步响应参数
type OrderPushRsp struct {
	Err
	Body string `json:"body"` // POI 等关联业务推送结果，非 POI 订单为空，JSON 字符串
}
