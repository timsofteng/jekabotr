package config

import (
	"errors"
	"os"
)

type config struct {
	GRPCPort string
	HTTPPort string
}

func New() (*config, error) {
	grpcPort := ":" + os.Getenv("GRPC_PORT")
	httpPort := ":" + os.Getenv("HTTP_PORT")

	if len(grpcPort) < 1 {
		return nil, errors.New("port doesnt specified")
	}

	return &config{GRPCPort: grpcPort, HTTPPort: httpPort}, nil

}
