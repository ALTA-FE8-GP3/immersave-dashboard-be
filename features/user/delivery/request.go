package delivery

import "project/immersive-dashboard/features/user"

type UserRequest struct {
	Nama_User string `json:"nama_user" form:"nama_user"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Role      string `json:"role" form:"role"`
	Team      string `json:"team" form:"team"`
	Status    string `json:"status" form:"status"`
}

func ToCore(dataRequest UserRequest) user.UserCore {
	return user.UserCore{
		Nama_User: dataRequest.Nama_User,
		Email:     dataRequest.Email,
		Password:  dataRequest.Password,
		Role:      dataRequest.Role,
		Team:      dataRequest.Team,
		Status:    dataRequest.Status,
	}
}
