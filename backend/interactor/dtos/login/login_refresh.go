package login

import (
	entity "auth_service/domain/entities/user"
	"net"
)

type LoginRefreshInputDTO struct {
	AccessToken  string
	RefreshToken string
	IP           net.IP
}

type LoginRefreshOutputDTO struct {
	AccessToken entity.TokenData
}
