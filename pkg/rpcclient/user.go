package rpcclient

import (
	"context"
	"project/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type User struct {
	conn   grpc.ClientConnInterface
	Client user.UserServiceClient
}

// CreateUser implements usecase.UserRpcClient.
func (u *User) CreateUser(ctx context.Context, firstName string, lastName string, email string, avatar string) (newId int, err error) {
	panic("unimplemented")
}

// NewUser initializes and returns a User instance based on the provided service discovery registry.
func NewUser(ctx context.Context) *User {
	conn, err := grpc.NewClient(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := user.NewUserServiceClient(conn)
	return &User{Client: client, conn: conn}
}

// UserRpcClient represents the structure for a User RPC client.
type UserRpcClient User
