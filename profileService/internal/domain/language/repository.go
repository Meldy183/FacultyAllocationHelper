package language

type Repository interface {
	GetAll() ([]*Language, error)
	GetByCode(code string) (*Language, error)
}
