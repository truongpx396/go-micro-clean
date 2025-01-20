package composers

import (
	"context"
	"log"
	pb "go-micro-clean/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func composeUserRPCClient(serviceCtx context.Context) pb.UserServiceClient {

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.NewClient("", opts)

	if err != nil {
		log.Fatal(err)
	}

	return pb.NewUserServiceClient(clientConn)
}
