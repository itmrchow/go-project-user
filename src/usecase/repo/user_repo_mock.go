package repo

import (
	"github.com/stretchr/testify/mock"

	"itmrchow/go-project/user/src/domain"
)

type UserRepoMock struct {
	mock.Mock
}

func NewUserRepoMock() *UserRepoMock {
	return &UserRepoMock{}
}

func (m *UserRepoMock) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepoMock) Get(userId string) (*domain.User, error) {
	args := m.Called(userId)

	user, err := args.Get(0), args.Error(1)

	if user == nil {
		return nil, err
	}

	return user.(*domain.User), err

}
func (m *UserRepoMock) ExistsByAccountOrEmailOrPhone(account string, email string, phone string) (bool, error) {
	args := m.Called(account, email, phone)
	return args.Bool(0), args.Error(1)
}

func (m *UserRepoMock) GetByAccountOrEmail(account string, email string) (*domain.User, error) {
	args := m.Called(account, email)

	user, err := args.Get(0), args.Error(1)

	if user == nil {
		return nil, err
	}

	return user.(*domain.User), err
}
