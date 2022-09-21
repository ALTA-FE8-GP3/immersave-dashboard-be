package delivery

import "project/immersive-dashboard/features/class"

type ClassRequest struct {
	Nama_Class string `json:"nama_class" form:"nama_class"`
}

func ToCore(dataRequest ClassRequest) class.ClassCore {
	return class.ClassCore{
		Nama_Class: dataRequest.Nama_Class,
	}
}
