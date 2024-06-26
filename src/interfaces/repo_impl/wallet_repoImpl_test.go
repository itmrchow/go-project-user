package repo_impl

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"

	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/domain/enum"
	"itmrchow/go-project/user/src/infrastructure/database"
	"itmrchow/go-project/user/src/infrastructure/initialization"
)

func TestWalletRepoImplSuite(t *testing.T) {
	suite.Run(t, new(WalletRepoImplTestSuite))
}

type WalletRepoImplTestSuite struct {
	suite.Suite
	repoImpl *WalletRepoImpl
}

func (s *WalletRepoImplTestSuite) SetupTest() {
	initialization.SetConfig()

	handler, _ := database.NewMySqlHandler()
	s.repoImpl = NewWalletRepoImpl(handler)

}

func (s *WalletRepoImplTestSuite) TestWalletRepoImpl_Create() {
	s.T().Skip()

	type args struct {
		wallet *domain.Wallet
	}
	type test struct {
		name       string
		args       args
		assertFunc func(err error)
	}

	testcase := &test{
		name: "Test Create",
		args: args{
			wallet: &domain.Wallet{
				UserId:     "fa791816-dd35-42e6-a475-00f87d4ac9aa",
				WalletType: enum.WalletType.PLATFORM,
				Currency:   enum.Currency.PHP,
				DefaultModel: domain.DefaultModel{
					CreatedBy: "system",
					UpdatedBy: "system",
				},
			},
		},
		assertFunc: func(err error) {
			s.Assert().Nil(err)
		},
	}

	s.Run(testcase.name, func() {
		err := s.repoImpl.Create(testcase.args.wallet)
		testcase.assertFunc(err)
	})
}

func TestWalletRepoImpl_Get(t *testing.T) {
	type fields struct {
		handler *database.MysqlHandler
	}
	type args struct {
		walletId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Wallet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalletRepoImpl{
				handler: tt.fields.handler,
			}
			got, err := w.Get(tt.args.walletId)
			if (err != nil) != tt.wantErr {
				t.Errorf("WalletRepoImpl.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WalletRepoImpl.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
