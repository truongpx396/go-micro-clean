package rpcclient

import (
	"context"
	"go-micro-clean/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type User struct {
	conn   grpc.ClientConnInterface
	Client user.UserServiceClient
}

// CreateUser implements usecase.UserRpcClient.
func (u *User) CreateUser(ctx context.Context, firstName string, lastName string, email string, avatar string) (newId int, err error) {
	//panic("unimplemented")
	data, err := u.Client.CreateUser(ctx, &user.CreateUserReq{FirstName: firstName, LastName: lastName, Email: email, Avatar: avatar})
	if err != nil {
		return -1, err
	}
	return int(data.Id), nil
}

// NewUser initializes and returns a User instance based on the provided service discovery registry.
func NewUser(ctx context.Context) *User {
	conn, err := grpc.NewClient(":10111", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := user.NewUserServiceClient(conn)
	return &User{Client: client, conn: conn}
}

// UserRpcClient represents the structure for a User RPC client.
type UserRpcClient User
