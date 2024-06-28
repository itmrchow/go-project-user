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

type FindWalletArgs struct {
	input *FindWalletInput
}

type FindWalletTest struct {
	name       string
	args       FindWalletArgs
	mockFunc   func(FindWalletArgs)
	assertFunc func(*[]FindWalletOutput, error, FindWalletArgs)
}

func (s *WalletTestSuite) Test_FindWallet() {

	tests := []FindWalletTest{
		{
			name: "query_error",
			args: FindWalletArgs{
				input: &FindWalletInput{},
			},
			mockFunc: func(args FindWalletArgs) {
				s.repoMock.On(
					"Find",        //mock func name
					mock.Anything, //mock args
				).Return(nil, errors.New("some error")).Once()
			},
			assertFunc: func(got *[]FindWalletOutput, err error, args FindWalletArgs) {
				s.Assert().Nil(got)
				s.Assert().ErrorIs(err, ErrDbFail)
			},
		},
		{
			name: "no_data",
			args: FindWalletArgs{
				input: &FindWalletInput{
					UserId:     "Jeff",
					WalletType: "P",
					Currency:   "PHP",
				},
			},
			mockFunc: func(args FindWalletArgs) {
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
						isValueEq := s.Assert().Equal(args.input.UserId, query.UserId) &&
							s.Assert().Equal(args.input.Currency, query.Currency) &&
							s.Assert().Equal(args.input.WalletType, query.WalletType)
						return isValueEq
					}),
				).Return([]domain.Wallet{}, nil).Once()
			},
			assertFunc: func(got *[]FindWalletOutput, err error, args FindWalletArgs) {
				s.Assert().Empty(got)
				s.Assert().Nil(err)
			},
		},
		{
			name: "success",
			args: FindWalletArgs{
				input: &FindWalletInput{
					UserId:     "Jeff",
					WalletType: "P",
					Currency:   "PHP",
				},
			},
			mockFunc: func(inputArgs FindWalletArgs) {
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
						isValueEq := s.Assert().Equal(inputArgs.input.UserId, query.UserId) &&
							s.Assert().Equal(inputArgs.input.Currency, query.Currency) &&
							s.Assert().Equal(inputArgs.input.WalletType, query.WalletType)
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
				}, nil).Once()
			},
			assertFunc: func(gotSlice *[]FindWalletOutput, err error, args FindWalletArgs) {
				s.repoMock.AssertExpectations(s.T())

				s.Assert().Len(*gotSlice, 1)
				s.Assert().Nil(err)

				got := (*gotSlice)[0]
				s.Assert().Equal(got.UserId, args.input.UserId)
				s.Assert().Equal(got.WalletType, args.input.WalletType)
				s.Assert().Equal(got.Currency, args.input.Currency)
				s.Assert().Equal(got.Balance, 10.0)
				s.Assert().Equal(got.CreatedBy, args.input.UserId)
				s.Assert().Equal(got.UpdatedBy, args.input.UserId)
			},
		},
	}

	for _, test := range tests {

		s.Run(test.name, func() {
			// mock
			test.mockFunc(test.args)

			// execute
			got, err := s.usecase.FindWallet(test.args.input)

			// assert
			test.assertFunc(got, err, test.args)

		})
	}
}
