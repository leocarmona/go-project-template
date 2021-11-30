package outbound

type (
	CreateBookResponse struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}

	ReadBookByIdResponse struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}

	UpdateBookByIdResponse struct {
		Updated bool `json:"updated"`
	}

	DeleteBookByIdResponse struct {
		Deleted bool `json:"deleted"`
	}

	ListBookResponse struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}
)
