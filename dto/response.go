package dto

type ResponseSuccess struct {
	Message string `json:"message"`
}

type ResponseParams struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
