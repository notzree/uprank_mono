package client

import (
	proto "github.com/notzree/uprank_mono/uprank-backend/queue-handler/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCInferenceClient(remoteAddr string) (proto.InferenceClient, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := proto.NewInferenceClient(conn)
	return c, nil
}
