package data

import (
	"errors"
	"project/immersive-dashboard/features/class"

	"gorm.io/gorm"
)

type classData struct {
	db *gorm.DB
}

func New(db *gorm.DB) class.DataInterface {
	return &classData{
		db: db,
	}

}

func (repo *classData) InsertClass(data class.ClassCore) (int, error) {
	newClass := fromCore(data)

	tx := repo.db.Create(&newClass)
	if tx.Error != nil {
		return 0, tx.Error
	}

	// token, errToken := middlewares.CreateToken(int(newUser.ID))
	// if errToken != nil {
	// 	return "", -1, errToken
	// }

	return int(tx.RowsAffected), nil
}

func (repo *classData) SelectAllClass() ([]class.ClassCore, error) {
	var allClass []Class
	tx := repo.db.Find(&allClass)

	if tx.Error != nil {
		return nil, tx.Error
	}

	produk_List := toCoreList(allClass)
	return produk_List, nil
}

func (repo *classData) UpdateClass(data class.ClassCore) (int, error) {
	var classUpdate Class
	txDataOld := repo.db.First(&classUpdate, data.ID)

	if txDataOld.Error != nil {
		return -1, txDataOld.Error
	}

	if data.Nama_Class != "" {
		classUpdate.Nama_Class = data.Nama_Class
	}

	tx := repo.db.Save(&classUpdate)
	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected != 1 {
		return 0, errors.New("zero row affected, fail update")
	}

	return int(tx.RowsAffected), nil
}

func (repo *classData) SelectById(id_class int) (class.ClassCore, error) {
	var dataClass Class
	dataClass.ID = uint(id_class)

	tx := repo.db.First(&dataClass)

	if tx.Error != nil {
		return class.ClassCore{}, tx.Error
	}

	dataClassCore := dataClass.toCore()
	return dataClassCore, nil

}
