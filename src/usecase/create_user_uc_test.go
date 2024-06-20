package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"

	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/interfaces/handlerimpl"
	"itmrchow/go-project/user/src/usecase/handler"
	"itmrchow/go-project/user/src/usecase/repo"
)

const (
	ExistFuncStr = "ExistsByAccountOrEmailOrPhone"
)

func TestCreateUserSuite(t *testing.T) {
	suite.Run(t, new(CreateUserUCTestSuite))
}

type CreateUserUCTestSuite struct {
	suite.Suite
	repoMock          *repo.UserRepoMock
	usecase           *CreateUserUseCase
	encryptionHandler handler.EncryptionHandler
}

func (s *CreateUserUCTestSuite) SetupTest() {
	s.repoMock = &repo.UserRepoMock{}
	s.encryptionHandler = &handlerimpl.BcryptHandler{}
	s.usecase = NewCreateUserUseCase(s.repoMock, s.encryptionHandler)
}

func (s *CreateUserUCTestSuite) Test_CreateUser_UserExists() {
	type args struct {
		input CreateUserInput
	}
	type test struct {
		name       string
		args       args
		mockFunc   func(repoMock *repo.UserRepoMock)
		assertFunc func(*CreateUserOutput, error)
	}

	testcase := &test{
		name: "user_exists",
		args: args{
			input: CreateUserInput{
				Account: "test",
				Email:   "XXXXXXXXXXXXX",
				Phone:   "1234567890",
			},
		},
		mockFunc: func(repoMock *repo.UserRepoMock) {
			repoMock.On(ExistFuncStr, mock.Anything, mock.Anything, mock.Anything).Return(false, ErrUserAlreadyExists)
		},
		assertFunc: func(got *CreateUserOutput, err error) {
			s.Assert().Nil(got)
			s.Assert().ErrorIs(err, ErrUserAlreadyExists)

			s.repoMock.AssertNotCalled(s.T(), "Create")
		},
	}

	s.Run(testcase.name, func() {
		testcase.mockFunc(s.repoMock)
		got, err := s.usecase.CreateUser(testcase.args.input)
		testcase.assertFunc(got, err)
	})
}

func (s *CreateUserUCTestSuite) Test_CreateUser_QueryUserError() {
	type args struct {
		input CreateUserInput
	}
	type test struct {
		name       string
		args       args
		mockFunc   func(repoMock *repo.UserRepoMock)
		assertFunc func(*CreateUserOutput, error)
	}

	testcase := &test{
		name: "query_user_error",
		args: args{
			input: CreateUserInput{
				UserName: "UserName",
				Account:  "test",
				Password: "XXXXXXXX",
				Email:    "XXXXXXXXXXXXX",
				Phone:    "1234567890",
			},
		},
		mockFunc: func(repoMock *repo.UserRepoMock) {
			repoMock.On(ExistFuncStr, mock.Anything, mock.Anything, mock.Anything).Return(true, nil)
		},
		assertFunc: func(got *CreateUserOutput, err error) {
			s.Assert().Nil(got)
			s.Assert().Error(err)

			s.repoMock.AssertNotCalled(s.T(), "Create")
		},
	}

	s.Run(testcase.name, func() {
		testcase.mockFunc(s.repoMock)
		got, err := s.usecase.CreateUser(testcase.args.input)
		testcase.assertFunc(got, err)
	})
}

func (s *CreateUserUCTestSuite) Test_CreateUser_HashPasswordError() {
	type args struct {
		input CreateUserInput
	}
	type test struct {
		name       string
		args       args
		mockFunc   func(repoMock *repo.UserRepoMock)
		assertFunc func(*CreateUserOutput, error)
	}

	testcase := &test{
		name: "query_user_error",
		args: args{
			input: CreateUserInput{
				UserName: "UserName",
				Account:  "test",
				Password: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Email:    "XXXXXXXXX",
				Phone:    "1234567890",
			},
		},
		mockFunc: func(repoMock *repo.UserRepoMock) {
			repoMock.On(ExistFuncStr, mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
		},
		assertFunc: func(got *CreateUserOutput, err error) {
			s.Assert().Nil(got)
			s.Assert().ErrorIs(err, bcrypt.ErrPasswordTooLong)

			s.repoMock.AssertNotCalled(s.T(), "Create")
		},
	}

	s.Run(testcase.name, func() {
		testcase.mockFunc(s.repoMock)
		got, err := s.usecase.CreateUser(testcase.args.input)
		testcase.assertFunc(got, err)
	})
}

func (s *CreateUserUCTestSuite) Test_CreateUser_InsertDbFail() {
	type args struct {
		input CreateUserInput
	}
	type test struct {
		name       string
		args       args
		mockFunc   func(repoMock *repo.UserRepoMock)
		assertFunc func(*CreateUserOutput, error)
	}

	testcase := &test{
		name: "query_user_error",
		args: args{
			input: CreateUserInput{
				UserName: "UserName",
				Account:  "test",
				Password: "XXXXXXXXX",
				Email:    "XXXXXXXXX",
				Phone:    "1234567890",
			},
		},
		mockFunc: func(repoMock *repo.UserRepoMock) {
			repoMock.On(ExistFuncStr, mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
			repoMock.On("Create", mock.MatchedBy(func(user *domain.User) bool {
				return s.Assert().NotNil(user.Id) && s.Assert().Equal("UserName", user.UserName)
			})).Return(errors.New("some error"))
		},
		assertFunc: func(got *CreateUserOutput, err error) {
			s.Assert().Nil(got)
			s.Assert().Error(err)

			s.repoMock.AssertCalled(s.T(), "Create", mock.Anything)
		},
	}

	s.Run(testcase.name, func() {
		testcase.mockFunc(s.repoMock)
		got, err := s.usecase.CreateUser(testcase.args.input)
		testcase.assertFunc(got, err)
	})
}

func (s *CreateUserUCTestSuite) Test_CreateUser_Pass() {
	type args struct {
		input CreateUserInput
	}
	type test struct {
		name       string
		args       args
		mockFunc   func(repoMock *repo.UserRepoMock)
		assertFunc func(*CreateUserOutput, error)
	}

	testInput := CreateUserInput{
		UserName: "UserName",
		Account:  "test",
		Password: "XXXXXXXXX",
		Email:    "XXXXXXXXX",
		Phone:    "1234567890",
	}

	testcase := &test{
		name: "query_user_error",
		args: args{
			input: testInput,
		},
		mockFunc: func(repoMock *repo.UserRepoMock) {
			repoMock.On(ExistFuncStr, mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
			repoMock.On("Create", mock.MatchedBy(func(user *domain.User) bool {
				return s.Assert().NotNil(user.Id) && s.Assert().Equal("UserName", user.UserName) && s.Assert().NotNil(user.Password)
			})).Return(nil)
		},
		assertFunc: func(got *CreateUserOutput, err error) {
			s.Assert().NotNil(got)
			s.Assert().Nil(err)

			s.repoMock.AssertCalled(s.T(), "Create", mock.Anything)

			s.Assert().NotNil(got.Id)
			s.Assert().Equal(testInput.UserName, got.UserName)
			s.Assert().Equal(testInput.Account, got.Account)
			s.Assert().Equal(testInput.Email, got.Email)
			s.Assert().Equal(testInput.Phone, got.Phone)

		},
	}

	s.Run(testcase.name, func() {
		testcase.mockFunc(s.repoMock)
		got, err := s.usecase.CreateUser(testcase.args.input)
		testcase.assertFunc(got, err)
	})
}
