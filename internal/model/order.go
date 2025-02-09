package model

import (
	"goframe-shop-v2/internal/model/do"
	"goframe-shop-v2/internal/model/entity"
)

type OrderListInput struct {
	Page           int    // 分页号码
	Size           int    // 分页数量，最大50
	Number         string `json:"number"           dc:"订单编号"`
	UserId         int    `json:"userId"           dc:"用户id"`
	PayType        int    `json:"payType"          dc:"支付方式 1微信 2支付宝 3云闪付"`
	PayAtGte       string `json:"payAtGte"         dc:"支付时间>="`
	PayAtLte       string `json:"payAtLte"         dc:"支付时间<="`
	Status         int    `json:"status"           dc:"订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价"`
	ConsigneePhone string `json:"consigneePhone"   dc:"收货人手机号"`
	PriceGte       int    `json:"priceGte"         dc:"订单金额>= 单位分"`
	PriceLte       int    `json:"priceLte"         dc:"订单金额<= 单位分"`
	DateGte        string `json:"dateGte"          dc:"創建时间>="`
	DateLte        string `json:"dateLte"          dc:"創建时间<="`
}

type OrderListOutput struct {
	List  []OrderListOutputItem
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}

type OrderListOutputItem struct {
	entity.OrderInfo
}

type OrderDetailInput struct {
	Id uint
}

type OrderDetailOutput struct {
	do.OrderInfo
	GoodsInfo []*do.OrderGoodsInfo `orm:"with:order_id=id"`
}
