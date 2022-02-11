package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

type Login struct {
	Username string `json:"username" binding:"required,eqfield=Username"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	jwt.StandardClaims `gorm:"-"`
	Name               string `json:"name"`
	Password           string `json:"password"`
	Status             int    `json:"status"`
	Org                string `json:"org"`
	Mobile             string `json:"mobile"`
	UserId             string `json:"user_id"`
	Code               string `json:"code"`
}

/*
	Todo这个Model的增删改查操作都放在这里
*/
// CreateATodo 创建todo
//func CreateATodo(todo *Todo) (err error){
//	err = dao.DB.Create(&todo).Error
//	return
//}
//
//func GetAllTodo() (todoList []*Todo, err error){
//	if err = dao.DB.Find(&todoList).Error; err != nil{
//		return nil, err
//	}
//	return
//}
//
//func GetATodo(id string)(todo *Todo, err error){
//	todo = new(Todo)
//	if err = dao.DB.Debug().Where("id=?", id).First(todo).Error; err!=nil{
//		return nil, err
//	}
//	return
//}
//
//func UpdateATodo(todo *Todo)(err error){
//	err = dao.DB.Save(todo).Error
//	return
//}
//
//func DeleteATodo(id string)(err error){
//	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
//	return
//}
