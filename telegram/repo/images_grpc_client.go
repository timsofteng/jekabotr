package repo

import (
	"errors"
	"log"
	"os"

	pb "github.com/jeka-designer/proto/gen/go"

	"google.golang.org/grpc"
)

func NewImagesGRPCClient() (pb.ImagesServiceClient, *grpc.ClientConn, error) {
	port := os.Getenv("IMAGES_GRPC_PORT")

	if len(port) < 1 {
		err := errors.New("port env var doesn't provided")
		return nil, nil, err
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		return nil, conn, err
	}

	c := pb.NewImagesServiceClient(conn)

	log.Println("images grpc client started")

	return c, conn, err
}
