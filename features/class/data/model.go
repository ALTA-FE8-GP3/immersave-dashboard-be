package data

import (
	"project/immersive-dashboard/features/class"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Nama_Class string
	UserID     uint
	Mentees    []Mentee
	User       User
}

type Mentee struct {
	gorm.Model
	Nama_Mentee      string
	Address          string
	Home_Address     string
	Email            string
	Gender           string
	Telegram         string
	Phone            uint
	Discord          string
	Nama_Emergency   string
	Phone_Emergency  uint
	Status_Emergency string
	Type             string
	Major            string
	Graduate         string
	Status           string
	ClassID          uint
	Class            Class
}

type User struct {
	gorm.Model
	Nama_User string
	Email     string
	Password  string
	Team      string
	Role      string
	Status    string
	Classes   []Class
}

func fromCore(dataCore class.ClassCore) Class {
	return Class{
		Nama_Class: dataCore.Nama_Class,
	}
}

func (dataClass *Class) toCore() class.ClassCore {
	return class.ClassCore{
		ID:         dataClass.ID,
		Nama_Class: dataClass.Nama_Class,
	}
}

func toCoreList(dataClass []Class) []class.ClassCore {
	var dataCore []class.ClassCore

	for key := range dataClass {
		dataCore = append(dataCore, dataClass[key].toCore())
	}
	return dataCore
}
