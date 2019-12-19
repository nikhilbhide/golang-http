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

func (r *loginUseCase) Login (context context.Context, signUp model.Login) (model.Login,model.Error) {
	//retrieve userid from the repo
	signUp,error:= r.signUpRepo.Create(context, &signUp)
	err:= model.Error{}
	if(error!=nil) {
		err.Message = error.Error()
	}

	return signUp,err
}