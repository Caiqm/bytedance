package bytedance

import (
	"github.com/goccy/go-json"
	"testing"
)

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
