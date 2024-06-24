package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/usecase/handler"
	"itmrchow/go-project/user/src/usecase/repo"
)

type LoginInput struct {
	Account  string
	Email    string
	Password string
}

type GetUserUseCase struct {
	userRepo          repo.UserRepo
	encryptionHandler handler.EncryptionHandler
}

func NewGetUserUseCase(userRepo repo.UserRepo, encryptionHandler handler.EncryptionHandler) *GetUserUseCase {
	return &GetUserUseCase{
		userRepo:          userRepo,
		encryptionHandler: encryptionHandler,
	}
}

// 定義output
type GetUserOutput struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Account  string `json:"account"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (c GetUserUseCase) GetUser(userId string) (*GetUserOutput, error) {
	user, err := c.userRepo.Get(userId)

	if err == nil {
		return &GetUserOutput{
			Id:       user.Id,
			UserName: user.UserName,
			Account:  user.Account,
			Email:    user.Email,
			Phone:    user.Phone,
		}, nil

	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else {
		return nil, errors.Join(ErrDbFail, err)
	}
}

// 定義output
type LoginOutput struct {
	Token string `json:"token"`
	Exp   int64  `json:"exp"`
}

func (c GetUserUseCase) Login(loginInput LoginInput) (LoginOutput, error) {
	// query user
	user, err := c.userRepo.GetByAccountOrEmail(loginInput.Account, loginInput.Email)
	if err != nil {
		return LoginOutput{}, errors.Join(ErrUserNotExists, err)
	}

	// check password
	psw := user.Password
	isCorrectPsw := c.encryptionHandler.CheckPasswordHash(loginInput.Password, psw)

	if !isCorrectPsw {
		return LoginOutput{}, ErrUnauthorized
	}

	// create token & return
	key := []byte(viper.GetString("privatekey"))

	exp := time.Now().Add(time.Hour).Unix()

	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp":      exp,
			"id":       user.Id,
			"userName": user.UserName,
			"account":  user.Account,
			"email":    user.Email,
			"phone":    user.Phone,
		})
	token, err := t.SignedString(key)

	return LoginOutput{
		Token: token,
		Exp:   exp,
	}, err
}
