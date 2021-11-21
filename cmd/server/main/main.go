package main

import (
	"bibit/app/controller"
	"bibit/app/dao"
	"bibit/app/infrastructure/connection"
	"bibit/app/infrastructure/server/grpc"
	"bibit/app/infrastructure/server/rest"
	"bibit/app/infrastructure/server/swagger"
	"bibit/app/service/movie_omdb"
	"bibit/internal/config"
	"context"
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfig()
	log := dao.NewLogApiClient(connection.DbConn)
	client := dao.NewMovieOmdbCLient(log)
	svc := movie_omdb.NewMovieImdbServiceImpl(client)
	controller := controller.NewMovieIMdbController(svc)

	// run HTTP gateway
	go func() {
		_ = rest.RunRestServer(ctx, cfg.Server.Grpc.Port, cfg.Server.Grpc.RestHost)
	}()
	//Run Swagger
	go func() {
		swagger.RunSwagger()
	}()
	grpc.RunServer(ctx, controller, cfg.Server.Grpc.Port)

}
