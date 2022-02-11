package dao

import (
	"Simple-login-demo/models"
	"fmt"
)

// 新建用户
func CreateUser(user *models.User) (err error) {
	err = MysqlDB.Create(&user).Error
	return
}

// 查询用户
func QueryUser(loginVo *models.Login) (user models.User, err error) {
	err = MysqlDB.Debug().Where(map[string]interface{}{
		"name":   loginVo.Username,
		"status": 0,
	}).Find(&user).Error
	if err != nil {
		fmt.Println("查询用户错误ERRO：", err)
		return
	}
	return
}

// 根据UserId查询用户
func QueryUserByUserId(userId string) (user models.User, err error) {
	err = MysqlDB.Debug().Where(map[string]interface{}{
		"user_id": userId,
		"status":  0,
	}).Find(&user).Error
	if err != nil {
		fmt.Println("查询用户错误ERRO：", err)
		return
	}
	return
}
