package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init(conn string, poolsize int) (err error) {
	DB, err = gorm.Open("mysql", conn)
	if err != nil {
		fmt.Println("gorm.Open mysql err", err)
		return err
	}
	DB.DB().SetMaxOpenConns(poolsize)
	DB.DB().SetMaxIdleConns(poolsize)
	migration()
	return
}
