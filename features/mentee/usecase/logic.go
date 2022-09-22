package usecase

import (
	"errors"
	"project/immersive-dashboard/features/class"
	"project/immersive-dashboard/features/mentee"
)

type menteeUsecase struct {
	menteeData mentee.DataInterface
	classData  class.DataInterface
}

func New(data mentee.DataInterface, dataClass class.DataInterface) mentee.UsecaseInterface {
	return &menteeUsecase{
		menteeData: data,
		classData:  dataClass,
	}
}

func (usecase *menteeUsecase) PostData(data mentee.MenteeCore) (int, error) {

	if data.Nama_Mentee == "" || data.Address == "" || data.Home_Address == "" || data.Email == "" || data.Gender == "" || data.Telegram == "" || data.Phone == 0 || data.Discord == "" || data.Nama_Emergency == "" || data.Phone_Emergency == 0 || data.Status_Emergency == "" || data.Type == "" || data.Major == "" || data.Graduate == "" || data.Status == "" || data.Id_Class == 0 {
		return -1, errors.New("data input ada yang kosong")
	}

	dataClass, errClass := usecase.classData.SelectById(int(data.Id_Class))
	if errClass != nil {
		return -1, errClass
	}

	data.Nama_Class = dataClass.Nama_Class
	row, err := usecase.menteeData.InsertData(data)
	if err != nil {
		return -1, err
	}

	return row, err

}
