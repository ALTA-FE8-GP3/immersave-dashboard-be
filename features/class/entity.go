package class

type ClassCore struct {
	ID         uint
	Nama_Class string
	UserID     uint
	Nama_User  string
}

type UsecaseInterface interface {
	GetAllClass() (data []ClassCore, err error)
	GetClassById(id int) (data ClassCore, err error)
	PostData(data ClassCore) (row int, err error)
	PutData(data ClassCore) (row int, err error)
	DeleteClass(id int) (row int, err error)
}

type DataInterface interface {
	SelectAllClass() (data []ClassCore, err error)
	SelectById(id int) (data ClassCore, err error)
	InsertClass(data ClassCore) (row int, err error)
	UpdateClass(data ClassCore) (row int, err error)
	ClassDelete(id int) (row int, err error)
}
