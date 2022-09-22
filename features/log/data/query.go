package data

import (
	"fmt"
	"project/immersive-dashboard/features/log"

	"gorm.io/gorm"
)

type logData struct {
	db *gorm.DB
}

func New(db *gorm.DB) log.DataInterface {
	return &logData{
		db: db,
	}

}

func (repo *logData) InsertData(data log.LogCore) (int, error) {

	newlog := fromCore(data)
	fmt.Println(newlog)
	tx := repo.db.Create(&newlog)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *logData) SelectLogById(id int) (log.LogCore, error) {
	var dataLog Log
	dataLog.ID = uint(id)

	tx := repo.db.First(&dataLog)

	if tx.Error != nil {
		return log.LogCore{}, tx.Error
	}

	dataMenteeCore := dataLog.toCore()
	return dataMenteeCore, nil

}
