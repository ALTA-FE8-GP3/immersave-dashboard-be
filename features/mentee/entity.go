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
	Id_Class         uint
	Created_At       time.Time
	Updated_At       time.Time
}
