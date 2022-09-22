package delivery

import (
	"project/immersive-dashboard/features/log"
	"time"
)

type LogResponse struct {
	ID         uint      `json:"id" form:"id"`
	Feedback   string    `json:"feedback" form:"feedback"`
	MenteeID   uint      `json:"mentee_id" form:"mentee_id"`
	Created_At time.Time `json:"created_at" form:"created_at"`
	UserID     uint      `json:"user_id" form:"user_id"`
	Url_File   string    `json:"url_file" form:"url_file"`
	Status     string    `json:"status" form:"status"`
}

func FromCore(dataCore log.LogCore) LogResponse {
	return LogResponse{
		ID:         dataCore.ID,
		Feedback:   dataCore.Feedback,
		MenteeID:   dataCore.MenteeID,
		Created_At: dataCore.Created_At,
		UserID:     dataCore.UserID,
		Url_File:   dataCore.Url_File,
		Status:     dataCore.Status,
	}
}

func FromCoreList(dataCore []log.LogCore) []LogResponse {
	var dataResponse []LogResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse

}
