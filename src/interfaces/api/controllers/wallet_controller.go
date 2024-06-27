package controllers

import (
	"github.com/jinzhu/copier"

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

func (c *WalletController) CreateWallet(createWalletReq *reqdto.CreateWalletReq, authWallet *reqdto.AuthUser) (*respdto.CreateWalletResp, error) {

	input := usecase.CreateWalletInput{}

	if err := copier.Copy(&input, &createWalletReq); err != nil {
		return nil, err
	}

	out, err := c.walletUC.CreateWallet(&input, *authWallet)

	if err != nil {
		return nil, err
	}

	resp := respdto.CreateWalletResp{}

	if err := copier.Copy(&resp, &out); err != nil {
		return nil, err
	}

	return &resp, err
}
