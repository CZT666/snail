package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"snail/teacher_backend/settings"
)

var (
	DB *gorm.DB
)

func InitMySQL(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.PassWord, cfg.Host, cfg.Port, cfg.DB)
	fmt.Printf("dsn %v", dsn)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect mysql failed: %v", err.Error())
		return err
	}
	return DB.DB().Ping()
}

func Close() {
	DB.DB().Close()
}
