package usecase

import (
	"errors"
	"project/immersive-dashboard/features/log"
)

type logUsecase struct {
	logData log.DataInterface
}

func New(data log.DataInterface) log.UsecaseInterface {
	return &logUsecase{
		logData: data,
	}
}

func (usecase *logUsecase) PostData(data log.LogCore) (int, error) {

	if data.Feedback == "" || data.Status == "" || data.UserID == 0 || data.MenteeID == 0 {
		return -1, errors.New("data input ada yang kosong")
	}

	row, err := usecase.logData.InsertData(data)
	if err != nil {
		return -1, err
	}

	return row, err

}

func (usecase *logUsecase) GetLogById(id int) (log.LogCore, error) {
	dataLog, err := usecase.logData.SelectLogById(id)
	return dataLog, err
}

func (usecase *logUsecase) GetAlllog() ([]log.LogCore, error) {
	results, err := usecase.logData.SelectAlllog()
	return results, err

}
