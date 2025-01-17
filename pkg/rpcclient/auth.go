package rpcclient

import (
	"context"
	"log"

	"project/proto/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuth(ctx context.Context) *Auth {
	// conn, err := discov.GetConn(context.Background(), config.Config.RpcRegisterName.OpenImAuthName)
	conn, err := grpc.NewClient(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	client := auth.NewAuthServiceClient(conn)
	return &Auth{conn: conn, Client: client}
}

type Auth struct {
	conn   grpc.ClientConnInterface
	Client auth.AuthServiceClient
}
