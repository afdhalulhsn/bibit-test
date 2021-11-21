package grpc

import (
	"bibit/app/controller"
	"bibit/app/dao"
	"bibit/app/infrastructure/connection"
	proto "bibit/app/infrastructure/grpc/proto/movie"
	"bibit/app/service/movie_omdb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func RunServer(ctx context.Context, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	// register service
	server := grpc.NewServer()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	RegisterServer(server)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")
			server.GracefulStop()

			<-ctx.Done()
		}
	}()
	// start gRPC server
	log.Println("starting gRPC server...", port)
	return server.Serve(listen)
}

func RegisterServer(server *grpc.Server){
	logsAPi := dao.NewLogApiClient(connection.DbConn)
	client := dao.NewMovieOmdbCLient(logsAPi)
	svc := movie_omdb.NewMovieImdbServiceImpl(client)
	controller := controller.NewMovieIMdbController(svc)
	proto.RegisterOmdbMovieServiceServer(server, controller)
}
