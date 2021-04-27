package dao

import (
	"XCloud/model"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfoDao interface {
	AddUser(UserInfo model.UserInfo) error
	DeleteUser(userName string) error
}

var db *gorm.DB

func Connect(ip, dbName, password string) {
	var err error
	db, err = gorm.Open("mysql", dbName+":"+password+"@("+ip+")/xcoluddb?charset=utf8mb4&parseTime=True&loc=Local")
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

func UserAuthentication(userName, password string) (isPass bool, err error) {
	if "" == userName {
		err = errors.New("userName:nil")
		return false, err
	}
	if "" == password {
		err = errors.New("password:nil")
		return false, err
	}

	var user model.UserInfo
	db.Where("UserName = ?", userName).First(&user)

	if user.Password != password {
		return false, nil
	}
	return true, nil
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
