package auth

import (
	"car-rent/internal/common"
	"car-rent/internal/entity"
	"car-rent/internal/presentations"
	"car-rent/internal/repositories"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type Contract interface {
	Authorization(ctx context.Context, payload entity.Authorization) (*presentations.Authorization, error)
}

type business struct {
	repo *repositories.Repository
	jwt  common.JwtCode
}

func NewBusiness(repo *repositories.Repository) Contract {
	return &business{
		repo: repo,
		jwt:  common.NewJwt(),
	}
}

func (b *business) Authorization(ctx context.Context, payload entity.Authorization) (*presentations.Authorization, error) {

	var (
		users *presentations.Users
		err   error
	)

	users, err = b.repo.Users.GetUserByUsername(ctx, payload.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(payload.Password))
	if err != nil {
		return nil, common.ErrUnauthorized
	}

	accesstoken, err := b.jwt.GenerateAuthorizartionCode(entity.Claim{
		UserID:   users.UserID,
		Username: users.Username,
		IsAdmin:  users.IsAdmin,
	})
	if err != nil {
		return nil, err
	}

	return &presentations.Authorization{
		AccessToken: accesstoken,
	}, nil

}
