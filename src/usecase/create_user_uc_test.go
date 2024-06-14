package usecase

import (
	"itmrchow/go-project/user/src/usecase/handler"
	"itmrchow/go-project/user/src/usecase/repo"
	"reflect"
	"testing"
)

func TestNewCreateUserUseCase(t *testing.T) {
	type args struct {
		userRepo          repo.UserRepo
		encryptionHandler handler.EncryptionHandler
	}
	tests := []struct {
		name string
		args args
		want CreateUserUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCreateUserUseCase(tt.args.userRepo, tt.args.encryptionHandler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCreateUserUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUserUseCase_CreateUser(t *testing.T) {
	type fields struct {
		userRepo          repo.UserRepo
		encryptionHandler handler.EncryptionHandler
	}
	type args struct {
		input CreateUserInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateUserOutput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateUserUseCase{
				userRepo:          tt.fields.userRepo,
				encryptionHandler: tt.fields.encryptionHandler,
			}
			got, err := c.CreateUser(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUserUseCase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUserUseCase.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
