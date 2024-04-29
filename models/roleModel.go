package models

import (
	"bookingAPI/database"

	"gorm.io/gorm"
)

// Role model
type Role struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
}

// CreateRole Create a role
func CreateRole(Role *Role) (err error) {
	err = database.DB.Create(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// GetRoles Get all roles
func GetRoles(Role *[]Role) (err error) {
	err = database.DB.Find(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// GetRole Get role by id
func GetRole(Role *Role, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateRole Update role
func UpdateRole(Role *Role) (err error) {
	database.DB.Save(Role)
	return nil
}
