package delivery

import (
	"context"
	"log"
	"net"
	"os"
	pb "proto"

	models "youtube/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const YT_LINK_CAPTION = "Взгляните на это видео:\n\n"

type server struct {
	// type embedded to comply with Google lib
	pb.UnimplementedYoutubeServiceServer
	uc models.YoutubeUsecases
}

func NewGRPCServer(uc models.YoutubeUsecases) {
	port := ":" + os.Getenv("YOUTUBE_GRPC_PORT")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterYoutubeServiceServer(s, &server{uc: uc})

	log.Println("starting youtube service on port", port)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetRandomVideo(ctx context.Context, request *pb.GetRandomVideoRequest) (*pb.GetRandomVideoResponse, error) {
	log.Println("Random video called")
	var url string
	var err error

	url, err = s.uc.GetRandomVideoUrl()

	if err != nil {
		log.Println("Failded to get video")
	}

	log.Printf("Url fetched successfully %v", url)

	return &pb.GetRandomVideoResponse{Url: url, Caption: YT_LINK_CAPTION}, nil
}
