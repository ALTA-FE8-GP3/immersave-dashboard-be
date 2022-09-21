package data

import (
	"project/immersive-dashboard/features/class"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Nama_Class string
}

func fromCore(dataCore class.ClassCore) Class {
	return Class{
		Nama_Class: dataCore.Nama_Class,
	}
}

func (dataClass *Class) toCore() class.ClassCore {
	return class.ClassCore{
		ID:         dataClass.ID,
		Nama_Class: dataClass.Nama_Class,
	}
}

func toCoreList(dataClass []Class) []class.ClassCore {
	var dataCore []class.ClassCore

	for key := range dataClass {
		dataCore = append(dataCore, dataClass[key].toCore())
	}
	return dataCore
}
