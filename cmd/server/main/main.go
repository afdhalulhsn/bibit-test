package main

import (
	"TEST/bibi_test/app/controller"
	"TEST/bibi_test/app/dao"
	"TEST/bibi_test/app/infrastructure/connection"
	"TEST/bibi_test/app/infrastructure/server/grpc"
	"TEST/bibi_test/app/infrastructure/server/rest"
	"TEST/bibi_test/app/service/movie_omdb"
	"TEST/bibi_test/internal/config"
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
	grpc.RunServer(ctx, controller, cfg.Server.Grpc.Port)

}
