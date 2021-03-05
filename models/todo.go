package models

import "bubble/dao"

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
 Todo的增删改查
*/

// InitModel 创建表
func InitModel() {
	dao.DB.AutoMigrate(Todo{})
}

// CreateATodo 在数据库创建一个todo
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

// GetAllTodo 在数据库中查找所有todo
func GetAllTodo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return
	}
	return
}

// GetATodo 在数据库中查找一个todo
func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return
	}
	return
}

// UpdateATodo 在数据库中更新一个todo
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

// DeleteATodo 从数据库中删除一个todo
func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(Todo{}).Error
	return
}
