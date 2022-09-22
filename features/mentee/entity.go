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
	Type             string
	Major            string
	Graduate         string
	Status           string
	ClassID          uint
	Nama_Class       string
	Created_At       time.Time
	Updated_At       time.Time
}

type UsecaseInterface interface {
	GetAllMentee() (data []MenteeCore, err error)
	PostData(data MenteeCore) (row int, err error)
	// PutData(data UserCore) (row int, err error)
	// DeleteUser(id int) (row int, err error)
}

type DataInterface interface {
	SelectAllMentee() (data []MenteeCore, err error)
	InsertData(data MenteeCore) (row int, err error)
	// 	UpdateUser(data UserCore) (row int, err error)
	// 	DeleteUser(id int) (row int, err error)
}
