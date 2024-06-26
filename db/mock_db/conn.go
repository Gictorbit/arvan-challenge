// Code generated by MockGen. DO NOT EDIT.
// Source: conn.go

// Package mock_db is a generated GoMock package.
package mock_db

import (
	context "context"
	reflect "reflect"

	squirrel "github.com/Masterminds/squirrel"
	discountv1 "github.com/gictorbit/arvan-challenge/protos/gen/discount/v1"
	walletv1 "github.com/gictorbit/arvan-challenge/protos/gen/wallet/v1"
	gomock "github.com/golang/mock/gomock"
	pgxpool "github.com/jackc/pgx/v5/pgxpool"
)

// MockArvanDBConn is a mock of ArvanDBConn interface.
type MockArvanDBConn struct {
	ctrl     *gomock.Controller
	recorder *MockArvanDBConnMockRecorder
}

// MockArvanDBConnMockRecorder is the mock recorder for MockArvanDBConn.
type MockArvanDBConnMockRecorder struct {
	mock *MockArvanDBConn
}

// NewMockArvanDBConn creates a new mock instance.
func NewMockArvanDBConn(ctrl *gomock.Controller) *MockArvanDBConn {
	mock := &MockArvanDBConn{ctrl: ctrl}
	mock.recorder = &MockArvanDBConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArvanDBConn) EXPECT() *MockArvanDBConnMockRecorder {
	return m.recorder
}

// ApplyGiftCode mocks base method.
func (m *MockArvanDBConn) ApplyGiftCode(ctx context.Context, phone, code string) (string, float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyGiftCode", ctx, phone, code)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(float64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ApplyGiftCode indicates an expected call of ApplyGiftCode.
func (mr *MockArvanDBConnMockRecorder) ApplyGiftCode(ctx, phone, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyGiftCode", reflect.TypeOf((*MockArvanDBConn)(nil).ApplyGiftCode), ctx, phone, code)
}

// EventUsers mocks base method.
func (m *MockArvanDBConn) EventUsers(ctx context.Context, eventCode string) ([]*discountv1.UserCodeUsage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventUsers", ctx, eventCode)
	ret0, _ := ret[0].([]*discountv1.UserCodeUsage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EventUsers indicates an expected call of EventUsers.
func (mr *MockArvanDBConnMockRecorder) EventUsers(ctx, eventCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventUsers", reflect.TypeOf((*MockArvanDBConn)(nil).EventUsers), ctx, eventCode)
}

// GetPgConn mocks base method.
func (m *MockArvanDBConn) GetPgConn() *pgxpool.Pool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPgConn")
	ret0, _ := ret[0].(*pgxpool.Pool)
	return ret0
}

// GetPgConn indicates an expected call of GetPgConn.
func (mr *MockArvanDBConnMockRecorder) GetPgConn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPgConn", reflect.TypeOf((*MockArvanDBConn)(nil).GetPgConn))
}

// GetPublishedEvents mocks base method.
func (m *MockArvanDBConn) GetPublishedEvents(ctx context.Context) ([]*discountv1.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublishedEvents", ctx)
	ret0, _ := ret[0].([]*discountv1.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublishedEvents indicates an expected call of GetPublishedEvents.
func (mr *MockArvanDBConnMockRecorder) GetPublishedEvents(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublishedEvents", reflect.TypeOf((*MockArvanDBConn)(nil).GetPublishedEvents), ctx)
}

// GetSQLBuilder mocks base method.
func (m *MockArvanDBConn) GetSQLBuilder() squirrel.StatementBuilderType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSQLBuilder")
	ret0, _ := ret[0].(squirrel.StatementBuilderType)
	return ret0
}

// GetSQLBuilder indicates an expected call of GetSQLBuilder.
func (mr *MockArvanDBConnMockRecorder) GetSQLBuilder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSQLBuilder", reflect.TypeOf((*MockArvanDBConn)(nil).GetSQLBuilder))
}

// MyWallet mocks base method.
func (m *MockArvanDBConn) MyWallet(ctx context.Context, userID uint32) (*walletv1.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MyWallet", ctx, userID)
	ret0, _ := ret[0].(*walletv1.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MyWallet indicates an expected call of MyWallet.
func (mr *MockArvanDBConnMockRecorder) MyWallet(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MyWallet", reflect.TypeOf((*MockArvanDBConn)(nil).MyWallet), ctx, userID)
}

// PublishEvent mocks base method.
func (m *MockArvanDBConn) PublishEvent(ctx context.Context, eventCode string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishEvent", ctx, eventCode)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishEvent indicates an expected call of PublishEvent.
func (mr *MockArvanDBConnMockRecorder) PublishEvent(ctx, eventCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishEvent", reflect.TypeOf((*MockArvanDBConn)(nil).PublishEvent), ctx, eventCode)
}
