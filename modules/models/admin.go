package models

import "time"

type Admin struct {
	ID         int       `gorm:"id;primaryKey;autoIncrement"`
	Name       string    `gorm:"name"`
	Username   string    `gorm:"email"`
	Password   string    `gorm:"password"`
	CreatedAt  time.Time `gorm:"created_at"`
	CreatedBy  string    `gorm:"created_by"`
	ModifiedAt time.Time `gorm:"modified_at"`
	ModifiedBy string    `gorm:"modified_by"`
}