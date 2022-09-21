package log

import "time"

type LogCore struct {
	ID         uint
	Feedback   string
	Id_Mentee  uint
	Created_At time.Time
	Id_User    uint
	Url_File   string
	Status     string
}
