package main

import (
	"errors"
	"log"
	"os"
	"runtime/debug"

	pb "proto"

	"google.golang.org/grpc"
)

func getCurrentModuleName() (string, error) {
	bi, ok := debug.ReadBuildInfo()

	if !ok {
		return "", errors.New("failed to get current module name")
	}

	return bi.Path, nil

}

func NewYoutubeGRPCClient() (pb.YoutubeServiceClient, *grpc.ClientConn, error) {
	port := ":" + os.Getenv("YOUTUBE_GRPC_PORT")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	c := pb.NewYoutubeServiceClient(conn)

	currentModuleName, err := getCurrentModuleName()

	if err != nil {
		log.Println("Failed to read current module name")
	}

	log.Printf("%v: youtube grpc client started", currentModuleName)

	return c, conn, err
}
