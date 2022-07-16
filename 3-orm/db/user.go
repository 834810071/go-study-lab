package db

import (
	"go-unit-test/3-orm/table"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var _DB *gorm.DB

func DB() *gorm.DB {
	return _DB
}

func initDB() *gorm.DB {
	// In our docker dev environment use
	//db, err := gorm.Open("mysql", "root:superpass@tcp(database:3306)/go_web?charset=utf8&parseTime=True&loc=Local")
	// Out of docker use
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(time.Second * 300)
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	return db
}

func init() {
	_DB = initDB()
}

func CreateUser(user *table.User) (err error) {
	err = DB().Create(user).Error

	return
}

func GetUserByNameAndPassword(name, password string) (user *table.User, err error) {
	user = new(table.User)
	err = DB().Where("username = ? AND secret = ?", name, password).
		First(&user).Error

	return
}

func UpdateUserNameById(userName string, userId int64) (err error) {
	user := new(table.User)
	updated := map[string]interface{}{
		"username": userName,
	}
	err = DB().Model(user).Where("id = ?", userId).Updates(updated).Error
	return
}

func GetAllUsers() (users []*table.User, err error) {
	err = DB().Find(&users).Error
	return
}
