package dao

import (
	"XCloud/model"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfoDao interface {
	AddUser(UserInfo model.UserInfo) error
	DeleteUser(userName string) error
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:root@(127.0.0.1:3306)/xcoluddb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	// 自动迁移
	db.AutoMigrate(&model.UserInfo{})
}

func CreateUserInfoDao() UserInfoDao {
	var userInfo UserInfoDao
	userInfo = new(UserInfoDaoMysql)
	return userInfo
}

type UserInfoDaoMysql struct {
}

func (user *UserInfoDaoMysql) AddUser(UserInfo model.UserInfo) error {
	fmt.Println("123")
	return nil
}
func (user *UserInfoDaoMysql) DeleteUser(UserName string) error {
	fmt.Println(UserName)
	return nil
}
