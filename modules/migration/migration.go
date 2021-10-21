package migration

import (
	"go-hexagonal-auth/modules/admin"
	"go-hexagonal-auth/modules/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{}, &admin.Admin{})
}
