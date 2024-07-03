package repo_impl

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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
	ctx      *gin.Context
}

func (s *WalletRecordRepoImplTestSuite) SetupTest() {
	initialization.SetConfig()

	handler, _ := database.NewMySqlHandler()
	s.repoImpl = NewWalletRecordRepoImpl(handler)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	s.ctx = ctx
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

	err := s.repoImpl.Create(s.ctx, &record)

	s.Assert().Nil(err)

}

func (s *WalletRecordRepoImplTestSuite) TestWalletRecordRepoImpl_Get() {
	s.T().Skip()

	record, _ := s.repoImpl.Get(s.ctx, 3)

	s.Assert().Equal(uint(1), record.ID)
	s.Assert().Equal(uint(12), record.WalletId)
	s.Assert().Equal(uint(12), record.Wallet.ID)

}
