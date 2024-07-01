package repo_impl

import (
	"context"
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

func (s *WalletRepoImplTestSuite) TestWalletRepoImpl_Find() {
	s.T().Skip()

	type test struct {
		name       string
		assertFunc func(wallets []domain.Wallet, err error)
	}

	testcase := &test{
		name: "Test Find",
		assertFunc: func(wallets []domain.Wallet, err error) {
			s.Assert().Len(wallets, 2)
		},
	}

	s.Run(testcase.name, func() {
		query := &domain.Wallet{UserId: "fa791816-dd35-42e6-a475-00f87d4ac9aa", WalletType: enum.WalletType.PLATFORM}

		wallets, err := s.repoImpl.Find(query)

		testcase.assertFunc(wallets, err)
	})
}

func (s *WalletRepoImplTestSuite) TestWalletRepoImpl_GetByUserIdAndWalletType() {
	s.T().Skip()

	type test struct {
		name       string
		assertFunc func(wallet domain.Wallet, err error)
	}

	testcase := &test{
		name: "Test GetByUserIdAndWalletType",
		assertFunc: func(wallet domain.Wallet, err error) {
			s.Assert().Equal("fa791816-dd35-42e6-a475-00f87d4ac9aa", wallet.UserId)
			s.Assert().Equal("Combo1", wallet.WalletType)
		},
	}

	s.Run(testcase.name, func() {
		ctx := context.Background()

		wallet, err := s.repoImpl.GetByUserIdAndWalletType(ctx, "fa791816-dd35-42e6-a475-00f87d4ac9aa", "Combo1")

		testcase.assertFunc(*wallet, err)
	})
}
