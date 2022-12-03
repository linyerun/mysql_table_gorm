package entity

import "gorm.io/gorm"

type PurchaseOrder struct { // 进货单
	gorm.Model
	GoodsId    uint  // 商品ID
	GoodsCount uint  // 商品数量
	State      uint8 // 进货单状态 (0是刚创建, 1是完成进货)
}

type PurchaseOrderInfo struct {
	gorm.Model
	PurchaseOrderId uint    // 进货单ID
	UnitPrice       float64 `gorm:"type:decimal(10,5)"` // 进货单价
	CommentInfo     string  `gorm:"type:varchar(500)"`
}
