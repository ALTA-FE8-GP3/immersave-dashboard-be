package delivery

import "project/immersive-dashboard/features/log"

type LogRequest struct {
	Feedback string `json:"feedback" form:"feedback"`
	MenteeID uint   `json:"mentee_id" form:"mentee_id"`
	UserID   uint   `json:"user_id" form:"user_id"`
	Url_File string `json:"url_file" form:"url_file"`
	Status   string `json:"status" form:"status"`
}

func ToCore(dataRequest LogRequest) log.LogCore {
	return log.LogCore{
		Feedback: dataRequest.Feedback,
		MenteeID: dataRequest.MenteeID,
		UserID:   dataRequest.UserID,
		Url_File: dataRequest.Url_File,
		Status:   dataRequest.Status,
	}
}
