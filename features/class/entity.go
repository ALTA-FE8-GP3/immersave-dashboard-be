package class

type ClassCore struct {
	ID         uint
	Nama_Class string
}

type UsecaseInterface interface {
	GetAllClass() (data []ClassCore, err error)
	PostLogin(data ClassCore) (token string, err error)
	PostData(data ClassCore) (row int, err error)
	PutData(data ClassCore) (row int, err error)
}

type DataInterface interface {
	SelectAllClassr() (data []ClassCore, err error)
	InsertData(data ClassCore) (row int, err error)
	LoginClass(data ClassCore) (token string, err error)
	UpdateClass(data ClassCore) (row int, err error)
}
