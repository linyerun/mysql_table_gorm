package entity

import "gorm.io/gorm"

type Contract struct {
	gorm.Model
	ClientId      uint    // 客户ID
	UserId        uint    // 销售人员ID
	TotalAmount   float64 `gorm:"type:decimal(10,5)"` //涉及金额
	ContractPic   string  `gorm:"type:varchar(200)"`  //合同图片
	ContractState uint8   // 合同的状态(已经发货表示订单完成了好吧)
	// 0: 还没付款, 1: 已经付款, 2: 完成所有发货
}

type PurchasingList struct {
	gorm.Model
	ContractId uint   // 合同ID
	Address    string `gorm:"type:varchar(100)"` // 付款之后由用户决定地址
}

type PurchasingGoods struct {
	gorm.Model
	PurchasingListId uint  // 采购清单
	GoodsId          uint  // 商品ID
	PurchasingCount  uint  // 购买数量
	State            uint8 // 货物是否已被生成发货单
}
