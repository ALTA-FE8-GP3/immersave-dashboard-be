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

	newUser := fromCore(data)
	fmt.Println(newUser)
	tx := repo.db.Create(&newUser)
	if tx.Error != nil {
		return 0, tx.Error
	}

	// token, errToken := middlewares.CreateToken(int(newUser.ID))
	// if errToken != nil {
	// 	return "", -1, errToken
	// }

	return int(tx.RowsAffected), nil
}

func (repo *menteeData) SelectAllMentee() ([]mentee.MenteeCore, error) {
	var allMentee []Mentee

	// txPreload := repo.db.Preload("Class").Find(&classData
	tx := repo.db.Preload("Class").Find(&allMentee)
	if tx.Error != nil {
		return nil, tx.Error
	}

	menteList := toCoreList(allMentee)
	return menteList, nil
}
