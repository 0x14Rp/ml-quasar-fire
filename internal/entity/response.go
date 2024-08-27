package entity

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type TopSecretResponse struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}

type ErrorResponse struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
	Timestamp string `json:"timestamp"`
}
