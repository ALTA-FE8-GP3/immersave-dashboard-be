package delivery

import "project/immersive-dashboard/features/class"

type ClassResponse struct {
	ID         uint   `json:"id" form:"id"`
	Nama_Class string `json:"nama_class" form:"nama_class"`
}

func FromCore(dataCore class.ClassCore) ClassResponse {
	return ClassResponse{
		ID:         dataCore.ID,
		Nama_Class: dataCore.Nama_Class,
	}
}

func FromCoreList(dataCore []class.ClassCore) []ClassResponse {
	var dataResponse []ClassResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse

}
