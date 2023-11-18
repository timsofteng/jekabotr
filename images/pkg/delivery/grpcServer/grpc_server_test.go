package grpcServer

import "testing"

func TestGRPCServer(t *testing.T) {

	go New(nil, ":8089")

	// if err != nil {
	// 	t.Errorf(`error during starting grpc server, %v`, err)
	// }

	// s.Stop()

}
