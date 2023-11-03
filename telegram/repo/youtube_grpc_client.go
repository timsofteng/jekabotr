package repo

import (
	"log"
	"os"

	pb "proto"

	"google.golang.org/grpc"
)

func NewYoutubeGRPCClient() (pb.YoutubeServiceClient, *grpc.ClientConn, error) {
	port := ":" + os.Getenv("YOUTUBE_GRPC_PORT")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	c := pb.NewYoutubeServiceClient(conn)

	log.Println("youtube grpc client started")

	return c, conn, err
}
