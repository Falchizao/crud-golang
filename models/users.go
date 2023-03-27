package models

import "gorm.io/gorm"

type Users struct {
	ID    uint    `gorm:"primary key;autoIncrement" json:"id"`
	Nome  *string `json:"nome"`
	Senha *string `json:"senha"`
	Idade *int    `json:"idade"`
	Email *string `json:"email"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err
}
