package login

import (
	"net"
	"net/http"

	api "auth_service/app/rest_server/api/login"
	utility "auth_service/app/rest_server/controllers/private"
	presenter "auth_service/app/rest_server/presenters/login"
	emailgate "auth_service/infra/gateway/email"
	sessionrepo "auth_service/infra/repository/session"
	userrepo "auth_service/infra/repository/user"

	"github.com/gin-gonic/gin"

	dto "auth_service/interactor/dtos/login"
	usecase "auth_service/interactor/use_cases/login"
)

// @summary Login refresh
// @Description Generate new access token from access and refresh token
// @Tags Login
// @Accept json
// @Produce json
// @Param request body login.LoginRefreshRequest true "Login refresh request body"
// @Success 200 {object} login.LoginRefreshResponse "Token With User ID "
// @Router /login/refresh [post]
func (s *LoginServer) LoginRefresh(c *gin.Context) {

	var request api.LoginRefreshRequest

	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	clientIPString := c.ClientIP()
	clientIP := net.ParseIP(clientIPString)

	inputDTO := dto.LoginRefreshInputDTO{
		AccessToken:  request.AccessToken,
		RefreshToken: request.RefreshToken,
		IP:           clientIP,
	}

	presenter := &presenter.LoginRefreshPresenter{}

	userRepo := userrepo.NewUserPostgresRepo(s.postgresPool)
	sessionRepo := sessionrepo.NewSessionPostgresRepo(s.postgresPool)
	emailGate := emailgate.NewEmailGate()

	useCase := usecase.LoginRefreshUseCase{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
		EmailGate:   emailGate,
	}

	outputDTO, err := useCase.Execute(s.ctx, inputDTO)

	if err != nil {
		code, err := utility.TransformErrorToHttpError(err)
		c.AbortWithStatusJSON(code, gin.H{"error": err})
		return
	}

	response := presenter.Present(outputDTO)

	c.JSON(http.StatusOK, response)
}
