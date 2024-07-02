package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/infrastructure/api/reqdto"
	"itmrchow/go-project/user/src/infrastructure/api/respdto"
	"itmrchow/go-project/user/src/usecase"
)

type WalletController struct {
	walletUC *usecase.WalletUseCase
}

func NewWalletController(walletUC *usecase.WalletUseCase) *WalletController {
	return &WalletController{walletUC: walletUC}
}

func (c *WalletController) CreateWallet(createWalletReq *reqdto.CreateWalletReq, authUser *reqdto.AuthUser) (*respdto.CreateWalletResp, error) {

	input := usecase.CreateWalletInput{}

	if err := copier.Copy(&input, &createWalletReq); err != nil {
		return nil, err
	}

	out, err := c.walletUC.CreateWallet(&input, *authUser)

	if err != nil {
		return nil, err
	}

	resp := respdto.CreateWalletResp{}

	if err := copier.Copy(&resp, &out); err != nil {
		return nil, err
	}

	return &resp, err
}

func (c *WalletController) FindWallets(req *reqdto.FindWalletsReq, authUser *reqdto.AuthUser) (*[]respdto.FindWalletResp, error) {

	input := usecase.FindWalletInput{
		UserId: authUser.Id,
	}

	if err := copier.Copy(&input, &req); err != nil {
		return nil, err
	}

	out, err := c.walletUC.FindWallet(&input)

	if err != nil {
		return nil, err
	}

	respSlice := []respdto.FindWalletResp{}
	if err := copier.Copy(&respSlice, out); err != nil {
		return nil, err
	}

	return &respSlice, nil
}

func (c *WalletController) TransferFunds(ctx *gin.Context, req *reqdto.TransferFundsReq, authUser *reqdto.AuthUser) error {
	log.Print("[WalletController]...TransferFunds")

	txHandle := ctx.MustGet("db_trx").(*gorm.DB) // 取得transaction

	uc := c.walletUC.WithTrx(txHandle)

	input := usecase.TransferFundsInput{}

	if err := copier.Copy(&input, req); err != nil {
		return err
	}

	if err := uc.TransferFunds(ctx, &input, *authUser); err != nil {
		return err
	}

	return nil
}
