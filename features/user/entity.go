package user

import "time"

type UserCore struct {
	ID         uint
	Nama_User  string
	Email      string
	Password   string
	Team       string
	Role       string
	Status     string
	Created_At time.Time
	Updated_At time.Time
}

type UsecaseInterface interface {
	PostLogin(data UserCore) (token string, err error)
	PostData(data UserCore) (row int, err error)
}

type DataInterface interface {
	InsertData(data UserCore) (row int, err error)
	LoginUser(data UserCore) (token string, err error)
}
