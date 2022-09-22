package data

import (
	classModel "project/immersive-dashboard/features/class/data"
	"project/immersive-dashboard/features/mentee"

	"gorm.io/gorm"
)

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
	Class            classModel.Class
	Logs             []Log
}

type Log struct {
	gorm.Model
	Feedback string
	MenteeID uint
	UserID   uint
	Url_file string
	Status   string
	Mentee   Mentee
}

func fromCore(dataCore mentee.MenteeCore) Mentee {
	return Mentee{
		Nama_Mentee:      dataCore.Nama_Mentee,
		Address:          dataCore.Address,
		Home_Address:     dataCore.Home_Address,
		Email:            dataCore.Email,
		Gender:           dataCore.Gender,
		Telegram:         dataCore.Telegram,
		Phone:            dataCore.Phone,
		Discord:          dataCore.Discord,
		Nama_Emergency:   dataCore.Nama_Emergency,
		Phone_Emergency:  dataCore.Phone_Emergency,
		Status_Emergency: dataCore.Status_Emergency,
		Type:             dataCore.Type,
		Major:            dataCore.Major,
		Graduate:         dataCore.Graduate,
		Status:           dataCore.Status,
		ClassID:          dataCore.ClassID,
	}
}

func (dataMentee *Mentee) toCore() mentee.MenteeCore {
	return mentee.MenteeCore{
		ID:               dataMentee.ID,
		Nama_Mentee:      dataMentee.Nama_Mentee,
		Address:          dataMentee.Address,
		Home_Address:     dataMentee.Home_Address,
		Email:            dataMentee.Email,
		Gender:           dataMentee.Gender,
		Telegram:         dataMentee.Telegram,
		Phone:            dataMentee.Phone,
		Discord:          dataMentee.Discord,
		Nama_Emergency:   dataMentee.Nama_Emergency,
		Phone_Emergency:  dataMentee.Phone_Emergency,
		Status_Emergency: dataMentee.Status_Emergency,
		Type:             dataMentee.Type,
		Major:            dataMentee.Major,
		Graduate:         dataMentee.Graduate,
		Status:           dataMentee.Status,
		ClassID:          dataMentee.ClassID,
	}
}

func toCoreList(dataMentee []Mentee) []mentee.MenteeCore {
	var dataCore []mentee.MenteeCore

	for key := range dataMentee {
		dataCore = append(dataCore, dataMentee[key].toCore())
	}
	return dataCore
}
