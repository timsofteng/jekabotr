package delivery

import (
	"context"
	"log"
	"net"
	"os"
	pb "proto"
	models "images/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	// type embedded to comply with Google lib
	pb.UnimplementedImagesServiceServer
	uc models.ImagesUsecases
}

func NewGRPCServer(uc models.ImagesUsecases) {
	port := ":" + os.Getenv("IMAGES_GRPC_PORT")

	listener, err := net.Listen("tcp", port)

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterImagesServiceServer(s, &server{uc: uc})

	log.Println("starting images service on port", port)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s server) GetRandomTaksa(ctx context.Context, request *pb.GetRandomTaksaRequest) (*pb.GetRandomTaksaResponse, error) {
	log.Println("Random taksa called")

	bin, id, caption, err := s.uc.RandomTaksaGetter()

	if err != nil {
		log.Println("Failded to get video")
	}

	log.Println("Taksa fetched successfully")

	return &pb.GetRandomTaksaResponse{Id: id, Bin: bin, Caption: caption}, nil
}
