package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/usecase/handler"
	"itmrchow/go-project/user/src/usecase/repo"
)

func TestGetUserSuite(t *testing.T) {
	suite.Run(t, new(GetUserUCTestSuite))
}

type GetUserUCTestSuite struct {
	suite.Suite
	repoMock              *repo.UserRepoMock
	encryptionHandlerMock *handler.EncryptionHandlerMock
	usecase               *GetUserUseCase
}

func (s *GetUserUCTestSuite) SetupTest() {
	s.repoMock = &repo.UserRepoMock{}
	s.encryptionHandlerMock = &handler.EncryptionHandlerMock{}
	s.usecase = NewGetUserUseCase(s.repoMock, s.encryptionHandlerMock)
}

func (s *GetUserUCTestSuite) Test_GetUser_NoUser() {
	type args struct {
		userId string
	}
	type test struct {
		name     string
		args     args
		mockFunc func(repoMock *repo.UserRepoMock)
		want     *GetUserOutput
		wantErr  error
	}

	testcase := &test{
		name: "no_user",
		args: args{
			userId: "userId",
		},
		mockFunc: func(repoMock *repo.UserRepoMock) {
			repoMock.On("Get", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
		},
		want:    nil,
		wantErr: nil,
	}

	s.Run(testcase.name, func() {

		s.repoMock.On("Get", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
		got, err := s.usecase.GetUser(testcase.args.userId)

		s.Assert().Nil(got)
		s.Assert().Nil(err)
	})
}

func (s *GetUserUCTestSuite) Test_GetUser_QueryError() {
	type args struct {
		userId string
	}
	type test struct {
		name    string
		args    args
		want    *GetUserOutput
		wantErr error
	}

	testcase := &test{
		name: "query_error",
		args: args{
			userId: "userId",
		},
		want:    nil,
		wantErr: errors.New("some error"),
	}

	s.repoMock.On("Get", mock.Anything).Return(nil, errors.New("some error"))
	got, err := s.usecase.GetUser(testcase.args.userId)

	s.Assert().Nil(got)
	s.Assert().ErrorIs(err, ErrDbFail)
}

func (s *GetUserUCTestSuite) Test_GetUser_HasUser() {
	type args struct {
		userId string
	}
	type test struct {
		name    string
		args    args
		want    *GetUserOutput
		wantErr error
	}

	testcase := &test{
		name: "has_user",
		args: args{
			userId: "has_user",
		},
		want: &GetUserOutput{
			Id:       "has_user",
			UserName: "name",
			Account:  "Account",
			Email:    "Email",
			Phone:    "Phone",
		},
		wantErr: nil,
	}

	s.repoMock.On("Get", mock.Anything).Return(&domain.User{
		Id:       "has_user",
		UserName: "name",
		Account:  "Account",
		Password: "Password",
		Email:    "Email",
		Phone:    "Phone",
	}, nil)
	got, err := s.usecase.GetUser(testcase.args.userId)

	s.Assert().EqualValues(got, testcase.want)
	s.Assert().Nil(err)
}

// Login_Test

func (s *GetUserUCTestSuite) Test_Login_NoUser() {
	type test struct {
		name     string
		args     LoginInput
		mockFunc func(repoMock *repo.UserRepoMock)
		want     string
		wantErr  error
	}

	testcase := &test{
		name: "NoUser",
		args: LoginInput{
			Account:  "Account",
			Email:    "Email",
			Password: "XXXXXXXX",
		},
		mockFunc: func(repoMock *repo.UserRepoMock) {
			repoMock.On("GetByAccountOrEmail", mock.Anything, mock.Anything).Return(nil, gorm.ErrRecordNotFound)
		},
		want:    "",
		wantErr: ErrUserNotExists,
	}

	s.Run(testcase.name, func() {

		// mock
		testcase.mockFunc(s.repoMock)

		// run
		gotToken, err := s.usecase.Login(testcase.args)

		// assert
		s.Assert().Equal("", gotToken)
		s.Assert().ErrorIs(err, ErrUserNotExists)

	})
}

func (s *GetUserUCTestSuite) Test_Login_InvalidPsw() {
	type test struct {
		name     string
		args     LoginInput
		mockFunc func(repoMock *repo.UserRepoMock, encrypMock *handler.EncryptionHandlerMock)
		want     string
		wantErr  error
	}

	testcase := &test{
		name: "InvalidPsw",
		args: LoginInput{
			Account:  "Account",
			Email:    "Email",
			Password: "XXXXXXXX",
		},
		mockFunc: func(repoMock *repo.UserRepoMock, encrypMock *handler.EncryptionHandlerMock) {
			repoMock.On("GetByAccountOrEmail", mock.Anything, mock.Anything).Return(
				&domain.User{
					Id:       "id",
					UserName: "UserName",
					Account:  "Account",
					Password: "12345678",
					Email:    "Email",
					Phone:    "Phone",
				}, nil)

			encrypMock.On("CheckPasswordHash", mock.Anything, mock.Anything).Return(false)
		},
		want:    "",
		wantErr: ErrUnauthorized,
	}

	s.Run(testcase.name, func() {

		// mock
		testcase.mockFunc(s.repoMock, s.encryptionHandlerMock)

		// run
		gotToken, err := s.usecase.Login(testcase.args)

		// assert
		s.Assert().Equal("", gotToken)
		s.Assert().ErrorIs(err, ErrUnauthorized)

	})
}

func (s *GetUserUCTestSuite) Test_Login_Success() {
	type test struct {
		name     string
		args     LoginInput
		mockFunc func(repoMock *repo.UserRepoMock, encrypMock *handler.EncryptionHandlerMock)
		want     string
		wantErr  error
	}

	testcase := &test{
		name: "JWTError",
		args: LoginInput{
			Account:  "Account",
			Email:    "Email",
			Password: "XXXXXXXX",
		},
		mockFunc: func(repoMock *repo.UserRepoMock, encrypMock *handler.EncryptionHandlerMock) {
			repoMock.On("GetByAccountOrEmail", mock.Anything, mock.Anything).Return(
				&domain.User{
					Id:       "id",
					UserName: "UserName",
					Account:  "Account",
					Password: "12345678",
					Email:    "Email",
					Phone:    "Phone",
				}, nil)

			encrypMock.On("CheckPasswordHash", mock.Anything, mock.Anything).Return(true)
		},
		want:    "",
		wantErr: ErrUnauthorized,
	}

	s.Run(testcase.name, func() {

		// mock
		testcase.mockFunc(s.repoMock, s.encryptionHandlerMock)

		// run
		gotToken, err := s.usecase.Login(testcase.args)

		// assert
		s.Assert().NotNil(gotToken)
		s.Assert().Nil(err)
	})
}
