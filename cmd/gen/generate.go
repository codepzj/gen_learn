package main

import (
	"gen_learn/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const mysqlDSN = "root:pzj20162116@tcp(127.0.0.1:3306)/gen_learn?charset=utf8mb4&parseTime=True"

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("数据库连接失败")
	}
	return db
}

func main() {
	db := connectDB(mysqlDSN)
	g := gen.NewGenerator(gen.Config{
		OutPath:           "dao/query",
		FieldNullable:     true,  // 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldCoverable:    false, // 插入零值被 gorm 识别并赋予默认值覆盖
		FieldSignable:     false,
		FieldWithIndexTag: true, // 生成 gorm 标签的字段索引属性
		FieldWithTypeTag:  true, // 生成 gorm 标签的字段类型属性
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
	})
	g.UseDB(db)
	// 统一数字类型为int64,兼容protobuf
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
	}
	g.WithDataTypeMap(dataMap)
	g.ApplyBasic(domain.User{}, domain.Book{})
	g.Execute()
}
