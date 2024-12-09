package login

import (
	"context"
	"errors"

	entity "auth_service/domain/entities/user"
	token "auth_service/infra/token_manager"
	dto "auth_service/interactor/dtos/login"
	prjerror "auth_service/interactor/error"
	repo "auth_service/interactor/ifaces/repos"
)

type PostLoginUseCase struct {
	UserRepo     repo.UserRepo
	SessionRepo  repo.SessionRepo
	TokenManager token.JWTTokenManager
}

func (u *PostLoginUseCase) Execute(ctx context.Context, inputDTO dto.PostLoginInputDTO) (*dto.PostLoginOutputDTO, error) {
	user, err := u.UserRepo.GetUserByID(ctx, inputDTO.UserID)

	if err != nil {
		return nil, err
	}

	session, err := u.SessionRepo.GetSessionByUserID(ctx, user.ID)

	if !errors.Is(err, &prjerror.NotFoundError{}) {
		if session != nil {
			u.SessionRepo.DeleteSessionByID(ctx, session.ID)
		}
	}

	accessToken, accessClaims, err := u.TokenManager.GenerateAccessToken(user.ID, inputDTO.IP)

	if err != nil {
		return nil, err
	}

	refreshToken, refreshClaims, err := u.TokenManager.GenerateRefreshToken()

	if err != nil {
		return nil, err
	}

	hashRefreshToken, err := u.TokenManager.GenerateHashFromToken(*refreshToken)

	if err != nil {
		return nil, err
	}

	_, err = u.SessionRepo.CreateSession(ctx, repo.CreateSessionArgs{
		UserID:        user.ID,
		HashToken:     *hashRefreshToken,
		AccessTokenID: accessClaims.TokenID,
		ExpiresAt:     refreshClaims.ExpiresAt.Time.UTC(),
	})

	if err != nil {
		return nil, err
	}

	return &dto.PostLoginOutputDTO{
		AccessToken: entity.TokenData{
			Token:     *accessToken,
			ExpiresAt: accessClaims.ExpiresAt.Time.UTC(),
		},
		RefreshToken: entity.TokenData{
			Token:     *refreshToken,
			ExpiresAt: refreshClaims.ExpiresAt.Time.UTC(),
		},
	}, nil
}
