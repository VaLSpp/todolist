
package model

import "toDoList/db"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"` // false: not finished, true: finished
}

func InitModel() {
	db.DB.AutoMigrate(&Todo{})
}

// CRUD
// create
func CreateTodo(todo *Todo) error {
	return db.DB.Create(&todo).Error
}

// return all todo item
func GetTodoList(toDoList *[]Todo) error {
	return db.DB.Find(toDoList).Error
}

// delete
func DeleteTodo(id string) error {
	return db.DB.Where("id=?", id).Delete(Todo{}).Error
}

// update
func UpdateTodo(todo *Todo) error {
	return db.DB.Save(&todo).Error
}

// search
// search when need to update, need give todo obj in db
func SearchTodo1(id string, todo *Todo) error {
	return db.DB.Where("id=?", id).First(&todo).Error
}

// search when need to delete, only need to judge whether id is in db
func SearchTodo2(id string) error {
	return db.DB.Where("id=?", id).Error
}
