package delivery

import (
	"context"
	pb "github.com/jeka-designer/proto/gen/go"
	"log"
	"net"
	"os"

	"youtube/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	// type embedded to comply with Google lib
	pb.UnimplementedYoutubeServiceServer
	uc models.YoutubeUsecases
}

func NewGRPCServer(uc models.YoutubeUsecases) error {
	port := ":" + os.Getenv("GRPC_PORT")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterYoutubeServiceServer(s, &server{uc: uc})

	log.Println("starting grpc server on port", port)

	if err := s.Serve(listener); err != nil {
		return err
	}

	return err
}

func (s *server) GetRandomVideo(ctx context.Context, request *pb.GetRandomVideoRequest) (*pb.GetRandomVideoResponse, error) {
	var url, caption string
	var err error

	url, caption, err = s.uc.GetRandomVideoUrl()

	if err != nil {
		return nil, err
	}

	log.Printf("Url fetched successfully %v", url)

	return &pb.GetRandomVideoResponse{Url: url, Caption: caption}, nil
}
