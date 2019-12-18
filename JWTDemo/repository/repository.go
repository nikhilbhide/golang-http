package repsitory

import (
	"context"
	"github.com/nik/JWTDemo/model"
)

// SignupRepo
type PostRepo interface {
	Fetch(ctx context.Context, num int64) ([]*model.Signup, error)
	GetByID(ctx context.Context, id int64) (*model.Signup, error)
	Create(ctx context.Context, p *model.Signup) (int64, error)
	Update(ctx context.Context, p *model.Signup) (*model.Signup, error)
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