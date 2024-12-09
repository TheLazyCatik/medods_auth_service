package signup

import (
	"context"

	dto "auth_service/interactor/dtos/signup"
	repo "auth_service/interactor/ifaces/repos"
)

type SignupUseCase struct {
	UserRepo repo.UserRepo
}

func (u *SignupUseCase) Execute(ctx context.Context, inputDTO dto.SignupInputDTO) (*dto.SignupOutputDTO, error) {

	userID, err := u.UserRepo.CreateUser(
		ctx,
		repo.CreateUserArgs{
			Email: inputDTO.Email,
		},
	)

	if err != nil {
		return nil, err
	}

	return &dto.SignupOutputDTO{
		UserID: *userID,
	}, nil
}
