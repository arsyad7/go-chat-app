package entity

type (
	DefaultResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
		Error   any    `json:"error,omitempty"`
	}
)
