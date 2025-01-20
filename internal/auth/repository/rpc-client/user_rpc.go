package rpcclient

import (
	"context"
	"go-micro-clean/proto/user"

	"github.com/pkg/errors"
)

type rpcUserClient struct {
	client user.UserServiceClient
}

func NewClient(client user.UserServiceClient) *rpcUserClient {
	return &rpcUserClient{client: client}
}

func (c *rpcUserClient) CreateUser(ctx context.Context, firstName, lastName, email, avatar string) (newId int, err error) {
	resp, err := c.client.CreateUser(ctx, &user.CreateUserReq{
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
