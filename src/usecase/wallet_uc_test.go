package usecase

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/domain/enum"
	"itmrchow/go-project/user/src/infrastructure/api/reqdto"
	repo "itmrchow/go-project/user/src/usecase/repo/mocks"
)

func TestWalletSuite(t *testing.T) {
	suite.Run(t, new(WalletTestSuite))
}

func TestWalletOperationTestSuite(t *testing.T) {
	suite.Run(t, new(WalletOperationTestSuite))
}

type WalletOperationTestSuite struct {
	suite.Suite
	usecase              *WalletUseCase
	repoWalletMock       *repo.MockWalletRepo
	repoWalletRecordMock *repo.MockWalletRecordRepo
	authUser             reqdto.AuthUser
	dbMock               sqlmock.Sqlmock
	ctx                  *gin.Context
}

type WalletTestSuite struct {
	operationMock *operationMock
	WalletOperationTestSuite
}

type operationMock struct {
	mock.Mock
}

func (m *operationMock) CreateWalletRecord(ctx *gin.Context, walletId uint, amount float64, eventName string, depiction string, updateUserId string) (*domain.WalletRecord, error) {
	args := m.Called(ctx, walletId, amount, eventName, depiction, updateUserId)
	record, err := args.Get(0), args.Error(1)

	if record == nil {
		return nil, err
	}

	return record.(*domain.WalletRecord), err
}

func (o *operationMock) UpdateWalletByRecord(ctx *gin.Context, tx *gorm.DB, record *domain.WalletRecord) error {
	return o.Called(ctx, tx, record).Error(0)
}

func (o *operationMock) UpdateWalletRecord(ctx *gin.Context, tx *gorm.DB, record *domain.WalletRecord) error {
	return o.Called(ctx, tx, record).Error(0)
}

// func (u *mockWalletUc) IncrementMoney(ctx *gin.Context, walletId uint, amount float64, eventName string, depiction string, updateUserId string) error {
// 	return u.WalletUseCase.IncrementMoney(ctx, walletId, amount, eventName, depiction, updateUserId)
// }

func (s *WalletTestSuite) SetupTest() {
	s.repoWalletMock = &repo.MockWalletRepo{}
	s.repoWalletRecordMock = &repo.MockWalletRecordRepo{}
	s.operationMock = &operationMock{}
	s.usecase = NewWalletUseCase(s.repoWalletMock, s.repoWalletRecordMock)
	s.usecase.operation = s.operationMock

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	s.ctx = ctx

	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	gormDB, _ := gorm.Open(
		mysql.New(mysql.Config{
			Conn:                      db,
			SkipInitializeWithVersion: true,
		}), &gorm.Config{},
	)
	s.dbMock = mock

	s.ctx.Set("DB", gormDB)
}

func (s *WalletOperationTestSuite) SetupTest() {
	s.repoWalletMock = &repo.MockWalletRepo{}
	s.repoWalletRecordMock = &repo.MockWalletRecordRepo{}
	s.usecase = NewWalletUseCase(s.repoWalletMock, s.repoWalletRecordMock)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	s.ctx = ctx
	s.ctx.Set("DB", &gorm.DB{})
}

func (s *WalletOperationTestSuite) Test_CreateWallet_WalletExist() {
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
			s.repoWalletMock.On("Create", mock.Anything).Return(gorm.ErrDuplicatedKey).Once()
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

func (s *WalletOperationTestSuite) Test_CreateWallet_Success() {
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
			s.repoWalletMock.On("Create", mock.Anything).Return(nil).Once()
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

func (s *WalletOperationTestSuite) Test_FindWallet() {

	tests := []FindWalletTest{
		{
			name: "query_error",
			args: FindWalletArgs{
				input: &FindWalletInput{},
			},
			mockFunc: func(args FindWalletArgs) {
				s.repoWalletMock.On(
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
					UserId:     "Jeff_NoData",
					WalletType: "P",
					Currency:   "PHP",
				},
			},
			mockFunc: func(args FindWalletArgs) {
				s.repoWalletMock.On(
					//mock func name
					"Find",
					//mock args
					mock.MatchedBy(func(i interface{}) bool {
						query, isTypeRight := i.(domain.Wallet)

						if !isTypeRight {
							return false
						}

						// return args.input.UserId == "Jeff_NoData"
						isValueEq := args.input.UserId == query.UserId &&
							args.input.Currency == query.Currency &&
							args.input.WalletType == query.WalletType
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
					UserId:     "Jeff_Success",
					WalletType: "P",
					Currency:   "PHP",
				},
			},
			mockFunc: func(args FindWalletArgs) {
				s.repoWalletMock.On(
					//mock func name
					"Find",
					//mock args
					mock.MatchedBy(func(i interface{}) bool {
						query, isTypeRight := i.(domain.Wallet)

						if !isTypeRight {
							return false
						}

						// 確認值
						isValueEq := args.input.UserId == query.UserId &&
							args.input.Currency == query.Currency &&
							args.input.WalletType == query.WalletType
						return isValueEq
					}),
				).Return([]domain.Wallet{
					{
						UserId:     "Jeff_Success",
						WalletType: enum.WalletType.PLATFORM,
						Currency:   enum.Currency.PHP,
						Balance:    10.0,
						DefaultModel: domain.DefaultModel{
							CreatedBy: "Jeff_Success",
							UpdatedBy: "Jeff_Success",
						},
					},
				}, nil).Once()
			},
			assertFunc: func(gotSlice *[]FindWalletOutput, err error, args FindWalletArgs) {
				s.repoWalletMock.AssertExpectations(s.T())

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

			s.repoWalletMock.AssertExpectations(s.T())
		})
	}
}

func (s *WalletOperationTestSuite) Test_UpdateWalletRecord() {
	type args struct {
		ctx    *gin.Context
		tx     *gorm.DB
		record *domain.WalletRecord
	}

	type testcase struct {
		name       string
		args       args
		mockFunc   func(args)
		assertFunc func(error)
	}

	tests := []testcase{
		{
			name: "db_fail",
			args: args{
				ctx:    s.ctx,
				tx:     nil,
				record: &domain.WalletRecord{
					// UserId:     "Jeff",
					// WalletType: enum.WalletType.PLATFORM,
					// Currency:   enum.Currency.PHP,
					// Amount:     10.0,
					// Balance:    10.0,
				},
			},
			mockFunc: func(a args) {
				s.repoWalletRecordMock.
					On("Update",
						mock.AnythingOfType("*gin.Context"),
						mock.AnythingOfType("*domain.WalletRecord")).
					Return(int64(0), gorm.ErrInvalidData).
					Once()
			},
			assertFunc: func(err error) {
				s.Assert().ErrorIs(err, ErrDbFail)
			},
		},
		{
			name: "data_not_exist",
			args: args{
				ctx:    s.ctx,
				tx:     nil,
				record: &domain.WalletRecord{},
			},
			mockFunc: func(a args) {
				s.repoWalletRecordMock.
					On("Update",
						mock.AnythingOfType("*gin.Context"),
						mock.AnythingOfType("*domain.WalletRecord")).
					Return(int64(0), nil).
					Once()
			},
			assertFunc: func(err error) {
				s.Assert().ErrorIs(err, ErrDataNotExists)
			},
		},
		{
			name: "use_tx_success",
			args: args{
				ctx:    s.ctx,
				tx:     &gorm.DB{},
				record: &domain.WalletRecord{},
			},
			mockFunc: func(a args) {
				s.repoWalletRecordMock.
					On("WithTrx",
						mock.AnythingOfType("*gorm.DB")).
					Return(s.repoWalletRecordMock).
					Once()

				s.repoWalletRecordMock.
					On("Update",
						mock.AnythingOfType("*gin.Context"),
						mock.AnythingOfType("*domain.WalletRecord")).
					Return(int64(1), nil).
					Once()
			},
			assertFunc: func(err error) {
				s.Assert().Nil(err)
			},
		},
		{
			name: "success",
			args: args{
				ctx:    s.ctx,
				tx:     nil,
				record: &domain.WalletRecord{},
			},
			mockFunc: func(a args) {
				s.repoWalletRecordMock.
					On("Update",
						mock.AnythingOfType("*gin.Context"),
						mock.AnythingOfType("*domain.WalletRecord")).
					Return(int64(1), nil).
					Once()
			},
			assertFunc: func(err error) {
				s.Assert().Nil(err)
			},
		},
	}

	for _, test := range tests {

		s.Run(test.name, func() {
			// mock
			test.mockFunc(test.args)

			// execute
			err := s.usecase.UpdateWalletRecord(test.args.ctx, test.args.tx, test.args.record)

			// assert
			test.assertFunc(err)

			s.repoWalletRecordMock.AssertExpectations(s.T())
		})
	}
}

type MockUc_UpdateWalletByRecord struct {
	mock.Mock
	WalletUseCase
}

func (m *MockUc_UpdateWalletByRecord) UpdateWalletRecord(ctx *gin.Context, tx *gorm.DB, record *domain.WalletRecord) error {
	args := m.Called(ctx, tx, record)
	return args.Error(0)
}

func (s *WalletOperationTestSuite) Test_UpdateWalletByRecord() {
	type args struct {
		ctx    *gin.Context
		tx     *gorm.DB
		record *domain.WalletRecord
	}
	type testcase struct {
		name       string
		args       args
		mockFunc   func(args)
		assertFunc func(error, args)
	}

	tests := []testcase{
		{
			name: "get_wallet_not_exist",
			args: args{
				ctx:    s.ctx,
				record: &domain.WalletRecord{
					// UserId:     "Jeff",
					// WalletType: enum.WalletType.PLATFORM,
					// Currency:   enum.Currency.PHP,
					// Amount:     10.0,
					// Balance:    10.0,
				},
			},
			mockFunc: func(a args) {
				s.repoWalletMock.
					On("WithTrx",
						mock.AnythingOfType("*gorm.DB")).
					Return(s.repoWalletMock).
					Once()

				s.repoWalletMock.
					On("GetWalletWithLock",
						mock.AnythingOfType("*gin.Context"),
						mock.AnythingOfType("uint"),
					).
					Return(nil, gorm.ErrRecordNotFound).
					Once()

			},
			assertFunc: func(err error, args args) {
				s.Assert().Nil(err)
				s.Assert().Equal(args.record.Status, domain.WALLET_RECORD_STATUS_FAILED)
			},
		},
		{
			name: "get_wallet_db_fail",
			args: args{
				ctx:    s.ctx,
				record: &domain.WalletRecord{
					// UserId:     "Jeff",
					// WalletType: enum.WalletType.PLATFORM,
					// Currency:   enum.Currency.PHP,
					// Amount:     10.0,
					// Balance:    10.0,
				},
			},
			mockFunc: func(a args) {
				s.repoWalletMock.
					On("WithTrx",
						mock.AnythingOfType("*gorm.DB")).
					Return(s.repoWalletMock).
					Once()

				s.repoWalletMock.
					On("GetWalletWithLock",
						mock.AnythingOfType("*gin.Context"),
						mock.AnythingOfType("uint"),
					).
					Return(nil, gorm.ErrInvalidDB).
					Once()

			},
			assertFunc: func(err error, args args) {
				s.Assert().ErrorIs(err, ErrDbFail)
			},
		},
		{
			name: "check_decrement_amount_fail",
			args: args{
				ctx: s.ctx,
				record: &domain.WalletRecord{
					Amount: -10.0,
				},
			},
			mockFunc: func(a args) {
				s.repoWalletMock.
					On("WithTrx",
						mock.AnythingOfType("*gorm.DB")).
					Return(s.repoWalletMock).
					Once()

				s.repoWalletMock.
					On("GetWalletWithLock",
						mock.AnythingOfType("*gin.Context"),
						mock.AnythingOfType("uint"),
					).
					Return(&domain.Wallet{
						Balance: 3.5,
					}, nil).
					Once()

			},
			assertFunc: func(err error, args args) {
				s.Assert().Nil(err)
				s.Assert().Equal(args.record.Status, domain.WALLET_RECORD_STATUS_FAILED)
			},
		},
		{
			name: "update_db_fail",
			args: args{
				ctx: s.ctx,
				record: &domain.WalletRecord{
					Amount: -2.0,
				},
			},
			mockFunc: func(a args) {
				s.repoWalletMock.
					On("WithTrx",
						mock.AnythingOfType("*gorm.DB")).
					Return(s.repoWalletMock).
					Once()

				s.repoWalletMock.
					On("GetWalletWithLock",
						mock.AnythingOfType("*gin.Context"),
						mock.AnythingOfType("uint"),
					).
					Return(&domain.Wallet{
						Balance: 3.5,
					}, nil).
					Once()

				s.repoWalletMock.
					On("Update",
						mock.AnythingOfType("*domain.Wallet"),
						mock.AnythingOfType("float64"),
					).
					Return(int64(0), gorm.ErrInvalidDB).
					Once()

			},
			assertFunc: func(err error, args args) {
				s.Assert().ErrorIs(err, ErrDbFail)
			},
		},
		{
			name: "update_wallet_count_zero",
			args: args{
				ctx: s.ctx,
				record: &domain.WalletRecord{
					WalletId: 1,
					Amount:   -2.0,
				},
			},
			mockFunc: func(a args) {
				s.repoWalletMock.
					On("WithTrx", mock.AnythingOfType("*gorm.DB")).
					Return(s.repoWalletMock).
					Once()

				s.repoWalletMock.
					On("GetWalletWithLock", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("uint")).
					Return(&domain.Wallet{
						Balance: 3.5,
					}, nil).
					Once()

				s.repoWalletMock.
					On("Update", mock.AnythingOfType("*domain.Wallet"), mock.AnythingOfType("float64")).
					Return(int64(0), nil).
					Once()
			},
			assertFunc: func(err error, args args) {
				s.Assert().NoError(err)
				s.Assert().Equal(domain.WALLET_RECORD_STATUS_FAILED, args.record.Status)
			},
		},
		{
			name: "update_wallet_success",
			args: args{
				ctx: s.ctx,
				record: &domain.WalletRecord{
					WalletId: 1,
					Amount:   -2.0,
				},
			},
			mockFunc: func(a args) {
				s.repoWalletMock.
					On("WithTrx", mock.AnythingOfType("*gorm.DB")).
					Return(s.repoWalletMock).
					Once()

				s.repoWalletMock.
					On("GetWalletWithLock", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("uint")).
					Return(&domain.Wallet{
						Balance: 3.5,
					}, nil).
					Once()

				s.repoWalletMock.
					On("Update", mock.AnythingOfType("*domain.Wallet"), mock.AnythingOfType("float64")).
					Return(int64(1), nil).
					Once()
			},
			assertFunc: func(err error, args args) {
				s.Assert().NoError(err)
				s.Assert().Equal(domain.WALLET_RECORD_STATUS_SUCCESS, args.record.Status)
			},
		},
	}

	for _, test := range tests {

		s.Run(test.name, func() {
			// mock
			test.mockFunc(test.args)

			// execute
			err := s.usecase.UpdateWalletByRecord(test.args.ctx, test.args.tx, test.args.record)

			// assert
			test.assertFunc(err, test.args)

			s.repoWalletMock.AssertExpectations(s.T())
			s.repoWalletRecordMock.AssertExpectations(s.T())
		})
	}
}

func (s *WalletOperationTestSuite) Test_CreateWalletRecord() {
	type args struct {
		ctx          *gin.Context
		walletId     uint
		amount       float64
		eventName    string
		depiction    string
		updateUserId string
	}
	type testcase struct {
		name       string
		args       args
		mockFunc   func(args)
		assertFunc func(*domain.WalletRecord, error, args)
	}

	tests := []testcase{
		{
			name: "get_wallet_not_exist",
			args: args{
				walletId: 13,
			},
			mockFunc: func(a args) {
				s.repoWalletMock.
					On("Get",
						mock.AnythingOfType("*gin.Context"),
						a.walletId,
					).Return(nil, gorm.ErrRecordNotFound).Once()

			},
			assertFunc: func(record *domain.WalletRecord, err error, args args) {
				s.Assert().Nil(record)
				s.Assert().ErrorIs(err, ErrDataNotExists)
			},
		},
		{
			name: "get_wallet_db_fail",
			args: args{
				walletId: 13,
			},
			mockFunc: func(a args) {
				s.repoWalletMock.
					On("Get",
						mock.AnythingOfType("*gin.Context"),
						a.walletId,
					).Return(nil, gorm.ErrInvalidDB).Once()

			},
			assertFunc: func(record *domain.WalletRecord, err error, args args) {
				s.Assert().Nil(record)
				s.Assert().ErrorIs(err, ErrDbFail)
			},
		},
		{
			name: "create_record_fail",
			args: args{
				walletId:     13,
				amount:       54,
				eventName:    "轉入",
				depiction:    "AAA轉入",
				updateUserId: "AAA",
			},
			mockFunc: func(a args) {
				s.repoWalletMock.
					On("Get",
						mock.AnythingOfType("*gin.Context"),
						a.walletId,
					).Return(
					&domain.Wallet{
						Currency: "PHP",
						DefaultModel: domain.DefaultModel{
							Model: gorm.Model{
								ID: a.walletId,
							},
						},
					}, nil).Once()

				s.repoWalletRecordMock.
					On("Create",
						mock.AnythingOfType("*gin.Context"),
						mock.MatchedBy(func(record *domain.WalletRecord) bool {
							return record.WalletId == a.walletId &&
								record.Currency == "PHP" &&
								record.RecordName == a.eventName &&
								record.Status == domain.WALLET_RECORD_STATUS_PENDING &&
								record.Description == a.depiction
						}),
					).Return(gorm.ErrInvalidDB).Once()
			},
			assertFunc: func(record *domain.WalletRecord, err error, args args) {
				s.Assert().Nil(record)
				s.Assert().ErrorIs(err, ErrDbFail)
			},
		},
		{
			name: "create_record_success",
			args: args{
				walletId:     13,
				amount:       54,
				eventName:    "轉入",
				depiction:    "AAA轉入",
				updateUserId: "AAA",
			},
			mockFunc: func(a args) {
				s.repoWalletMock.
					On("Get",
						mock.AnythingOfType("*gin.Context"),
						a.walletId,
					).Return(
					&domain.Wallet{
						Currency: "PHP",
						DefaultModel: domain.DefaultModel{
							Model: gorm.Model{
								ID: a.walletId,
							},
						},
					}, nil).Once()

				s.repoWalletRecordMock.
					On("Create",
						mock.AnythingOfType("*gin.Context"),
						mock.MatchedBy(func(record *domain.WalletRecord) bool {
							return record.WalletId == a.walletId &&
								record.Currency == "PHP" &&
								record.RecordName == a.eventName &&
								record.Status == domain.WALLET_RECORD_STATUS_PENDING &&
								record.Description == a.depiction
						}),
					).Return(nil).Once()
			},
			assertFunc: func(record *domain.WalletRecord, err error, a args) {
				s.Assert().Nil(err)
				s.Assert().Equal(record.WalletId, a.walletId)
				s.Assert().Equal(record.Currency, "PHP")
				s.Assert().Equal(record.RecordName, a.eventName)
				s.Assert().Equal(record.Status, domain.WALLET_RECORD_STATUS_PENDING)
				s.Assert().Equal(record.Description, a.depiction)
			},
		},
	}

	for _, test := range tests {

		s.Run(test.name, func() {
			// mock
			test.mockFunc(test.args)

			// execute
			record, err := s.usecase.CreateWalletRecord(test.args.ctx, test.args.walletId, test.args.amount, test.args.eventName, test.args.depiction, test.args.updateUserId)

			// assert
			test.assertFunc(record, err, test.args)

			s.repoWalletMock.AssertExpectations(s.T())
			s.repoWalletRecordMock.AssertExpectations(s.T())
		})
	}
}

func (s *WalletTestSuite) Test_IncrementMoney() {
	type args struct {
		ctx          *gin.Context
		walletId     uint
		amount       float64
		eventName    string
		depiction    string
		updateUserId string
	}
	type testcase struct {
		name       string
		args       args
		mockFunc   func(args)
		assertFunc func(error, args)
	}

	tests := []testcase{
		{
			name: "create_wallet_record_fail",
			args: args{
				ctx:          s.ctx,
				walletId:     12,
				amount:       12,
				eventName:    "test",
				depiction:    "test",
				updateUserId: "test",
			},
			mockFunc: func(a args) {
				s.operationMock.On("CreateWalletRecord", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, ErrDbFail).
					Once()

			},
			assertFunc: func(err error, args args) {
				s.Assert().ErrorIs(err, ErrDbFail)
			},
		},
		{
			name: "update_wallet_by_record_fail",
			args: args{
				ctx:          s.ctx,
				walletId:     12,
				amount:       12,
				eventName:    "test",
				depiction:    "test",
				updateUserId: "test",
			},
			mockFunc: func(a args) {
				s.operationMock.On("CreateWalletRecord", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(&domain.WalletRecord{}, nil).
					Once()

				s.dbMock.ExpectBegin()

				s.operationMock.On("UpdateWalletByRecord", mock.Anything, mock.Anything, mock.Anything).
					Return(ErrDbFail).
					Once()
			},
			assertFunc: func(err error, args args) {
				s.Assert().ErrorIs(err, ErrDbFail)
			},
		},
		{
			name: "UpdateWalletRecord_fail",
			args: args{
				ctx:          s.ctx,
				walletId:     12,
				amount:       12,
				eventName:    "test",
				depiction:    "test",
				updateUserId: "test",
			},
			mockFunc: func(a args) {
				s.operationMock.On("CreateWalletRecord", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(&domain.WalletRecord{}, nil).
					Once()

				s.dbMock.ExpectBegin()

				s.operationMock.On("UpdateWalletByRecord", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Once()

				s.operationMock.On("UpdateWalletRecord", mock.Anything, mock.Anything, mock.Anything).
					Return(ErrDbFail).
					Once()
			},
			assertFunc: func(err error, args args) {
				s.Assert().ErrorIs(err, ErrDbFail)
			},
		},
	}

	for _, test := range tests {

		s.Run(test.name, func() {
			// mock
			test.mockFunc(test.args)

			// execute
			err := s.usecase.IncrementMoney(test.args.ctx, test.args.walletId, test.args.amount, test.args.eventName, test.args.depiction, test.args.updateUserId)

			// assert
			test.assertFunc(err, test.args)

			s.operationMock.AssertExpectations(s.T())
			s.repoWalletMock.AssertExpectations(s.T())
			s.repoWalletRecordMock.AssertExpectations(s.T())
		})
	}
}

func (s *WalletTestSuite) Test_DecrementMoney() {
	type args struct {
		ctx          *gin.Context
		walletId     uint
		amount       float64
		eventName    string
		depiction    string
		updateUserId string
	}
	type testcase struct {
		name       string
		args       args
		mockFunc   func(args)
		assertFunc func(error, args)
	}

	tests := []testcase{
		{
			name: "UpdateWalletRecord_fail",
			args: args{
				ctx:          s.ctx,
				walletId:     12,
				amount:       12,
				eventName:    "test",
				depiction:    "test",
				updateUserId: "test",
			},
			mockFunc: func(a args) {
				s.operationMock.On("CreateWalletRecord", mock.Anything, mock.Anything, float64(-12), mock.Anything, mock.Anything, mock.Anything).
					Return(&domain.WalletRecord{}, nil).
					Once()

				s.dbMock.ExpectBegin()

				s.operationMock.On("UpdateWalletByRecord", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Once()

				s.operationMock.On("UpdateWalletRecord", mock.Anything, mock.Anything, mock.Anything).
					Return(ErrDbFail).
					Once()
			},
			assertFunc: func(err error, args args) {
				s.Assert().ErrorIs(err, ErrDbFail)
			},
		},
	}

	for _, test := range tests {

		s.Run(test.name, func() {
			// mock
			test.mockFunc(test.args)

			// execute
			err := s.usecase.DecrementMoney(test.args.ctx, test.args.walletId, test.args.amount, test.args.eventName, test.args.depiction, test.args.updateUserId)

			// assert
			test.assertFunc(err, test.args)

			s.operationMock.AssertExpectations(s.T())
			s.repoWalletMock.AssertExpectations(s.T())
			s.repoWalletRecordMock.AssertExpectations(s.T())
		})
	}
}
