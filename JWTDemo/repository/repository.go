package repsitory

import (
	"context"
	"github.com/nik/JWTDemo/model"
)

// SignupRepo
type LoginRepo interface {
	Fetch(ctx context.Context, num int64) ([]*model.Login, error)
	GetByUserID(ctx context.Context, id int64) (*model.Login, error)
	GetUserByEmail(email string) (*model.Login, error)
	Create(ctx context.Context, p *model.Login) (*model.Login, error)
	Update(ctx context.Context, p *model.Login) (*model.Login, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

// TokenRepo
type TokenRepo interface {
	Fetch(ctx context.Context, num int64) ([]*model.Token, error)
	GetByID(ctx context.Context, id int64) (*model.Token, error)
	Create(ctx context.Context, p *model.Token) (int64, error)
	Update(ctx context.Context, p *model.Token) (*model.Token, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

// ErrorRepo
type ErrorRepo interface {
	Fetch(ctx context.Context, num int64) ([]*model.Error, error)
	GetByID(ctx context.Context, id int64) (*model.Error, error)
	Create(ctx context.Context, p *model.Error) (int64, error)
	Update(ctx context.Context, p *model.Error) (*model.Error, error)
	Delete(ctx context.Context, id int64) (bool, error)
}