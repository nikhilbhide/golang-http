package bo

import (
	"context"
	"github.com/nik/JWTDemo/model"
)

// Usecase represent the article's usecases
type LoginUseCase interface {
	Login(ctx context.Context, signup *model.Login) (*model.Login, *model.Error)
	FetchUserByEmail(email string)(*model.Login, *model.Error)
}