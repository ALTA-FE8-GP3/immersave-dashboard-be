package mentee

import (
	"time"
)

type MenteeCore struct {
	ID               uint
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
	Category         string
	Major            string
	Graduate         string
	Status           string
	ClassID          uint
	Nama_Class       string
	Created_At       time.Time
	Updated_At       time.Time
}

type UsecaseInterface interface {
	GetAllMentee(class_id int, category, status string) (data []MenteeCore, err error)
	PostData(data MenteeCore) (row int, err error)
	UpdateMenteeId(data MenteeCore) (row int, err error)
	Delete(id int) (row int, err error)
	GetMenteeById(id int) (data MenteeCore, err error)
}

type DataInterface interface {
	SelectAllMentee(class_id int, category, status string) (data []MenteeCore, err error)
	InsertData(data MenteeCore) (row int, err error)
	UpdateMentee(data MenteeCore) (row int, err error)
	DeleteMentee(id int) (row int, err error)
	SelectMenteeById(id int) (data MenteeCore, err error)
}
