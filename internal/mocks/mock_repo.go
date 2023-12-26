package mocks

import (
	reflect "reflect"

	model "github.com/akmyrzza/go-musthave-shortener/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateShortURL mocks base method.
func (m *MockRepository) CreateShortURL(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShortURL", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateShortURL indicates an expected call of CreateShortURL.
func (mr *MockRepositoryMockRecorder) CreateShortURL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShortURL", reflect.TypeOf((*MockRepository)(nil).CreateShortURL), arg0, arg1)
}

// CreateShortURLs mocks base method.
func (m *MockRepository) CreateShortURLs(arg0 []model.ReqURL) ([]model.ReqURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShortURLs", arg0)
	ret0, _ := ret[0].([]model.ReqURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateShortURLs indicates an expected call of CreateShortURLs.
func (mr *MockRepositoryMockRecorder) CreateShortURLs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShortURLs", reflect.TypeOf((*MockRepository)(nil).CreateShortURLs), arg0)
}

// GetOriginalURL mocks base method.
func (m *MockRepository) GetOriginalURL(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOriginalURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOriginalURL indicates an expected call of GetOriginalURL.
func (mr *MockRepositoryMockRecorder) GetOriginalURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOriginalURL", reflect.TypeOf((*MockRepository)(nil).GetOriginalURL), arg0)
}

// PingStore mocks base method.
func (m *MockRepository) PingStore() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingStore")
	ret0, _ := ret[0].(error)
	return ret0
}

// PingStore indicates an expected call of PingStore.
func (mr *MockRepositoryMockRecorder) PingStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingStore", reflect.TypeOf((*MockRepository)(nil).PingStore))
}
