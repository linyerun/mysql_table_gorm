package entity

import "gorm.io/gorm"

type DispatchBill struct { // 发货单
	gorm.Model
	ContractId    uint   // 合同ID
	GoodsId       uint   // 商品ID
	GoodsCount    uint   // 商品数量
	CourierNumber string `gorm:"type:varchar(200)"` //快递单号,由仓库管理员填写
}
