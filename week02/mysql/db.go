package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbinstance *gorm.DB
func getDB() *gorm.DB {
	if dbinstance == nil {
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: "golearn:golearn@tcp(172.24.10.228:3306)/golearn?charset=utf8mb4&parseTime=True&loc=Local",
		}))
		if err != nil {
			panic(err)
		}
		dbinstance = db
	}

	return dbinstance
}