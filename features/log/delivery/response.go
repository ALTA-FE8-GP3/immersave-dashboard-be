package delivery

import (
	"project/immersive-dashboard/features/log"
	"time"
)

type LogResponse struct {
	ID          uint      `json:"id" form:"id"`
	Feedback    string    `json:"feedback" form:"feedback"`
	Created_At  time.Time `json:"created_at" form:"created_at"`
	Url_File    string    `json:"url_file" form:"url_file"`
	Status      string    `json:"status" form:"status"`
	Nama_User   string    `json:"nama_user" form:"nama_user"`
	Nama_Mentee string    `json:"nama_mentee" form:"nama_mentee"`
}

func FromCore(dataCore log.LogCore) LogResponse {
	return LogResponse{
		ID:          dataCore.ID,
		Feedback:    dataCore.Feedback,
		Created_At:  dataCore.Created_At,
		Url_File:    dataCore.Url_File,
		Status:      dataCore.Status,
		Nama_User:   dataCore.Nama_User,
		Nama_Mentee: dataCore.Nama_Mentee,
	}
}

func FromCoreList(dataCore []log.LogCore) []LogResponse {
	var dataResponse []LogResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse

}
