package data

import (
	"errors"
	"project/immersive-dashboard/features/user"
	"project/immersive-dashboard/middlewares"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userData{
		db: db,
	}

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (repo *userData) InsertData(data user.UserCore) (int, error) {
	hash_pass, errHash := HashPassword(data.Password)

	if errHash != nil {
		return -1, errHash
	}
	data.Password = hash_pass //memasukkan hasil enskripsi data password

	newUser := fromCore(data)

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

func (repo *userData) LoginUser(data user.UserCore) (string, error) {
	var userData User
	tx := repo.db.Where("email = ?", data.Email).First(&userData)

	check_result := CheckPasswordHash(data.Password, userData.Password)

	if !check_result {
		return "", errors.New("password salah")
	}

	if tx.Error != nil {
		return "", tx.Error
	}

	if tx.RowsAffected != 1 {
		return "", errors.New("login failed")
	}

	token, errToken := middlewares.CreateToken(int(userData.ID), userData.Role)
	if errToken != nil {
		return "", errToken
	}

	return token, nil

}
func (repo *userData) UpdateUser(data user.UserCore) (int, error) {
	var userUpdate User
	txDataOld := repo.db.First(&userUpdate, data.ID)

	if txDataOld.Error != nil {
		return -1, txDataOld.Error
	}

	if data.Nama_User != "" {
		userUpdate.Nama_User = data.Nama_User
	}

	if data.Email != "" {
		userUpdate.Email = data.Email
	}

	if data.Password != "" {
		hash_pass, errHash := HashPassword(data.Password)
		if errHash != nil {
			return -1, errHash
		}
		userUpdate.Password = hash_pass
	}

	if data.Role != "" {
		userUpdate.Role = data.Role
	}

	if data.Team != "" {
		userUpdate.Team = data.Team
	}

	if data.Status != "" {
		userUpdate.Status = data.Status
	}
	tx := repo.db.Save(&userUpdate)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
