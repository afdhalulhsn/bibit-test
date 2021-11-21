package main

import (
	"bibit/app/infrastructure/server/grpc"
	"bibit/app/infrastructure/server/rest"
	"bibit/app/infrastructure/server/swagger"
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
	grpc.RunServer(ctx,cfg.Server.Grpc.Port)

}
