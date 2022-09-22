package data

import (
	class "project/immersive-dashboard/features/class/data"
	"project/immersive-dashboard/features/log"
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

func fromCore(dataCore log.LogCore) Log {
	return Log{
		Feedback: dataCore.Feedback,
		MenteeID: dataCore.MenteeID,
		UserID:   dataCore.UserID,
		Url_file: dataCore.Url_File,
		Status:   dataCore.Status,
	}
}

func (dataLog *Log) toCore() log.LogCore {
	return log.LogCore{
		ID:       dataLog.ID,
		Feedback: dataLog.Feedback,
		MenteeID: dataLog.MenteeID,
		UserID:   dataLog.UserID,
		Url_File: dataLog.Url_file,
		Status:   dataLog.Status,
	}
}

func toCoreList(dataLog []Log) []log.LogCore {
	var dataCore []log.LogCore

	for key := range dataLog {
		dataCore = append(dataCore, dataLog[key].toCore())
	}
	return dataCore
}
