package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/usecase/repo"
)

func TestGetUserSuite(t *testing.T) {
	suite.Run(t, new(GetUserUCTestSuite))
}

type GetUserUCTestSuite struct {
	suite.Suite
	repoMock *repo.UserRepoMock
	usecase  *GetUserUseCase
}

func (s *GetUserUCTestSuite) SetupTest() {
	s.repoMock = &repo.UserRepoMock{}
	s.usecase = NewGetUserUseCase(s.repoMock)
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
	s.Assert().EqualError(err, testcase.wantErr.Error())
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
