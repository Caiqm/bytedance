# 抖音相关接口（简易版）
抖音担保交易、小程序登录、生成小程序二维码

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