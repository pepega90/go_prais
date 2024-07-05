// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/submission_service.go
//
// Generated by this command:
//
//	mockgen -source=./service/submission_service.go -destination=./test/mock/services/submission_service_mock.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	entity "assignment_1/entity"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockISubmissionService is a mock of ISubmissionService interface.
type MockISubmissionService struct {
	ctrl     *gomock.Controller
	recorder *MockISubmissionServiceMockRecorder
}

// MockISubmissionServiceMockRecorder is the mock recorder for MockISubmissionService.
type MockISubmissionServiceMockRecorder struct {
	mock *MockISubmissionService
}

// NewMockISubmissionService creates a new mock instance.
func NewMockISubmissionService(ctrl *gomock.Controller) *MockISubmissionService {
	mock := &MockISubmissionService{ctrl: ctrl}
	mock.recorder = &MockISubmissionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISubmissionService) EXPECT() *MockISubmissionServiceMockRecorder {
	return m.recorder
}

// CreateSubmission mocks base method.
func (m *MockISubmissionService) CreateSubmission(ctx context.Context, submission *entity.Submission) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubmission", ctx, submission)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSubmission indicates an expected call of CreateSubmission.
func (mr *MockISubmissionServiceMockRecorder) CreateSubmission(ctx, submission any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubmission", reflect.TypeOf((*MockISubmissionService)(nil).CreateSubmission), ctx, submission)
}

// DeleteSubmission mocks base method.
func (m *MockISubmissionService) DeleteSubmission(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubmission", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubmission indicates an expected call of DeleteSubmission.
func (mr *MockISubmissionServiceMockRecorder) DeleteSubmission(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubmission", reflect.TypeOf((*MockISubmissionService)(nil).DeleteSubmission), ctx, id)
}

// GetAllSubmissions mocks base method.
func (m *MockISubmissionService) GetAllSubmissions(ctx context.Context, limit, offset int) ([]entity.Submission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubmissions", ctx, limit, offset)
	ret0, _ := ret[0].([]entity.Submission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSubmissions indicates an expected call of GetAllSubmissions.
func (mr *MockISubmissionServiceMockRecorder) GetAllSubmissions(ctx, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubmissions", reflect.TypeOf((*MockISubmissionService)(nil).GetAllSubmissions), ctx, limit, offset)
}

// GetSubmissionByID mocks base method.
func (m *MockISubmissionService) GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubmissionByID", ctx, id)
	ret0, _ := ret[0].(entity.Submission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubmissionByID indicates an expected call of GetSubmissionByID.
func (mr *MockISubmissionServiceMockRecorder) GetSubmissionByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubmissionByID", reflect.TypeOf((*MockISubmissionService)(nil).GetSubmissionByID), ctx, id)
}

// GetSubmissionByUserID mocks base method.
func (m *MockISubmissionService) GetSubmissionByUserID(ctx context.Context, id int) (entity.Submission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubmissionByUserID", ctx, id)
	ret0, _ := ret[0].(entity.Submission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubmissionByUserID indicates an expected call of GetSubmissionByUserID.
func (mr *MockISubmissionServiceMockRecorder) GetSubmissionByUserID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubmissionByUserID", reflect.TypeOf((*MockISubmissionService)(nil).GetSubmissionByUserID), ctx, id)
}

// GetTotalSubs mocks base method.
func (m *MockISubmissionService) GetTotalSubs(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalSubs", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalSubs indicates an expected call of GetTotalSubs.
func (mr *MockISubmissionServiceMockRecorder) GetTotalSubs(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalSubs", reflect.TypeOf((*MockISubmissionService)(nil).GetTotalSubs), ctx)
}

// MockISubmissionRepository is a mock of ISubmissionRepository interface.
type MockISubmissionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockISubmissionRepositoryMockRecorder
}

// MockISubmissionRepositoryMockRecorder is the mock recorder for MockISubmissionRepository.
type MockISubmissionRepositoryMockRecorder struct {
	mock *MockISubmissionRepository
}

// NewMockISubmissionRepository creates a new mock instance.
func NewMockISubmissionRepository(ctrl *gomock.Controller) *MockISubmissionRepository {
	mock := &MockISubmissionRepository{ctrl: ctrl}
	mock.recorder = &MockISubmissionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISubmissionRepository) EXPECT() *MockISubmissionRepositoryMockRecorder {
	return m.recorder
}

// CreateSubmission mocks base method.
func (m *MockISubmissionRepository) CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubmission", ctx, submission)
	ret0, _ := ret[0].(entity.Submission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubmission indicates an expected call of CreateSubmission.
func (mr *MockISubmissionRepositoryMockRecorder) CreateSubmission(ctx, submission any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubmission", reflect.TypeOf((*MockISubmissionRepository)(nil).CreateSubmission), ctx, submission)
}

// DeleteSubmission mocks base method.
func (m *MockISubmissionRepository) DeleteSubmission(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubmission", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubmission indicates an expected call of DeleteSubmission.
func (mr *MockISubmissionRepositoryMockRecorder) DeleteSubmission(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubmission", reflect.TypeOf((*MockISubmissionRepository)(nil).DeleteSubmission), ctx, id)
}

// GetAllSubmissions mocks base method.
func (m *MockISubmissionRepository) GetAllSubmissions(ctx context.Context, limit, offset int) ([]entity.Submission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubmissions", ctx, limit, offset)
	ret0, _ := ret[0].([]entity.Submission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSubmissions indicates an expected call of GetAllSubmissions.
func (mr *MockISubmissionRepositoryMockRecorder) GetAllSubmissions(ctx, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubmissions", reflect.TypeOf((*MockISubmissionRepository)(nil).GetAllSubmissions), ctx, limit, offset)
}

// GetSubmissionByID mocks base method.
func (m *MockISubmissionRepository) GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubmissionByID", ctx, id)
	ret0, _ := ret[0].(entity.Submission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubmissionByID indicates an expected call of GetSubmissionByID.
func (mr *MockISubmissionRepositoryMockRecorder) GetSubmissionByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubmissionByID", reflect.TypeOf((*MockISubmissionRepository)(nil).GetSubmissionByID), ctx, id)
}

// GetSubmissionByUserID mocks base method.
func (m *MockISubmissionRepository) GetSubmissionByUserID(ctx context.Context, id int) (entity.Submission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubmissionByUserID", ctx, id)
	ret0, _ := ret[0].(entity.Submission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubmissionByUserID indicates an expected call of GetSubmissionByUserID.
func (mr *MockISubmissionRepositoryMockRecorder) GetSubmissionByUserID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubmissionByUserID", reflect.TypeOf((*MockISubmissionRepository)(nil).GetSubmissionByUserID), ctx, id)
}

// GetTotalSubs mocks base method.
func (m *MockISubmissionRepository) GetTotalSubs(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalSubs", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalSubs indicates an expected call of GetTotalSubs.
func (mr *MockISubmissionRepositoryMockRecorder) GetTotalSubs(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalSubs", reflect.TypeOf((*MockISubmissionRepository)(nil).GetTotalSubs), ctx)
}
