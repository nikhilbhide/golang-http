package bo

import (
	"context"
	"github.com/nik/JWTDemo/driver/postgres"
	"github.com/nik/JWTDemo/repository/users"
	"reflect"
	"testing"

	"github.com/nik/JWTDemo/model"
	repsitory "github.com/nik/JWTDemo/repository"
)

func TestNewLoginUseCase(t *testing.T) {
	type args struct {
		useCaseInstance repsitory.LoginRepo
	}
	var usecasePostgresInstance = users.NewLoginPostGresRepo(postgres.InitDB("postgres://nenyejsk:80PuDCmjzIAlMJb8xLaS7rgBwE7gXHeN@rajje.db.elephantsql.com:5432/nenyejsk"))
	var useCaseInstance = NewLoginUseCase(usecasePostgresInstance)

	tests := []struct {
		name string
		args args
		want LoginUseCase
	}{
		{
			"unit_test_for_identical_objects",
			args{
				usecasePostgresInstance,
			},
			useCaseInstance,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoginUseCase(tt.args.useCaseInstance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loginUseCase_Login(t *testing.T) {
	type args struct {
		context context.Context
		signUp  *model.Login
	}

	var usecasePostgresInstance = users.NewLoginPostGresRepo(postgres.InitDB("postgres://nenyejsk:80PuDCmjzIAlMJb8xLaS7rgBwE7gXHeN@rajje.db.elephantsql.com:5432/nenyejsk"))
	var useCaseInstance = NewLoginUseCase(usecasePostgresInstance)

	tests := []struct {
		name  string
		r     LoginUseCase
		args  args
		want  *model.Login
		want1 *model.Error
	}{
		{
			"unit_test_successful_login",
			useCaseInstance,
			args{
				signUp:&model.Login{
					Password:"jkl1",
					Email:"jkl1@test.com",
				},
				context:nil,
			},
			&model.Login{
					Password:"jkl1",
					Email:"jkl1@test.com",
					UserID:"26",
				},
				nil,
			},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.r.Login(tt.args.context, tt.args.signUp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loginUseCase.Login() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("loginUseCase.Login() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_loginUseCase_FetchUserByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name  string
		r     *loginUseCase
		args  args
		want  *model.Login
		want1 *model.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.r.FetchUserByEmail(tt.args.email)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loginUseCase.FetchUserByEmail() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("loginUseCase.FetchUserByEmail() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}