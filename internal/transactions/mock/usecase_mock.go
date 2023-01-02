// Code generated by MockGen. DO NOT EDIT.
// Source: internal/transactions/usecase.go

// Package mock_transactions is a generated GoMock package.
package mock_transactions

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	models "github.com/katsuragawaa/btc-billionaire/internal/models"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUseCase) Create(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, transaction)
	ret0, _ := ret[0].(*models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUseCaseMockRecorder) Create(ctx, transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), ctx, transaction)
}

// GetBalance mocks base method.
func (m *MockUseCase) GetBalance(ctx context.Context) (*models.TransactionsBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx)
	ret0, _ := ret[0].(*models.TransactionsBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockUseCaseMockRecorder) GetBalance(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockUseCase)(nil).GetBalance), ctx)
}

// GetPerHours mocks base method.
func (m *MockUseCase) GetPerHours(ctx context.Context, start, end time.Time) (*models.TransactionsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPerHours", ctx, start, end)
	ret0, _ := ret[0].(*models.TransactionsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPerHours indicates an expected call of GetPerHours.
func (mr *MockUseCaseMockRecorder) GetPerHours(ctx, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPerHours", reflect.TypeOf((*MockUseCase)(nil).GetPerHours), ctx, start, end)
}