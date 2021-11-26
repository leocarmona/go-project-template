package outbound

type HealthResponse struct {
	Healthy bool   `json:"healthy"`
	ReadDB  string `json:"read_db"`
	WriteDB string `json:"write_db"`
}
