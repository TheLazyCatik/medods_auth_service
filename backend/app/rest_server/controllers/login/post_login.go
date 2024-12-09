package login

import (
	"net"
	"net/http"

	utility "auth_service/app/rest_server/controllers/private"
	presenter "auth_service/app/rest_server/presenters/login"

	sessionrepo "auth_service/infra/repository/session"
	userrepo "auth_service/infra/repository/user"

	"github.com/gin-gonic/gin"

	api "auth_service/app/rest_server/api/login"

	dto "auth_service/interactor/dtos/login"
	usecase "auth_service/interactor/use_cases/login"
)

// @summary Post login
// @Description Retrieve access token, refresh token
// @Tags Login
// @Accept json
// @Produce json
// @Param request body login.PostLoginRequest true "Access refresh tokens"
// @Success 200 {object} login.PostLoginResponse "Token With User ID "
// @Router /login [post]
func (s *LoginServer) PostLogin(c *gin.Context) {
	var request api.PostLoginRequest

	err := c.BindJSON(&request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	clientIPString := c.ClientIP()
	clientIP := net.ParseIP(clientIPString)

	inputDTO := dto.PostLoginInputDTO{
		UserID: request.UserID,
		IP:     clientIP,
	}

	userRepo := userrepo.NewUserPostgresRepo(s.postgresPool)
	sessionRepo := sessionrepo.NewSessionPostgresRepo(s.postgresPool)

	presenter := &presenter.PostLoginPresenter{}

	useCase := usecase.PostLoginUseCase{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
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
