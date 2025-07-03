package institute

type Repository interface {
	GetById(instituteID int64) (*Institute, error)
	GetAll() ([]*Institute, error)
}
