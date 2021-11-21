package main

import (
	"bibit/app/controller"
	"bibit/app/dao"
	"bibit/app/infrastructure/connection"
	"bibit/app/infrastructure/server/grpc"
	"bibit/app/infrastructure/server/rest"
	"bibit/app/infrastructure/server/swagger"
	"bibit/app/service/movie_omdb"
	"bibit/cmd/server/main/docs"
	"bibit/internal/config"
	"bytes"
	"context"
	"io"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfig()
	logsAPi := dao.NewLogApiClient(connection.DbConn)
	client := dao.NewMovieOmdbCLient(logsAPi)
	svc := movie_omdb.NewMovieImdbServiceImpl(client)
	controller := controller.NewMovieIMdbController(svc)

	swJson := "./app/infrastructure/swagger/movie.swagger.json"
	file, _ := os.Open(swJson)
	contentBytes := bytes.NewBuffer(nil)
	_, errReadFile := io.Copy(contentBytes, file)
	if errReadFile != nil {
		log.Fatal("Error Read Swagger File")
	}
	docs.SetDoc(contentBytes.String())
	docs.RegisterSwagger()

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
