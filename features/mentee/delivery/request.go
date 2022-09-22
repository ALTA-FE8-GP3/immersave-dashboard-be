package delivery

import "project/immersive-dashboard/features/mentee"

type MenteeRequest struct {
	Nama_Mentee      string `json:"nama_mentee" form:"nama_mentee"`
	Address          string `json:"address" form:"address"`
	Home_Address     string `json:"home_address" form:"home_address"`
	Email            string `json:"email" form:"email"`
	Gender           string `json:"gender" form:"gender"`
	Telegram         string `json:"telegram" form:"telegram"`
	Phone            uint   `json:"phone" form:"phone"`
	Discord          string `json:"discord" form:"discord"`
	Nama_Emergency   string `json:"nama_emergency" form:"nama_emergency"`
	Phone_Emergency  uint   `json:"phone_emergency" form:"phone_emergency"`
	Status_Emergency string `json:"status_emergency" form:"status_emergency"`
	Type             string `json:"type" form:"type"`
	Major            string `json:"major" form:"major"`
	Graduate         string `json:"graduate" form:"graduate"`
	Status           string `json:"status" form:"status"`
	ClassID          uint   `json:"class_id" form:"class_id"`
}

func ToCore(dataRequest MenteeRequest) mentee.MenteeCore {
	return mentee.MenteeCore{
		Nama_Mentee:      dataRequest.Nama_Mentee,
		Address:          dataRequest.Address,
		Home_Address:     dataRequest.Home_Address,
		Email:            dataRequest.Email,
		Gender:           dataRequest.Gender,
		Telegram:         dataRequest.Telegram,
		Phone:            dataRequest.Phone,
		Discord:          dataRequest.Discord,
		Nama_Emergency:   dataRequest.Nama_Emergency,
		Phone_Emergency:  dataRequest.Phone_Emergency,
		Status_Emergency: dataRequest.Status_Emergency,
		Type:             dataRequest.Type,
		Major:            dataRequest.Major,
		Graduate:         dataRequest.Graduate,
		Status:           dataRequest.Status,
		ClassID:          dataRequest.ClassID,
	}
}
