package data

import (
	"fmt"
	"project/immersive-dashboard/features/mentee"

	"gorm.io/gorm"
)

type menteeData struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.DataInterface {
	return &menteeData{
		db: db,
	}

}

func (repo *menteeData) InsertData(data mentee.MenteeCore) (int, error) {

	newmentee := fromCore(data)
	fmt.Println(newmentee)
	tx := repo.db.Create(&newmentee)
	if tx.Error != nil {
		return 0, tx.Error
	}

	// token, errToken := middlewares.CreateToken(int(newmentee.ID))
	// if errToken != nil {
	// 	return "", -1, errToken
	// }

	return int(tx.RowsAffected), nil
}

func (repo *menteeData) SelectAllMentee() ([]mentee.MenteeCore, error) {
	var allMentee []Mentee
	tx := repo.db.Find(&allMentee)

	if tx.Error != nil {
		return nil, tx.Error
	}

	menteList := toCoreList(allMentee)
	return menteList, nil
}

func (repo *menteeData) DeleteMentee(id int) (int, error) {
	var menteeData Mentee
	menteeData.ID = uint(id)
	tx := repo.db.Delete(&menteeData)

	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *menteeData) SelectMenteeById(id int) (mentee.MenteeCore, error) {
	var dataMentee Mentee
	dataMentee.ID = uint(id)

	tx := repo.db.First(&dataMentee)

	if tx.Error != nil {
		return mentee.MenteeCore{}, tx.Error
	}

	dataMenteeCore := dataMentee.toCore()
	return dataMenteeCore, nil

}
