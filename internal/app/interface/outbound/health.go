package outbound

type HealthResponse struct {
	Healthy bool   `json:"healthy"`
	ReadDB  string `json:"read_db"`
	WriteDB string `json:"write_db"`
	RedisDB string `json:"redis_db"`
}
