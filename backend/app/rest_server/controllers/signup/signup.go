package signup

import (
	"net/http"
	"net/mail"

	api "auth_service/app/rest_server/api/signup"
	utility "auth_service/app/rest_server/controllers/private"
	presenter "auth_service/app/rest_server/presenters/signup"

	"github.com/gin-gonic/gin"

	userrepo "auth_service/infra/repository/user"

	dto "auth_service/interactor/dtos/signup"
	usecase "auth_service/interactor/use_cases/signup"
)

// @summary Signup
// @Description Retrieve user ID
// @Tags Signup
// @Accept json
// @Produce json
// @Param request body signup.SignupRequest true "Signup request body"
// @Success 200 {object} signup.SignupResponse "User ID"
// @Router /signup [post]
func (s *SignupServer) Signup(c *gin.Context) {
	var request api.SignupRequest

	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	_, err = mail.ParseAddress(request.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	inputDTO := dto.SignupInputDTO{
		Email: request.Email,
	}

	userRepo := userrepo.NewUserPostgresRepo(s.postgresPool)

	presenter := &presenter.SignupPresenter{}

	useCase := usecase.SignupUseCase{
		UserRepo: userRepo,
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
