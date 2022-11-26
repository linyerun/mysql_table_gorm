package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"mysql_table_gorm/entity"
)

func main() {
	db := getDB()

	// user.go
	AutoMigrate(db, entity.User{})

	// client.go
	AutoMigrate(db, entity.Client{})

	// contract.go
	AutoMigrate(db, entity.Contract{})
	AutoMigrate(db, entity.PurchasingList{})
	AutoMigrate(db, entity.PurchasingGoods{})

	// dispatch_bill.go
	AutoMigrate(db, entity.DispatchBill{})

	// purchase_order.go
	AutoMigrate(db, entity.PurchaseOrder{})

	// goods.go
	AutoMigrate(db, entity.Goods{})
}

func getDB() (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/company_sales_management_system?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	return db
}

func AutoMigrate(db *gorm.DB, obj any) {
	err := db.AutoMigrate(obj)
	if err != nil {
		panic(err)
	}
}
