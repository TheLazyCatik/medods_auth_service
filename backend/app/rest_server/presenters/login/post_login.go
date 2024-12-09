package login

import (
	api "auth_service/app/rest_server/api/login"
	"auth_service/domain/entities/user"
	dto "auth_service/interactor/dtos/login"

	"github.com/gin-gonic/gin"
)

type PostLoginPresenter struct {
	Context *gin.Context
}

func (p *PostLoginPresenter) Present(outputDTO *dto.PostLoginOutputDTO) *api.PostLoginResponse {
	if outputDTO == nil {
		return nil
	}

	response := &api.PostLoginResponse{
		Result: api.PostLoginResult{
			AccessToken: user.TokenData{
				Token:     outputDTO.AccessToken.Token,
				ExpiresAt: outputDTO.AccessToken.ExpiresAt,
			},
			RefreshToken: user.TokenData{
				Token:     outputDTO.RefreshToken.Token,
				ExpiresAt: outputDTO.RefreshToken.ExpiresAt,
			},
		},
	}
	return response
}
