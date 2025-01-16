package rpcserver

import (
	"context"
	"encoding/json"
	"project/common"
	"project/internal/user/entity"

	"project/proto/user"

	"github.com/btcsuite/btcutil/base58"
)

type UserStore interface {
	GetUsers(ctx context.Context, ids []int) ([]entity.SimpleUser, error)
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*entity.User, error)
	CreateUser(ctx context.Context, data *entity.UserCreation) error
}
type grpcService struct {
	dbStore UserStore
}

func NewGRPCSerivce(dbStore UserStore) *grpcService {
	return &grpcService{dbStore: dbStore}
}

func (s *grpcService) GetUserById(context.Context, *user.GetUserByIdReq) (*user.PublicUserInfoResp, error) {
	return nil, nil
}

func (s *grpcService) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.NewUserIdResp, error) {
	newUserData := entity.NewUserForCreation(req.FirstName, req.LastName, req.Email, req.Avatar)

	if err := s.dbStore.CreateUser(ctx, &newUserData); err != nil {
		return nil, common.ErrInternal(err)
	}

	return &user.NewUserIdResp{Id: int32(newUserData.ID)}, nil
}

func (s *grpcService) GetUsersByIds(ctx context.Context, request *user.GetUsersByIdsReq) (*user.PublicUsersInfoResp, error) {
	userIds := make([]int, len(request.GetIds()))

	for i := range userIds {
		userIds[i] = int(request.GetIds()[i])
	}

	rs, err := s.dbStore.GetUsers(ctx, userIds)

	if err != nil {
		return nil, err
	}

	users := make([]*user.PublicUserInfo, len(rs))

	for i, item := range rs {
		//item.Mask(common.MaskTypeUser)

		userAvatar, _ := json.Marshal(item.Avatar)
		avatarData := base58.Encode(userAvatar)

		users[i] = &user.PublicUserInfo{
			// Id:        item.FakeId(),
			Id:        int32(item.ID),
			FirstName: item.FirstName,
			LastName:  item.LastName,
			Avatar:    avatarData,
			// Role:      item.Role,
		}
	}

	return &user.PublicUsersInfoResp{Users: users}, nil
}
