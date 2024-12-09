package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	logincontrollers "auth_service/app/rest_server/controllers/login"
	signupcontrollers "auth_service/app/rest_server/controllers/signup"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Serve(ctx context.Context, port int, postgresPool *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	registerLogin(ctx, postgresPool, router)
	registerSignup(ctx, postgresPool, router)

	return router
}

func registerLogin(ctx context.Context, postgresPool *pgxpool.Pool, router *gin.Engine) {
	loginServer := logincontrollers.New(ctx, postgresPool)

	router.POST("/login", loginServer.PostLogin)
	router.POST("/login/refresh", loginServer.LoginRefresh)
}

func registerSignup(ctx context.Context, postgresPool *pgxpool.Pool, router *gin.Engine) {
	signupServer := signupcontrollers.New(ctx, postgresPool)

	router.POST("/signup", signupServer.Signup)
}
