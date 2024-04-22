package errorhandler

import (
	"encoding/json"
	"net/http"

	"github.com/tafhdytllah/customer-list/dto"
	"github.com/tafhdytllah/customer-list/helper"
)

type ServiceError struct {
	Message string `json:"message"`
}

func HandlerError(res http.ResponseWriter, err error) {
	var statusCode int

	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	}

	response := helper.Response(dto.ResponseParams{
		Message: err.Error(),
	})

	res.WriteHeader(statusCode)
	json.NewEncoder(res).Encode(response)
}
