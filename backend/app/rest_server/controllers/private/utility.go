package utility

import (
	"errors"
	"fmt"
	"net/http"

	prjerror "auth_service/interactor/error"
)

func TransformErrorToHttpError(err error) (int, string) {
	switch {
	case errors.Is(err, &prjerror.NotFoundError{}):
		return http.StatusNotFound, "Not found"
	case errors.Is(err, &prjerror.NotAuthorizedError{}):
		return http.StatusUnauthorized, "Not authorized"
	case errors.Is(err, &prjerror.InvalidJWTError{}):
		return http.StatusBadRequest, "Invalid JWT token format"
	case errors.Is(err, &prjerror.AlreadyExistsError{}):
		return http.StatusConflict, "Element already exists"
	case errors.Is(err, &prjerror.InvalidArgumentError{}):
		return http.StatusBadRequest, "Invalid argument error"
	default:
		fmt.Println(err)
		return http.StatusInternalServerError, "Internal server error"
	}
}
