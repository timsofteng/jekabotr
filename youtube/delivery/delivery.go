package delivery

import (
	"context"
	"log"
	"net"
	pb "github.com/timsofteng/jekabot"

	models "youtube/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	// type embedded to comply with Google lib
	pb.UnimplementedYoutubeServer
	uc models.YoutubeUsecases
}

func NewGRPCServer(uc models.YoutubeUsecases) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterYoutubeServer(s, &server{uc: uc})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) GetRandomVideo(ctx context.Context, request *pb.RandomVideoRequest) (*pb.RandomVideoResponse, error) {
	log.Println("Random video called")
	var url string
	var err error

	url, err = s.uc.GetRandomVideoUrl()

	if err != nil {
		log.Println("Failded to get video")
	}

	log.Printf("Url fetched successfully %v", url)

	return &pb.RandomVideoResponse{Url: url}, nil
}
