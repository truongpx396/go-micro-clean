package rpcclient

import (
	"context"
	pb "project/proto/gen"

	"github.com/pkg/errors"
)

type rpcUserClient struct {
	client pb.UserServiceClient
}

func NewClient(client pb.UserServiceClient) *rpcUserClient {
	return &rpcUserClient{client: client}
}

func (c *rpcUserClient) CreateUser(ctx context.Context, firstName, lastName, email, avatar string) (newId int, err error) {
	resp, err := c.client.CreateUser(ctx, &pb.CreateUserReq{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Avatar:    avatar,
	})

	if err != nil {
		return 0, errors.WithStack(err)
	}

	return int(resp.Id), nil
}
