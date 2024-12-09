package login

import (
	"context"
	"time"

	entity "auth_service/domain/entities/user"
	token "auth_service/infra/token_manager"
	dto "auth_service/interactor/dtos/login"
	prjerror "auth_service/interactor/error"
	gate "auth_service/interactor/ifaces/gates"
	repo "auth_service/interactor/ifaces/repos"
)

type LoginRefreshUseCase struct {
	UserRepo     repo.UserRepo
	SessionRepo  repo.SessionRepo
	TokenManager token.JWTTokenManager
	EmailGate    gate.EmailGate
}

func (u *LoginRefreshUseCase) Execute(ctx context.Context, inputDTO dto.LoginRefreshInputDTO) (*dto.LoginRefreshOutputDTO, error) {
	accessClaims, err := u.TokenManager.ValidateAccessToken(inputDTO.AccessToken)

	if err != nil {
		return nil, &prjerror.InvalidJWTError{}
	}

	_, err = u.TokenManager.ValidateRefreshToken(inputDTO.RefreshToken)

	if err != nil {
		return nil, &prjerror.InvalidJWTError{}
	}

	session, err := u.SessionRepo.GetSessionByUserID(ctx, accessClaims.UserID)

	if err != nil {
		return nil, err
	}

	if session.ExpiresAt.Before(time.Now()) {
		return nil, &prjerror.NotAuthorizedError{}
	}

	if !u.TokenManager.CompareHashTokenAndToken(session.HashToken, inputDTO.RefreshToken) {
		return nil, &prjerror.NotAuthorizedError{}
	}

	if accessClaims.TokenID != session.AccessTokenID {
		return nil, &prjerror.NotAuthorizedError{}
	}

	user, err := u.UserRepo.GetUserByID(ctx, accessClaims.UserID)

	if err != nil {
		return nil, err
	}

	if !accessClaims.IP.Equal(inputDTO.IP) {
		u.EmailGate.SendNewIPLoginNotificationEmail(ctx, inputDTO.IP.String(), user.Email, accessClaims.IssuedAt.Time.UTC())
	}

	newAccessToken, newAccessClaims, err := u.TokenManager.GenerateAccessToken(accessClaims.UserID, accessClaims.IP)

	if err != nil {
		return nil, err
	}

	err = u.SessionRepo.UpdateSessionByID(
		ctx, session.ID,
		repo.UpdateSessionArgs{
			AccessTokenID: newAccessClaims.TokenID,
			ExpiresAt:     newAccessClaims.ExpiresAt.Time.UTC(),
		},
	)

	if err != nil {
		return nil, err
	}

	return &dto.LoginRefreshOutputDTO{
		AccessToken: entity.TokenData{
			Token:     *newAccessToken,
			ExpiresAt: newAccessClaims.ExpiresAt.Time.UTC(),
		},
	}, nil
}
