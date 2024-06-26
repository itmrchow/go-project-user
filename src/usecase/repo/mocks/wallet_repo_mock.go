// Code generated by mockery v2.43.2. DO NOT EDIT.

package repo

import (
	mock "github.com/stretchr/testify/mock"

	domain "itmrchow/go-project/user/src/domain"
)

// MockWalletRepo is an autogenerated mock type for the WalletRepo type
type MockWalletRepo struct {
	mock.Mock
}

type MockWalletRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWalletRepo) EXPECT() *MockWalletRepo_Expecter {
	return &MockWalletRepo_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: wallet
func (_m *MockWalletRepo) Create(wallet *domain.Wallet) error {
	ret := _m.Called(wallet)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Wallet) error); ok {
		r0 = rf(wallet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWalletRepo_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockWalletRepo_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - wallet *domain.Wallet
func (_e *MockWalletRepo_Expecter) Create(wallet interface{}) *MockWalletRepo_Create_Call {
	return &MockWalletRepo_Create_Call{Call: _e.mock.On("Create", wallet)}
}

func (_c *MockWalletRepo_Create_Call) Run(run func(wallet *domain.Wallet)) *MockWalletRepo_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.Wallet))
	})
	return _c
}

func (_c *MockWalletRepo_Create_Call) Return(_a0 error) *MockWalletRepo_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWalletRepo_Create_Call) RunAndReturn(run func(*domain.Wallet) error) *MockWalletRepo_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Find provides a mock function with given fields: query, args
func (_m *MockWalletRepo) Find(query interface{}, args ...interface{}) ([]domain.Wallet, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 []domain.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) ([]domain.Wallet, error)); ok {
		return rf(query, args...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) []domain.Wallet); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Wallet)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, ...interface{}) error); ok {
		r1 = rf(query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWalletRepo_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type MockWalletRepo_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - query interface{}
//   - args ...interface{}
func (_e *MockWalletRepo_Expecter) Find(query interface{}, args ...interface{}) *MockWalletRepo_Find_Call {
	return &MockWalletRepo_Find_Call{Call: _e.mock.On("Find",
		append([]interface{}{query}, args...)...)}
}

func (_c *MockWalletRepo_Find_Call) Run(run func(query interface{}, args ...interface{})) *MockWalletRepo_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *MockWalletRepo_Find_Call) Return(_a0 []domain.Wallet, _a1 error) *MockWalletRepo_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWalletRepo_Find_Call) RunAndReturn(run func(interface{}, ...interface{}) ([]domain.Wallet, error)) *MockWalletRepo_Find_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: walletId
func (_m *MockWalletRepo) Get(walletId string) (*domain.Wallet, error) {
	ret := _m.Called(walletId)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *domain.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Wallet, error)); ok {
		return rf(walletId)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Wallet); ok {
		r0 = rf(walletId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Wallet)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(walletId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWalletRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockWalletRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - walletId string
func (_e *MockWalletRepo_Expecter) Get(walletId interface{}) *MockWalletRepo_Get_Call {
	return &MockWalletRepo_Get_Call{Call: _e.mock.On("Get", walletId)}
}

func (_c *MockWalletRepo_Get_Call) Run(run func(walletId string)) *MockWalletRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockWalletRepo_Get_Call) Return(_a0 *domain.Wallet, _a1 error) *MockWalletRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWalletRepo_Get_Call) RunAndReturn(run func(string) (*domain.Wallet, error)) *MockWalletRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWalletRepo creates a new instance of MockWalletRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWalletRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWalletRepo {
	mock := &MockWalletRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
