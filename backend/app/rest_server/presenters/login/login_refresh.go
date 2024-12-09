package login

import (
	api "auth_service/app/rest_server/api/login"
	"auth_service/domain/entities/user"
	dto "auth_service/interactor/dtos/login"

	"github.com/gin-gonic/gin"
)

type LoginRefreshPresenter struct {
	Context *gin.Context
}

func (p *LoginRefreshPresenter) Present(outputDTO *dto.LoginRefreshOutputDTO) *api.LoginRefreshResponse {
	if outputDTO == nil {
		return nil
	}

	response := &api.LoginRefreshResponse{
		Result: api.LoginRefreshResult{
			AccessToken: user.TokenData{
				Token:     outputDTO.AccessToken.Token,
				ExpiresAt: outputDTO.AccessToken.ExpiresAt,
			},
		},
	}
	return response
}
