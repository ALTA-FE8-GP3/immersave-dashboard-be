package usecase

import (
	"errors"
	"project/immersive-dashboard/features/class"
)

type classUsecase struct {
	classData class.DataInterface
}

func New(data class.DataInterface) class.UsecaseInterface {
	return &classUsecase{
		classData: data,
	}
}

func (usecase *classUsecase) PostData(data class.ClassCore) (int, error) {
	if data.Nama_Class == "" {
		return -1, errors.New("data input ada yang kosong")
	}

	row, err := usecase.classData.InsertClass(data)
	if err != nil {
		return -1, err
	}

	return row, err

}

func (usecase *classUsecase) PutData(data class.ClassCore) (int, error) {
	row, err := usecase.classData.UpdateClass(data)
	return row, err
}

func (usecase *classUsecase) GetAllClass() ([]class.ClassCore, error) {
	results, err := usecase.classData.SelectAllClass()
	return results, err

}

func (usecase *classUsecase) GetClassById(id_class int) (class.ClassCore, error) {
	dataClass, err := usecase.classData.SelectById(id_class)
	return dataClass, err
}
