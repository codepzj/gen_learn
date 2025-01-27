package dao

import (
	"gen_learn/domain"
	"gen_learn/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func AutoMigrateTables(db *gorm.DB) {
	tables := []any{
		domain.User{},
		domain.Book{},
	}

	for _, t := range tables {
		_ = db.AutoMigrate(&t)
	}
}

func InitDB() {
	mysqlDSN := "root:pzj20162116@tcp(127.0.0.1:3306)/gen_learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlDSN,
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("数据库连接失败")
	}
	AutoMigrateTables(db)
	global.DB = db
}
