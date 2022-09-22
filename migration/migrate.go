package migration

import (
	classModel "project/immersive-dashboard/features/class/data"
	logModel "project/immersive-dashboard/features/log/data"
	menteeModel "project/immersive-dashboard/features/mentee/data"
	userModel "project/immersive-dashboard/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&classModel.Class{})
	db.AutoMigrate(&menteeModel.Mentee{})
	db.AutoMigrate(&logModel.Log{})
}
