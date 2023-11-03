package repo

import (
	"log"
	"os"

	pb "proto"

	"google.golang.org/grpc"
)

func NewImagesGRPCClient() (pb.ImagesServiceClient, *grpc.ClientConn, error) {
	port := ":" + os.Getenv("IMAGES_GRPC_PORT")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	c := pb.NewImagesServiceClient(conn)

	log.Println("images grpc client started")

	return c, conn, err
}
