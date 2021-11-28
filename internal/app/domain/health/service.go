package health

import "context"

type Service struct {
	repository Repository
}

func (s *Service) HealthReadDB(ctx context.Context) *Health {
	return s.repository.HealthReadDB(ctx)
}

func (s *Service) HealthWriteDB(ctx context.Context) *Health {
	return s.repository.HealthWriteDB(ctx)
}

func (s *Service) HealthRedisDB(ctx context.Context) *Health {
	return s.repository.HealthRedisDB(ctx)
}
