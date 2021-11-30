package book

func New(id int64, name string) *Book {
	return &Book{
		Id:   id,
		Name: name,
	}
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
