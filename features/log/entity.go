package log

import "time"

type LogCore struct {
	ID         uint
	Feedback   string
	MenteeID   uint
	Created_At time.Time
	UserID     uint
	Url_File   string
	Status     string
}

type UsecaseInterface interface {
	PostData(data LogCore) (row int, err error)
	GetLogById(id int) (data LogCore, err error)
}

type DataInterface interface {
	InsertData(data LogCore) (row int, err error)
	SelectLogById(id int) (data LogCore, err error)
}
