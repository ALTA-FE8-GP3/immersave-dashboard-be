package usecase

import (
	"errors"
	"project/immersive-dashboard/features/mentee"
)

type menteeUsecase struct {
	menteeData mentee.DataInterface
}

func New(data mentee.DataInterface) mentee.UsecaseInterface {
	return &menteeUsecase{
		menteeData: data,
	}
}

func (usecase *menteeUsecase) PostData(data mentee.MenteeCore) (int, error) {

	if data.Nama_Mentee == "" || data.Address == "" || data.Home_Address == "" || data.Email == "" || data.Gender == "" || data.Telegram == "" || data.Phone == 0 || data.Discord == "" || data.Nama_Emergency == "" || data.Phone_Emergency == 0 || data.Status_Emergency == "" || data.Type == "" || data.Major == "" || data.Graduate == "" || data.Status == "" || data.ClassID == 0 {
		return -1, errors.New("data input ada yang kosong")
	}

	row, err := usecase.menteeData.InsertData(data)
	if err != nil {
		return -1, err
	}

	return row, err

}

func (usecase *menteeUsecase) GetAllMentee() ([]mentee.MenteeCore, error) {
	results, err := usecase.menteeData.SelectAllMentee()
	return results, err

}
