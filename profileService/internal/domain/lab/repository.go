package lab

type Repository interface {
	GetAll() ([]*Lab, error)
	GetLabsByInstituteID(InstituteID int64) ([]*Lab, error)
}
