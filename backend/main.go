package main

import (
	"context"
	"flag"
	"fmt"

	server "auth_service/app/rest_server"
	"auth_service/config"

	_ "auth_service/docs"
)

var (
	port = flag.Int("port", 8040, "The server port")
)

// @title           Medods Auth API
// @version         0.1.0
// @securityDefinitions.basic BasicAuth
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	flag.Parse()
	ctx := context.Background()

	postgresPool, errPostgres := config.NewPostgresPool(ctx)

	if errPostgres != nil {
		panic("could not create postgres pool: " + errPostgres.Error())
	}

	server := server.Serve(ctx, *port, postgresPool)
	server.Run(fmt.Sprintf(":%v", *port))

	defer postgresPool.Close()
}
