package data

import (
	class "project/immersive-dashboard/features/class/data"
	mentee "project/immersive-dashboard/features/mentee/data"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Feedback string
	MenteeID uint
	UserID   uint
	Url_file string
	Status   string
	Mentee   mentee.Mentee
	User     User
}

type User struct {
	gorm.Model
	Nama_User string
	Email     string
	Password  string
	Team      string
	Role      string
	Status    string
	Classes   []class.Class
	Logs      []Log
}
