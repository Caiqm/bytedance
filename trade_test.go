package bytedance

import (
	"testing"
)

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

// 订单查询
func TestClient_TradeOrderQuery(t *testing.T) {
	t.Log("========== TradeOrderQuery ==========")
	client.LoadOptionFunc(WithApiHost("https://developer.toutiao.com/api/apps/ecpay/v1/query_order"))
	var p = TradeOrderQuery{}
	p.OutOrderNo = "TEST2023112717521212345678"
	rsp, err := client.TradeOrderQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", rsp)
}

// 发起退款
func TestClient_TradeRefund(t *testing.T) {
	t.Log("========== TradeRefund ==========")
	client.LoadOptionFunc(WithApiHost("https://developer.toutiao.com/api/apps/ecpay/v1/create_refund"))
	var p = TradeRefund{}
	p.OutOrderNo = "TEST2023112717521212345678"
	p.OutRefundNo = "TEST2023112717521212345698"
	p.Reason = "测试退款"
	p.RefundAmount = 1
	rsp, err := client.TradeRefund(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", rsp)
}

// 退款结果查询
func TestClient_TradeRefundQuery(t *testing.T) {
	t.Log("========== TradeRefundQuery ==========")
	client.LoadOptionFunc(WithApiHost("https://developer.toutiao.com/api/apps/ecpay/v1/query_refund"))
	var p = TradeRefundQuery{}
	p.OutRefundNo = "TEST2023112717521212345678"
	rsp, err := client.TradeRefundQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", rsp)
}

// 通用交易系统
func TestClient_TradeTransaction(t *testing.T) {
	t.Log("========== TradeTransaction ==========")
	client.LoadOptionFunc(WithKeyVersion(""), WithPrivateKey(""))
	var p = TradeTransaction{
		OutOrderNo:  "TEST2023112717521212345678",
		TotalAmount: 1,
		OrderEntrySchema: OrderEntrySchema{
			Path:   "pages/order/detail",
			Params: "{\"order_id\":1234}",
		},
		SkuList: []SkuList{
			{
				SkuId:    "TEST2023112717521212345678",
				Title:    "支付测试",
				Price:    1,
				Quantity: 1,
				ImageList: []string{
					"",
				},
				Type:       1,
				TagGroupId: "",
			},
		},
	}
	rsp, err := client.TradeTransactionPay(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", rsp)
}
