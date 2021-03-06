package models

import (
	"Bubble/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// Todo 增删改查

// CreateTodo 创建 ToDo
func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err := dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetTodo(id string) (todo *Todo, err error) {
	//todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
