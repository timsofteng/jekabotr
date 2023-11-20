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
	RandomImg(ctx context.Context, query string) (url string, err error)
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

	log.Println("starting grpc server on port", port)

	if err := s.Serve(listener); err != nil {
		return nil, err
	}

	return s, err
}

func (s server) GetRandomImg(ctx context.Context, req *pb.GetRandomImgRequest) (*pb.GetRandomImgResponse, error) {
	log.Println("Random img called")

	query := req.Query

	url, err := s.uc.RandomImg(ctx, query)

	if err != nil {
		return nil, err
	}

	log.Println("Random img fetched successfully")

	return &pb.GetRandomImgResponse{Url: url}, nil
}
