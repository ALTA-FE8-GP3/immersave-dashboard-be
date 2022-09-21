package usecase

import (
	"errors"
	"project/immersive-dashboard/features/user"
)

type userUsecase struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &userUsecase{
		userData: data,
	}
}

func (usecase *userUsecase) PostLogin(data user.UserCore) (string, error) {
	if data.Email == "" || data.Password == "" {
		return "", errors.New("data input ada yang kosong")
	}

	token, err := usecase.userData.LoginUser(data)
	if err != nil {
		return "", err
	}

	return token, err

}

func (usecase *userUsecase) PostData(data user.UserCore) (int, error) {
	if data.Nama_User == "" || data.Email == "" || data.Password == "" {
		return -1, errors.New("data input ada yang kosong")
	}

	row, err := usecase.userData.InsertData(data)
	if err != nil {
		return -1, err
	}

	return row, err

}

func (usecase *userUsecase) PutData(data user.UserCore) (int, error) {
	row, err := usecase.userData.UpdateUser(data)
	return row, err
}
