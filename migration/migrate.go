package migration

import (
	classModel "project/immersive-dashboard/features/class/data"
	userModel "project/immersive-dashboard/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&classModel.Class{})
}
