package delivery

import "project/immersive-dashboard/features/mentee"

type menteeResponse struct {
	ID          uint   `json:"ID" form:"ID"`
	Nama_Mentee string `json:"nama_mentee" form:"nama_mentee"`
	Email       string `json:"email" form:"email"`
	Gender      string `json:"gender" form:"gender"`
	Telegram    string `json:"telegram" form:"telegram"`
	Phone       uint   `json:"phone" form:"phone"`
	Discord     string `json:"discord" form:"discord"`
	Type        string `json:"type" form:"type"`
	Major       string `json:"major" form:"major"`
	Graduate    string `json:"graduate" form:"graduate"`
	Status      string `json:"status" form:"status"`
	Id_Class    uint   `json:"id_class" form:"id_class"`
	Nama_Class  string `json:"nama_class" form:"nama_class"`
}

func FromCore(dataCore mentee.MenteeCore) menteeResponse {
	return menteeResponse{
		ID:          dataCore.ID,
		Nama_Mentee: dataCore.Nama_Mentee,
		Email:       dataCore.Email,
		Gender:      dataCore.Gender,
		Telegram:    dataCore.Telegram,
		Phone:       dataCore.Phone,
		Discord:     dataCore.Discord,
		Type:        dataCore.Type,
		Major:       dataCore.Major,
		Graduate:    dataCore.Graduate,
		Status:      dataCore.Status,
		Id_Class:    dataCore.Id_Class,
		Nama_Class:  dataCore.Nama_Class,
	}
}

func FromCoreList(dataCore []mentee.MenteeCore) []menteeResponse {
	var dataResponse []menteeResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse

}
