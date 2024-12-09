package login

import (
	entity "auth_service/domain/entities/user"
	"net"

	"github.com/google/uuid"
)

type PostLoginInputDTO struct {
	UserID uuid.UUID
	IP     net.IP
}

type PostLoginOutputDTO struct {
	AccessToken  entity.TokenData
	RefreshToken entity.TokenData
}
