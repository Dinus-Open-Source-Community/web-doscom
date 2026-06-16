package response

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data"`
	Error   any    `json:"error,omitempty"`
}
