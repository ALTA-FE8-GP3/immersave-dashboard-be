package delivery

import "time"

type LogResponse struct {
	ID         uint
	Feedback   string
	MenteeID   uint
	Created_At time.Time
	UserID     uint
	Url_File   string
	Status     string
}
