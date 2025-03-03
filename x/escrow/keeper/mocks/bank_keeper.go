// Code generated by mockery v2.24.0. DO NOT EDIT.

package mocks

import (
	types "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"
)

// BankKeeper is an autogenerated mock type for the BankKeeper type
type BankKeeper struct {
	mock.Mock
}

type BankKeeper_Expecter struct {
	mock *mock.Mock
}

func (_m *BankKeeper) EXPECT() *BankKeeper_Expecter {
	return &BankKeeper_Expecter{mock: &_m.Mock}
}

// SendCoinsFromAccountToModule provides a mock function with given fields: ctx, senderAddr, recipientModule, amt
func (_m *BankKeeper) SendCoinsFromAccountToModule(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	ret := _m.Called(ctx, senderAddr, recipientModule, amt)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, string, types.Coins) error); ok {
		r0 = rf(ctx, senderAddr, recipientModule, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BankKeeper_SendCoinsFromAccountToModule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendCoinsFromAccountToModule'
type BankKeeper_SendCoinsFromAccountToModule_Call struct {
	*mock.Call
}

// SendCoinsFromAccountToModule is a helper method to define mock.On call
//   - ctx types.Context
//   - senderAddr types.AccAddress
//   - recipientModule string
//   - amt types.Coins
func (_e *BankKeeper_Expecter) SendCoinsFromAccountToModule(ctx interface{}, senderAddr interface{}, recipientModule interface{}, amt interface{}) *BankKeeper_SendCoinsFromAccountToModule_Call {
	return &BankKeeper_SendCoinsFromAccountToModule_Call{Call: _e.mock.On("SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)}
}

func (_c *BankKeeper_SendCoinsFromAccountToModule_Call) Run(run func(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins)) *BankKeeper_SendCoinsFromAccountToModule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.AccAddress), args[2].(string), args[3].(types.Coins))
	})
	return _c
}

func (_c *BankKeeper_SendCoinsFromAccountToModule_Call) Return(_a0 error) *BankKeeper_SendCoinsFromAccountToModule_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BankKeeper_SendCoinsFromAccountToModule_Call) RunAndReturn(run func(types.Context, types.AccAddress, string, types.Coins) error) *BankKeeper_SendCoinsFromAccountToModule_Call {
	_c.Call.Return(run)
	return _c
}

// SendCoinsFromModuleToAccount provides a mock function with given fields: ctx, senderModule, recipientAddr, amt
func (_m *BankKeeper) SendCoinsFromModuleToAccount(ctx types.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	ret := _m.Called(ctx, senderModule, recipientAddr, amt)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, string, types.AccAddress, types.Coins) error); ok {
		r0 = rf(ctx, senderModule, recipientAddr, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BankKeeper_SendCoinsFromModuleToAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendCoinsFromModuleToAccount'
type BankKeeper_SendCoinsFromModuleToAccount_Call struct {
	*mock.Call
}

// SendCoinsFromModuleToAccount is a helper method to define mock.On call
//   - ctx types.Context
//   - senderModule string
//   - recipientAddr types.AccAddress
//   - amt types.Coins
func (_e *BankKeeper_Expecter) SendCoinsFromModuleToAccount(ctx interface{}, senderModule interface{}, recipientAddr interface{}, amt interface{}) *BankKeeper_SendCoinsFromModuleToAccount_Call {
	return &BankKeeper_SendCoinsFromModuleToAccount_Call{Call: _e.mock.On("SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)}
}

func (_c *BankKeeper_SendCoinsFromModuleToAccount_Call) Run(run func(ctx types.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins)) *BankKeeper_SendCoinsFromModuleToAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(string), args[2].(types.AccAddress), args[3].(types.Coins))
	})
	return _c
}

func (_c *BankKeeper_SendCoinsFromModuleToAccount_Call) Return(_a0 error) *BankKeeper_SendCoinsFromModuleToAccount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BankKeeper_SendCoinsFromModuleToAccount_Call) RunAndReturn(run func(types.Context, string, types.AccAddress, types.Coins) error) *BankKeeper_SendCoinsFromModuleToAccount_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewBankKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewBankKeeper creates a new instance of BankKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBankKeeper(t mockConstructorTestingTNewBankKeeper) *BankKeeper {
	mock := &BankKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
