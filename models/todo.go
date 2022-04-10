package models

import (
	"bubble/dao"
)

// Todo Model
type Todo struct {
	ID     int64  `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Status bool   `json:"status,omitempty"`
}

/*
 Todo这个model的增删改查操作都放在这里
*/

// CreateATdo 创建todo
func CreateATdo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return

}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=", id).Delete(&Todo{}).Error
	return

}
