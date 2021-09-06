package models

import (
	"EasyGo-Server/pkg/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB 
	err error 
)

func Init() {
	DB, err = gorm.Open(
		"mysql", 
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.Mysql.Username,
			config.Mysql.Password,
			config.Mysql.Host,
			config.Mysql.Port,
			config.Mysql.Database,
		),
	)
	if err != nil {
		panic(err)
	}
	// set up conn pool 
	DB.DB().SetMaxIdleConns(10) 
	DB.DB().SetMaxOpenConns(100)
}