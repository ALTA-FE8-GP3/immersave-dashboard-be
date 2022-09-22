package data

import (
	"errors"
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

	// txPreload := repo.db.Preload("Class").Find(&classData
	tx := repo.db.Preload("Class").Find(&allMentee)
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

func (repo *menteeData) UpdateMentee(data mentee.MenteeCore) (int, error) {
	var menteeUpdate Mentee
	txDataOld := repo.db.First(&menteeUpdate, data.ID)
	// result := repo.db.Model(&Mentee{}).Where("id = ?", data.ID).Updates(fromCore(data))
	if txDataOld.Error != nil {
		return -1, txDataOld.Error
	}

	if data.Nama_Mentee != "" {
		menteeUpdate.Nama_Mentee = data.Nama_Mentee
	}

	if data.Address != "" {
		menteeUpdate.Address = data.Address
	}

	if data.Home_Address != "" {
		menteeUpdate.Home_Address = data.Home_Address
	}

	if data.Email != "" {
		menteeUpdate.Email = data.Email
	}

	if data.Gender != "" {
		menteeUpdate.Gender = data.Gender
	}

	if data.Telegram != "" {
		menteeUpdate.Telegram = data.Telegram
	}

	if data.Phone != 0 {
		menteeUpdate.Phone = data.Phone
	}

	if data.Discord != "" {
		menteeUpdate.Discord = data.Discord
	}

	if data.Nama_Emergency != "" {
		menteeUpdate.Nama_Emergency = data.Nama_Emergency
	}
	if data.Phone_Emergency != 0 {
		menteeUpdate.Phone_Emergency = data.Phone_Emergency
	}

	if data.Status_Emergency != "" {
		menteeUpdate.Status_Emergency = data.Status_Emergency
	}

	if data.Type != "" {
		menteeUpdate.Type = data.Type
	}

	if data.Major != "" {
		menteeUpdate.Major = data.Major
	}

	if data.ClassID != 0 {
		menteeUpdate.ClassID = data.ClassID
	}

	tx := repo.db.Save(&menteeUpdate)
	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected != 1 {
		return 0, errors.New("zero row affected, fail update")
	}

	return int(tx.RowsAffected), nil
}
