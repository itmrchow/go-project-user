package repo_impl

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/infrastructure/database"
	"itmrchow/go-project/user/src/infrastructure/initialization"
)

func TestWalletRecordRepoImplSuite(t *testing.T) {
	suite.Run(t, new(WalletRecordRepoImplTestSuite))
}

type WalletRecordRepoImplTestSuite struct {
	suite.Suite
	repoImpl *WalletRecordRepoImpl
}

func (s *WalletRecordRepoImplTestSuite) SetupTest() {
	initialization.SetConfig()

	handler, _ := database.NewMySqlHandler()
	s.repoImpl = NewWalletRecordRepoImpl(handler)
}

func (s *WalletRecordRepoImplTestSuite) TestWalletRecordRepoImpl_Create() {
	s.T().Skip()

	record := domain.WalletRecord{
		WalletId:    12,
		RecordName:  "轉入測試",
		Currency:    "PHP",
		Amount:      50.5,
		Status:      domain.WALLETRECORDSTATUS_PENDING,
		Description: "Description test",
	}

	err := s.repoImpl.Create(&record)

	s.Assert().Nil(err)

}

func (s *WalletRecordRepoImplTestSuite) TestWalletRecordRepoImpl_Get() {
	s.T().Skip()

	record, _ := s.repoImpl.Get(1)

	s.Assert().Equal(uint(1), record.ID)
	s.Assert().Equal(uint(12), record.WalletId)
	s.Assert().Equal(uint(12), record.Wallet.ID)

}
