package task03

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

/**
 * 初始化数据库
 */
func GormInitDb(dst ...interface{}) *gorm.DB {
	dsn := "root:fhj5532899@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(dst...)
	return db
}

func InitDb() *sql.DB {
	dsn := "root:fhj5532899@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	return db
}

func SqlXInitDb() *sqlx.DB {
	dsn := "root:fhj5532899@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	return db
}
