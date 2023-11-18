package grpcServer

import (
	"context"
	"log"
	"net"

	pb "github.com/jeka-designer/proto/gen/go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type usecases interface {
	RandomTaksa(ctx context.Context) (url string, err error)
}

type server struct {
	// type embedded to comply with Google lib
	pb.UnimplementedImagesServiceServer
	uc usecases
}

func New(uc usecases, port string) (*grpc.Server, error) {
	listener, err := net.Listen("tcp", port)

	if err != nil {
		return nil, err
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterImagesServiceServer(s, &server{uc: uc})

	log.Println("starting images service on port", port)

	if err := s.Serve(listener); err != nil {
		return nil, err
	}

	return s, err
}

func (s server) GetRandomTaksa(ctx context.Context, request *pb.GetRandomTaksaRequest) (*pb.GetRandomTaksaResponse, error) {
	log.Println("Random taksa called")

	url, err := s.uc.RandomTaksa(ctx)

	if err != nil {
		return nil, err
	}

	log.Println("Taksa fetched successfully")

	return &pb.GetRandomTaksaResponse{Url: url}, nil
}
