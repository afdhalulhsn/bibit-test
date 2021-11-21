package grpc

import (
	proto "bibit/app/infrastructure/grpc/proto/movie"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func RunServer(ctx context.Context, v1API proto.OmdbMovieServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	// register service
	server := grpc.NewServer()
	proto.RegisterOmdbMovieServiceServer(server, v1API)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
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
