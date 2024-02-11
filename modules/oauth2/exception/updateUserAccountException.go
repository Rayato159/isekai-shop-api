package exception

import (
	"fmt"
	"net/http"

	"github.com/Rayato159/isekai-shop-api/server/writter"
)

type UpdateUserAccountException struct {
	*writter.ErrorMessage
}

func NewUpdateUserAccountException(userId string, err error) writter.CustomErrorResponse {
	return &UpdateUserAccountException{
		ErrorMessage: &writter.ErrorMessage{
			Message: fmt.Sprintf("Error updating user account with id: %s; %s", userId, err.Error()),
		},
	}
}

func (e *UpdateUserAccountException) GetErrorMessage() string {
	return e.Message
}

func (e *UpdateUserAccountException) GetStatusCode() int {
	return http.StatusInternalServerError
}
