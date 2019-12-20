package bo

import (
	"context"
	"github.com/nik/JWTDemo/model"
	repsitory "github.com/nik/JWTDemo/repository"
)

type loginUseCase struct {
	signUpRepo    repsitory.LoginRepo
}

func NewLoginUseCase(useCaseInstance repsitory.LoginRepo) LoginUseCase {
	return &loginUseCase{
		signUpRepo: useCaseInstance,
	}
}

func (r *loginUseCase) Login (context context.Context, signUp *model.Login) (*model.Login,*model.Error) {
	//retrieve userid from the repo
	signUp, error := r.signUpRepo.Create(context, signUp)

	if (error != nil) {
		err := &model.Error{error.Error()}
		return nil, err
	} else {
		return signUp, nil
	}
}

func (r *loginUseCase) FetchUserByEmail (email string) (*model.Login,*model.Error) {
	//retrieve userid from the repo
	signUp, error := r.signUpRepo.GetUserByEmail(email)

	if (error != nil) {
		err := &model.Error{error.Error()}
		return nil, err
	} else {
		return signUp, nil
	}
}