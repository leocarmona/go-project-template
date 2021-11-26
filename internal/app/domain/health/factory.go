package health

func New(up bool, err error) *Health {
	return &Health{
		Up:    up,
		Error: err,
	}
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
