package data

import (
	mentee "project/immersive-dashboard/features/mentee/data"
	"project/immersive-dashboard/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama_User string
	Email     string
	Password  string
	Team      string
	Role      string
	Status    string
	Classes   []Class
	Logs      []Log
}

type Class struct {
	gorm.Model
	Nama_Class string
	UserID     uint
	User       User
}

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

func fromCore(dataCore user.UserCore) User {
	return User{
		Nama_User: dataCore.Nama_User,
		Email:     dataCore.Email,
		Password:  dataCore.Password,
		Team:      dataCore.Team,
		Role:      dataCore.Role,
		Status:    dataCore.Status,
	}
}

func (dataUser *User) toCore() user.UserCore {
	return user.UserCore{
		ID:        dataUser.ID,
		Nama_User: dataUser.Nama_User,
		Email:     dataUser.Email,
		Password:  dataUser.Password,
		Team:      dataUser.Team,
		Role:      dataUser.Role,
		Status:    dataUser.Status,
	}
}

func toCoreList(dataUser []User) []user.UserCore {
	var dataCore []user.UserCore

	for key := range dataUser {
		dataCore = append(dataCore, dataUser[key].toCore())
	}
	return dataCore
}
