package usecase

import (
	"context"
	"encoding/json"
	"project/common"
	"project/modules/auth/entity"

	"github.com/btcsuite/btcutil/base58"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type AuthRepository interface {
	AddNewAuth(ctx context.Context, data *entity.Auth) error
	GetAuth(ctx context.Context, email string) (*entity.Auth, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, firstName, lastName, email, avatar string) (newId int, err error)
}

type Hasher interface {
	RandomStr(length int) (string, error)
	HashPassword(salt, password string) (string, error)
	CompareHashPassword(hashedPassword, salt, password string) bool
}

type JWTProviderComponent interface {
	IssueToken(ctx context.Context, id, sub string) (token string, expSecs int, err error)
	ParseToken(ctx context.Context, tokenString string) (claims *jwt.RegisteredClaims, err error)
}

type business struct {
	authRepository AuthRepository
	userRepository UserRepository
	jwtProvider    JWTProviderComponent
	hasher         Hasher
}

func NewBusiness(repository AuthRepository, userRepository UserRepository,
	jwtProvider JWTProviderComponent, hasher Hasher) *business {
	return &business{
		authRepository: repository,
		userRepository: userRepository,
		jwtProvider:    jwtProvider,
		hasher:         hasher,
	}
}

func (biz *business) Login(ctx context.Context, data *entity.AuthEmailPassword) (*entity.TokenResponse, error) {
	if err := data.Validate(); err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	authData, err := biz.authRepository.GetAuth(ctx, data.Email)

	if err != nil {
		return nil, common.ErrInvalidRequestWithMsg(err, entity.ErrLoginFailed.Error())
	}

	if !biz.hasher.CompareHashPassword(authData.Password, authData.Salt, data.Password) {
		return nil, common.ErrInvalidRequest(entity.ErrLoginFailed)
	}

	uid := common.NewUID(uint32(authData.UserId), 1, 1)
	sub := uid.String()
	tid := uuid.New().String()

	tokenStr, expSecs, err := biz.jwtProvider.IssueToken(ctx, tid, sub)

	if err != nil {
		return nil, common.ErrInternalWithMsg(err, entity.ErrLoginFailed.Error())
	}

	return &entity.TokenResponse{
		AccessToken: entity.Token{
			Token:     tokenStr,
			ExpiredIn: expSecs,
		},
	}, nil
}

func (biz *business) Register(ctx context.Context, data *entity.AuthRegister) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	_, err := biz.authRepository.GetAuth(ctx, data.Email)

	if err == nil {
		return common.ErrInvalidRequest(entity.ErrEmailHasExisted)
	} else if err != common.ErrRecordNotFound {
		return common.ErrInternalWithMsg(err, entity.ErrCannotRegister.Error())
	}

	avatar, err := json.Marshal(data.Avatar)
	if err != nil {
		return common.ErrInternalWithMsg(err, entity.ErrCannotRegister.Error())
	}
	avatarData := base58.Encode(avatar)

	newUserId, err := biz.userRepository.CreateUser(ctx, data.FirstName, data.LastName, data.Email, avatarData)

	if err != nil {
		return common.ErrInternalWithMsg(err, entity.ErrCannotRegister.Error())
	}

	salt, err := biz.hasher.RandomStr(16)

	if err != nil {
		return common.ErrInternalWithMsg(err, entity.ErrCannotRegister.Error())
	}

	passHashed, err := biz.hasher.HashPassword(salt, data.Password)

	if err != nil {
		return common.ErrInternalWithMsg(err, entity.ErrCannotRegister.Error())
	}

	newAuth := entity.NewAuthWithEmailPassword(newUserId, data.Email, salt, passHashed)

	if err := biz.authRepository.AddNewAuth(ctx, &newAuth); err != nil {
		return common.ErrInternalWithMsg(err, entity.ErrCannotRegister.Error())
	}

	return nil
}

func (biz *business) IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error) {
	claims, err := biz.jwtProvider.ParseToken(ctx, accessToken)

	if err != nil {
		return nil, common.ErrUnauthorized(err)
	}

	return claims, nil
}
