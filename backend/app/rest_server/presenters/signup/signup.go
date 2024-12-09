package signup

import (
	api "auth_service/app/rest_server/api/signup"
	dto "auth_service/interactor/dtos/signup"

	"github.com/gin-gonic/gin"
)

type SignupPresenter struct {
	Context *gin.Context
}

func (p *SignupPresenter) Present(outputDTO *dto.SignupOutputDTO) *api.SignupResponse {
	if outputDTO == nil {
		return nil
	}

	userIDStr := outputDTO.UserID.String()

	response := &api.SignupResponse{
		Result: api.SignupResult{
			UserID: userIDStr,
		},
	}
	return response
}
