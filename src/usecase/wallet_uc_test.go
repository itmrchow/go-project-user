package usecase

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain/enum"
	"itmrchow/go-project/user/src/infrastructure/api/reqdto"
	repo "itmrchow/go-project/user/src/usecase/repo/mocks"
)

func TestWalletSuite(t *testing.T) {
	suite.Run(t, new(WalletTestSuite))
}

type WalletTestSuite struct {
	suite.Suite
	usecase  *WalletUseCase
	repoMock *repo.MockWalletRepo
	authUser reqdto.AuthUser
}

func (s *WalletTestSuite) SetupTest() {
	s.repoMock = &repo.MockWalletRepo{}
	s.usecase = NewWalletUseCase(s.repoMock)
}

func (s *WalletTestSuite) Test_CreateWallet_WalletExist() {
	type args struct {
		input *CreateWalletInput
	}
	type test struct {
		name       string
		args       args
		mockFunc   func()
		assertFunc func(*CreateWalletOutput, error)
	}

	testcase := &test{
		name: "wallet_exist",
		args: args{
			input: &CreateWalletInput{
				UserId:     "Jeff",
				WalletType: enum.WalletType.PLATFORM,
				Currency:   enum.Currency.PHP,
				Balance:    10.0,
			},
		},
		mockFunc: func() {
			s.repoMock.On("Create", mock.Anything).Return(gorm.ErrDuplicatedKey)
		},
		assertFunc: func(got *CreateWalletOutput, err error) {
			s.Assert().Nil(got)
			s.Assert().ErrorIs(err, gorm.ErrDuplicatedKey)
		},
	}

	s.Run(testcase.name, func() {
		// mock
		testcase.mockFunc()

		// execute
		got, err := s.usecase.CreateWallet(testcase.args.input, s.authUser)

		// assert
		testcase.assertFunc(got, err)
	})
}

func (s *WalletTestSuite) Test_CreateWallet_Success() {
	type args struct {
		input *CreateWalletInput
	}
	type test struct {
		name       string
		args       args
		mockFunc   func()
		assertFunc func(*CreateWalletOutput, error)
	}

	testArgs := args{
		input: &CreateWalletInput{
			UserId:     "Jeff",
			WalletType: enum.WalletType.PLATFORM,
			Currency:   enum.Currency.PHP,
			Balance:    10.0,
		},
	}

	testcase := &test{
		name: "wallet_success",
		args: testArgs,
		mockFunc: func() {
			s.repoMock.On("Create", mock.Anything).Return(nil)
		},
		assertFunc: func(got *CreateWalletOutput, err error) {
			s.Assert().Nil(err)

			s.Assert().Equal(testArgs.input.UserId, got.UserId)
			s.Assert().Equal(testArgs.input.WalletType, got.WalletType)
			s.Assert().Equal(testArgs.input.Currency, got.Currency)
			s.Assert().Equal(
				testArgs.input.Balance,
				got.Balance,
			)
			s.Assert().Equal(s.authUser.Id, got.CreatedBy)
			s.Assert().Equal(s.authUser.Id, got.UpdatedBy)
		},
	}

	s.Run(testcase.name, func() {
		// mock
		testcase.mockFunc()

		// execute
		got, err := s.usecase.CreateWallet(testcase.args.input, s.authUser)

		// assert
		testcase.assertFunc(got, err)
	})
}
