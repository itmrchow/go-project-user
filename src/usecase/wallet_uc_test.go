package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
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

// FindWallet

func (s *WalletTestSuite) Test_FindWallet_QueryErr() {
	type args struct {
		input *FindWalletInput
	}
	type test struct {
		name       string
		args       args
		mockFunc   func()
		assertFunc func(*[]FindWalletOutput, error)
	}

	testcase := &test{
		name: "query_error",
		args: args{
			input: &FindWalletInput{},
		},
		mockFunc: func() {
			s.repoMock.On(
				"Find",        //mock func name
				mock.Anything, //mock args
			).Return(nil, errors.New("some error"))
		},
		assertFunc: func(got *[]FindWalletOutput, err error) {
			s.Assert().Nil(got)
			s.Assert().ErrorIs(err, ErrDbFail)
		},
	}

	s.Run(testcase.name, func() {
		// mock
		testcase.mockFunc()

		// execute
		got, err := s.usecase.FindWallet(testcase.args.input)

		// assert
		testcase.assertFunc(got, err)
	})
}

func (s *WalletTestSuite) Test_FindWallet_NoData() {
	type args struct {
		input *FindWalletInput
	}
	type test struct {
		name       string
		args       args
		mockFunc   func()
		assertFunc func(*[]FindWalletOutput, error)
	}

	inputArgs := FindWalletInput{
		UserId:     "Jeff",
		WalletType: "P",
		Currency:   "PHP",
	}

	testcase := &test{
		name: "no_data",
		args: args{
			input: &inputArgs,
		},
		mockFunc: func() {
			s.repoMock.On(
				//mock func name
				"Find",
				//mock args
				mock.MatchedBy(func(i interface{}) bool {
					query, isTypeRight := i.(domain.Wallet)

					if !isTypeRight {
						return false
					}

					// 確認值
					isValueEq := s.Assert().Equal(inputArgs.UserId, query.UserId) &&
						s.Assert().Equal(inputArgs.Currency, query.Currency) &&
						s.Assert().Equal(inputArgs.WalletType, query.WalletType)
					return isValueEq
				}),
			).Return([]domain.Wallet{}, nil)
		},
		assertFunc: func(got *[]FindWalletOutput, err error) {
			s.Assert().Empty(got)
			s.Assert().Nil(err)

			s.repoMock.AssertExpectations(s.T())
		},
	}

	s.Run(testcase.name, func() {
		// mock
		testcase.mockFunc()

		// execute
		got, err := s.usecase.FindWallet(testcase.args.input)

		// assert
		testcase.assertFunc(got, err)
	})
}

func (s *WalletTestSuite) Test_FindWallet_Success() {

	type args struct {
		input *FindWalletInput
	}
	type test struct {
		name       string
		args       args
		mockFunc   func()
		assertFunc func(*[]FindWalletOutput, error)
	}

	inputArgs := FindWalletInput{
		UserId:     "Jeff",
		WalletType: "P",
		Currency:   "PHP",
	}

	testcase := &test{
		name: "success",
		args: args{
			input: &inputArgs,
		},
		mockFunc: func() {
			s.repoMock.On(
				//mock func name
				"Find",
				//mock args
				mock.MatchedBy(func(i interface{}) bool {
					query, isTypeRight := i.(domain.Wallet)

					if !isTypeRight {
						return false
					}

					// 確認值
					isValueEq := s.Assert().Equal(inputArgs.UserId, query.UserId) &&
						s.Assert().Equal(inputArgs.Currency, query.Currency) &&
						s.Assert().Equal(inputArgs.WalletType, query.WalletType)
					return isValueEq
				}),
			).Return([]domain.Wallet{
				{
					UserId:     "Jeff",
					WalletType: enum.WalletType.PLATFORM,
					Currency:   enum.Currency.PHP,
					Balance:    10.0,
					DefaultModel: domain.DefaultModel{
						CreatedBy: "Jeff",
						UpdatedBy: "Jeff",
					},
				},
			}, nil)
		},
		assertFunc: func(gotSlice *[]FindWalletOutput, err error) {
			s.repoMock.AssertExpectations(s.T())

			s.Assert().Len(*gotSlice, 1)
			s.Assert().Nil(err)

			got := (*gotSlice)[0]
			s.Assert().Equal(got.UserId, inputArgs.UserId)
			s.Assert().Equal(got.WalletType, inputArgs.WalletType)
			s.Assert().Equal(got.Currency, inputArgs.Currency)
			s.Assert().Equal(got.Balance, 10.0)
			s.Assert().Equal(got.CreatedBy, "Jeff")
			s.Assert().Equal(got.UpdatedBy, "Jeff")
		},
	}

	s.Run(testcase.name, func() {
		// mock
		testcase.mockFunc()

		// execute
		got, err := s.usecase.FindWallet(testcase.args.input)

		// assert
		testcase.assertFunc(got, err)
	})

}
