package helper

import "github.com/tafhdytllah/customer-list/dto"

type ResponseWithData struct {
	Data any `json:"data"`
}

type ResponseWithoutData struct {
	Message string `json:"message"`
}

func Response(params dto.ResponseParams) any {
	var response any

	if params.Data != nil {
		response = &ResponseWithData{
			Data: params.Data,
		}
	} else {
		response = &ResponseWithoutData{
			Message: params.Message,
		}
	}

	return response
}
