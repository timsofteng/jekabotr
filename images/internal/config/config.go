package config

import (
	"errors"
	"os"
)

type config struct {
	GRPCPort string
}

func New() (*config, error) {
	grpcPort := ":" + os.Getenv("GRPC_PORT")

	if len(grpcPort) < 1 {
		return nil, errors.New("port doesnt specified")
	}

	return &config{GRPCPort: grpcPort}, nil

}
