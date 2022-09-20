package delivery

import "project/immersive-dashboard/features/user"

type UserResponse struct {
	ID        uint   `json:"id" form:"id"`
	Nama_User string `json:"nama_user" form:"nama_user"`
	Email     string `json:"email" form:"email"`
	Role      string `json:"role" form:"role"`
	Team      string `json:"team" form:"team"`
	Status    string `json:"status" form:"status"`
}

func FromCore(dataCore user.UserCore) UserResponse {
	return UserResponse{
		ID:        dataCore.ID,
		Nama_User: dataCore.Nama_User,
		Email:     dataCore.Email,
		Role:      dataCore.Role,
		Team:      dataCore.Team,
		Status:    dataCore.Status,
	}
}

func FromCoreList(dataCore []user.UserCore) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse

}
