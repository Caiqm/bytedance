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
	client.LoadOptionFunc(WithKeyVersion("1"), WithPrivateKey("MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDWOy0ZBwIEMVKl\n/sZnVE/hVYmZYW0YCzv8w771zODofvi3EHxbHS+6i/nO98vinI2PYFhN4uJDMXOh\nL+6avMmqjE0+IDV2W+Ceh0FBYPJzgFknIr13djP7spzCIxVXWqdPMEhDuaMWDBUK\nax4wPhzGlR/NIcm6RhqeL1ZcGAJSmQwni+fIC8efRPa4SAE6FElOxDqYusMcBKNT\nvR32iFkYPCbS9VZ+3tKkZU9MNdLI+6b1vRTbG3oxRVLR+qsZHyNMTceohmY5aNu2\nlJBLnFSN1Swb8/Ky8jSyYa9hXdA2cZVAwVs8AbRqJm4V5aXp8QcTuRYN/qj5zKtu\nuSoQ8SmPAgMBAAECggEARB6m9WlUC88/L4c2wU5+4/DWAw6GqPmRhosbon8WGPee\n0F8Om85JCfFOarBChXkwTiHdkk6Nkb4cvWlHXq6fFXrtaM7HES6f2ui5KDfSgEWU\nk6wpQN42I1elHtoXBCLQRo7cu6RAplcLLbfn7oEhl0d9yDVxNqjp99inhjsgKfDp\nQhdVD0b5FXHLYWtYTmMijEF8HSJTR/VIW5szGlQD6RCeD1XIB/P1dXHSovY+CFd9\nfweUCKwoRc+d8H/HH0KfvnskIgEw3nvVSz2zdzMvQNXzgj85S+MAMzddoHySZEoU\nLNpa7/ST1SU5gG/4lOczczMbOsywXhDgtH3FrU6BUQKBgQD5Ujj/9HqeLfJ6b5EZ\nl4/qjNcM4JqOf4EfOCpBWvHz3GlFtwCvfgXfzD6vOhmztd0sGU6MQMhRA5XF0/h/\nHkajZ9iLddsYC0WCfwLf6bpTF/eW9xZvZGXHgcqm9lzG2c9VRBZNE/XCkg+DXFJF\ni/dLYtEQ3Lns3O/muYl4P+omSQKBgQDb+FDKbx1vm4WBfXfZ3fUFcXlOHlp0SqRP\n7IEaCW95huauQFUeiUsUj46wB/n7fbsvXY4OJXI94OVMmDiliT2EKosYJAJ0TbN2\narKqJJlWHR5hQ3HxF+1jZXGj60QGvBFJB4JPRgoax60v6o+kRBUGfMaTf8S6A9wu\neWD0OfXxFwKBgQClB5vbMISJd3htOrQc8Im+g4NFbtZfRF6/v3mY0Q4ekOLN1piJ\nkS1qOdO0QQ0KO3Mu4YFwrODrXyggTnLKEcxIFVDs3xIrFCrqeg/5Dsjaf57POraK\n/TWxnWP6qFA4/6uRkQQB+RhPtka2eMGsbz09lSc1tiULeCMD4gPwqL3goQKBgQDL\neSDSgIIMESInq7XHlCboCBUbi7xEQh9Hxw+M1BmHop+To/KYsor+0+Q7NMWqOp2B\nrNqQf1bzoiq49T3A8fgzX18Wz4htqMpOyVfHRt+H5puJgOfPCkEOZnH+HMvqJuEe\nWpHRXopOR1IONrz0R/3i6FyPZ+rD3no3ixOCFnNHtQKBgQDfc2c4xaEF4T6FK1Ot\nemstoDAUkw4ANSlHCFGqepjs94IxFYHvmfjAjPf3LPUUgAcEE8OamP29lWdXO4pK\nCzW+y0fX/g6OvrkfuLWZ978rh2bBhLuT2B5CPiErquwAULuFj6IGg2DA0j5UT79E\nf6xPKURekD+ERkSec6GPsIFVnQ=="))
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
					"https://common-shop.oss-cn-shenzhen.aliyuncs.com/static/images/ba_shop_detail.png",
				},
				Type:       704,
				TagGroupId: "tag_group_7443548955339669558",
			},
		},
	}
	rsp, err := client.TradeTransactionPay(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", rsp)
}
