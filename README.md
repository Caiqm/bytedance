# 抖音相关接口（简易版）
抖音担保交易、小程序登录、生成小程序二维码、订单同步

## 安装

#### 启用 Go module

```go
go get github.com/Caiqm/bytedance
```

```go
import "github.com/Caiqm/bytedance"
```

#### 未启用 Go module

```go
go get github.com/Caiqm/bytedance
```

```go
import "github.com/Caiqm/bytedance"
```

## 如何使用

```go
var client, err = bytedance.New(appID, Secret)
```

#### 关于密钥（Secret）
是应用的唯一凭证密钥，可以在开发者后台获取

## 加载配置

```go
// 加载支付系统秘钥SALT
client.LoadOptionFunc(WithSalt(SALT))

// 或者一开始就加载支付系统秘钥SALT
var client, err = bytedance.New(appID, Secret, WithSalt(SALT))
```

#### 关于支持的配置

```go
// 设置请求链接，可自定义请求接口，传入host字符串
WithApiHost(HOST)

// 设置支付系统秘钥SALT
WithSalt(SALT)

// 设置支付系统回调TOKEN
WithToken(TOKEN)

// 也可自定义传入配置，返回以下类型即可
type OptionFunc func(c *Client)
```

## 小程序登录

```go
// 小程序登录
func TestClient_Code2Session(t *testing.T) {
	t.Log("========== Code2Session ==========")
	client.LoadOptionFunc(WithApiHost("https://open-sandbox.douyin.com/api/apps/v2/jscode2session"))
	var p = Code2Session{}
	p.Code = "647f16afe0b44c49a8eb1cb3c02aXX31"
	rsp, err := client.Code2Session(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rsp)
}
```

## 小程序支付

```go
// 担保支付
func TestClient_TradeEcPay(t *testing.T) {
	t.Log("========== TradeEcPay ==========")
	client.LoadOptionFunc(WithApiHost("https://developer.toutiao.com/api/apps/ecpay/v1/create_order"))
	var p = TradeEcPay{}
	p.OutOrderNo = "TEST2023112717521212345678"
	p.TotalAmount = 1
	p.Subject = "支付测试"
	p.Body = "支付测试"
	p.ValidTime = 300
	rsp, err := client.TradeEcPay(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", rsp)
}
```

## 小程序订单同步

```go
// 同步订单
func TestClient_OrderPush(t *testing.T) {
	t.Log("========== OrderPush ==========")
	client.LoadOptionFunc(WithApiHost("https://developer.toutiao.com/api/apps/order/v2/push"))
	// 小程序订单商品类型
	var itemLists []ItemListApplet
	var i ItemListApplet
	i.ItemCode = "111" // 开发者侧商品 ID
	i.Img = ""
	i.Title = "麻花零食"
	i.SubTitle = "开心麻花"
	i.Amount = 1 // 单类商品的数目
	i.Price = 1  // 单类商品的总价，单位为分
	// 添加到数组
	itemLists = append(itemLists, i)
	// 根据类型获取那种类型的POI订单
	var o OrderDetailPoiApplet
	o.OrderId = "" // 开发者侧业务单号。用作幂等控制。该订单号是和担保支付的支付单号绑定的，也就是预下单时传入的 out_order_no 字段
	o.CreateTime = 1694761323
	o.Status = "待支付"
	o.Amount = 1                            // 订单商品总数
	o.TotalPrice = 1                        // 订单总价，单位为分
	o.DetailUrl = "pages/order/orderDetail" // 小程序订单详情页 path
	o.ItemList = itemLists
	// 转化为json
	orderDetailJson, _ := json.Marshal(o)
	// 主要参数
	var p OrderPush
	p.AccessToken = ""
	p.AppName = "douyin"
	p.OpenId = ""
	p.OrderStatus = 0
	p.OrderType = 0
	p.UpdateTime = 1694761323
	p.Extra = "额外参数"
	p.OrderDetail = string(orderDetailJson)
	// 主方法
	r, err := client.OrderPush(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", r)
}
```