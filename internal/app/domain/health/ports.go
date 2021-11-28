package health

import "context"

type Repository interface {
	HealthReadDB(ctx context.Context) *Health
	HealthWriteDB(ctx context.Context) *Health
	HealthRedisDB(ctx context.Context) *Health
}
