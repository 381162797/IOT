package models

import (
	"mserver/setting"

	"github.com/jinzhu/gorm"
	// sqlite3
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB pool
var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open(setting.Config.DBType, setting.Config.DBURL)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	//建表及管理员账户
	if !db.HasTable(&User{}) {
		db.Set("gorm:table_options", "ENGINE = InnoDB DEFAULT CHARSET = utf8").
			CreateTable(&User{}, &Device{}, &Slaver{}, &Template{},
				&DataPoint{}, &DeviceTask{}, &Task{}, &PointData{})
		db.Create(&User{
			Username: "381162797",
			Password: "123456",
		})
	}

}
