package log

import "time"

type LogCore struct {
	ID          uint
	Feedback    string
	MenteeID    uint
	Created_At  time.Time
	UserID      uint
	Url_File    string
	Status      string
	Nama_User   string
	Nama_Mentee string
}

type UsecaseInterface interface {
	PostData(data LogCore) (row int, err error)
	GetLogById(id int) (data LogCore, err error)
	GetAlllog() (data []LogCore, err error)
}

type DataInterface interface {
	SelectAlllog() (data []LogCore, err error)
	InsertData(data LogCore) (row int, err error)
	SelectLogById(id int) (data LogCore, err error)
}
