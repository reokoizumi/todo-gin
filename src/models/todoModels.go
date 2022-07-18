package models

import (
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetAllTodos(todo *[]Todo) (err error) {
	if err = Db.Find(&todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateATodo(todo *Todo) (err error) {
	if err = Db.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func GetATodo(todo *Todo, id uint) (err error) {
	if err := Db.Where("id = ?", id).First(&todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(todo *Todo, id string) (err error) {
	if err := Db.Where("id = ?", id).Save(&todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteATodo(todo *Todo, id uint) (err error) {
	if err := Db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return err
	}
	return nil
}
