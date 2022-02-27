// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/drkennetz/codingchallenge/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/drkennetz/codingchallenge/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(arg0 context.Context, arg1 db.CreateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), arg0, arg1)
}

// CreateAccountUserTx mocks base method.
func (m *MockStore) CreateAccountUserTx(arg0 context.Context, arg1 db.CreateAccountUserTxParams) (db.CreateAccountUserTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccountUserTx", arg0, arg1)
	ret0, _ := ret[0].(db.CreateAccountUserTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccountUserTx indicates an expected call of CreateAccountUserTx.
func (mr *MockStoreMockRecorder) CreateAccountUserTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccountUserTx", reflect.TypeOf((*MockStore)(nil).CreateAccountUserTx), arg0, arg1)
}

// CreateCategory mocks base method.
func (m *MockStore) CreateCategory(arg0 context.Context, arg1 string) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockStoreMockRecorder) CreateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockStore)(nil).CreateCategory), arg0, arg1)
}

// CreateQuestion mocks base method.
func (m *MockStore) CreateQuestion(arg0 context.Context, arg1 db.CreateQuestionParams) (db.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateQuestion", arg0, arg1)
	ret0, _ := ret[0].(db.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQuestion indicates an expected call of CreateQuestion.
func (mr *MockStoreMockRecorder) CreateQuestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQuestion", reflect.TypeOf((*MockStore)(nil).CreateQuestion), arg0, arg1)
}

// CreateQuestionCategory mocks base method.
func (m *MockStore) CreateQuestionCategory(arg0 context.Context, arg1 db.CreateQuestionCategoryParams) (db.QuestionCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateQuestionCategory", arg0, arg1)
	ret0, _ := ret[0].(db.QuestionCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQuestionCategory indicates an expected call of CreateQuestionCategory.
func (mr *MockStoreMockRecorder) CreateQuestionCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQuestionCategory", reflect.TypeOf((*MockStore)(nil).CreateQuestionCategory), arg0, arg1)
}

// CreateQuestionCategoryTx mocks base method.
func (m *MockStore) CreateQuestionCategoryTx(arg0 context.Context, arg1 db.CreateQuestionCategoryTxParams) (db.CreateQuestionCategoryTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateQuestionCategoryTx", arg0, arg1)
	ret0, _ := ret[0].(db.CreateQuestionCategoryTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQuestionCategoryTx indicates an expected call of CreateQuestionCategoryTx.
func (mr *MockStoreMockRecorder) CreateQuestionCategoryTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQuestionCategoryTx", reflect.TypeOf((*MockStore)(nil).CreateQuestionCategoryTx), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateUserQuestionScore mocks base method.
func (m *MockStore) CreateUserQuestionScore(arg0 context.Context, arg1 db.CreateUserQuestionScoreParams) (db.UserQuestionScore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserQuestionScore", arg0, arg1)
	ret0, _ := ret[0].(db.UserQuestionScore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserQuestionScore indicates an expected call of CreateUserQuestionScore.
func (mr *MockStoreMockRecorder) CreateUserQuestionScore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserQuestionScore", reflect.TypeOf((*MockStore)(nil).CreateUserQuestionScore), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), arg0, arg1)
}

// DeleteAccountUserTx mocks base method.
func (m *MockStore) DeleteAccountUserTx(arg0 context.Context, arg1 db.DeleteAccountUserTxParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccountUserTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccountUserTx indicates an expected call of DeleteAccountUserTx.
func (mr *MockStoreMockRecorder) DeleteAccountUserTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccountUserTx", reflect.TypeOf((*MockStore)(nil).DeleteAccountUserTx), arg0, arg1)
}

// DeleteCategory mocks base method.
func (m *MockStore) DeleteCategory(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockStoreMockRecorder) DeleteCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockStore)(nil).DeleteCategory), arg0, arg1)
}

// DeleteQuestion mocks base method.
func (m *MockStore) DeleteQuestion(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteQuestion", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteQuestion indicates an expected call of DeleteQuestion.
func (mr *MockStoreMockRecorder) DeleteQuestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteQuestion", reflect.TypeOf((*MockStore)(nil).DeleteQuestion), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// DeleteUserQuestionScore mocks base method.
func (m *MockStore) DeleteUserQuestionScore(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserQuestionScore", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserQuestionScore indicates an expected call of DeleteUserQuestionScore.
func (mr *MockStoreMockRecorder) DeleteUserQuestionScore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserQuestionScore", reflect.TypeOf((*MockStore)(nil).DeleteUserQuestionScore), arg0, arg1)
}

// GetACategoryByID mocks base method.
func (m *MockStore) GetACategoryByID(arg0 context.Context, arg1 int64) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetACategoryByID", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetACategoryByID indicates an expected call of GetACategoryByID.
func (mr *MockStoreMockRecorder) GetACategoryByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetACategoryByID", reflect.TypeOf((*MockStore)(nil).GetACategoryByID), arg0, arg1)
}

// GetACategoryIDByName mocks base method.
func (m *MockStore) GetACategoryIDByName(arg0 context.Context, arg1 string) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetACategoryIDByName", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetACategoryIDByName indicates an expected call of GetACategoryIDByName.
func (mr *MockStoreMockRecorder) GetACategoryIDByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetACategoryIDByName", reflect.TypeOf((*MockStore)(nil).GetACategoryIDByName), arg0, arg1)
}

// GetAUser mocks base method.
func (m *MockStore) GetAUser(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAUser indicates an expected call of GetAUser.
func (mr *MockStoreMockRecorder) GetAUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAUser", reflect.TypeOf((*MockStore)(nil).GetAUser), arg0, arg1)
}

// GetAccount mocks base method.
func (m *MockStore) GetAccount(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockStoreMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockStore)(nil).GetAccount), arg0, arg1)
}

// GetAccountByEmail mocks base method.
func (m *MockStore) GetAccountByEmail(arg0 context.Context, arg1 string) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByEmail", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByEmail indicates an expected call of GetAccountByEmail.
func (mr *MockStoreMockRecorder) GetAccountByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByEmail", reflect.TypeOf((*MockStore)(nil).GetAccountByEmail), arg0, arg1)
}

// GetLastUserQuestionScore mocks base method.
func (m *MockStore) GetLastUserQuestionScore(arg0 context.Context, arg1 db.GetLastUserQuestionScoreParams) (db.UserQuestionScore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastUserQuestionScore", arg0, arg1)
	ret0, _ := ret[0].(db.UserQuestionScore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastUserQuestionScore indicates an expected call of GetLastUserQuestionScore.
func (mr *MockStoreMockRecorder) GetLastUserQuestionScore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastUserQuestionScore", reflect.TypeOf((*MockStore)(nil).GetLastUserQuestionScore), arg0, arg1)
}

// GetQuestion mocks base method.
func (m *MockStore) GetQuestion(arg0 context.Context, arg1 int64) (db.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestion", arg0, arg1)
	ret0, _ := ret[0].(db.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuestion indicates an expected call of GetQuestion.
func (mr *MockStoreMockRecorder) GetQuestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestion", reflect.TypeOf((*MockStore)(nil).GetQuestion), arg0, arg1)
}

// GetUserByUsername mocks base method.
func (m *MockStore) GetUserByUsername(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStoreMockRecorder) GetUserByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStore)(nil).GetUserByUsername), arg0, arg1)
}

// ListAccounts mocks base method.
func (m *MockStore) ListAccounts(arg0 context.Context, arg1 db.ListAccountsParams) ([]db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].([]db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockStoreMockRecorder) ListAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockStore)(nil).ListAccounts), arg0, arg1)
}

// ListAllQuestionScoresByUser mocks base method.
func (m *MockStore) ListAllQuestionScoresByUser(arg0 context.Context, arg1 db.ListAllQuestionScoresByUserParams) ([]db.UserQuestionScore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllQuestionScoresByUser", arg0, arg1)
	ret0, _ := ret[0].([]db.UserQuestionScore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllQuestionScoresByUser indicates an expected call of ListAllQuestionScoresByUser.
func (mr *MockStoreMockRecorder) ListAllQuestionScoresByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllQuestionScoresByUser", reflect.TypeOf((*MockStore)(nil).ListAllQuestionScoresByUser), arg0, arg1)
}

// ListAllQuestions mocks base method.
func (m *MockStore) ListAllQuestions(arg0 context.Context, arg1 db.ListAllQuestionsParams) ([]db.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllQuestions", arg0, arg1)
	ret0, _ := ret[0].([]db.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllQuestions indicates an expected call of ListAllQuestions.
func (mr *MockStoreMockRecorder) ListAllQuestions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllQuestions", reflect.TypeOf((*MockStore)(nil).ListAllQuestions), arg0, arg1)
}

// ListAllQuestionsByCategory mocks base method.
func (m *MockStore) ListAllQuestionsByCategory(arg0 context.Context, arg1 int64) ([]db.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllQuestionsByCategory", arg0, arg1)
	ret0, _ := ret[0].([]db.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllQuestionsByCategory indicates an expected call of ListAllQuestionsByCategory.
func (mr *MockStoreMockRecorder) ListAllQuestionsByCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllQuestionsByCategory", reflect.TypeOf((*MockStore)(nil).ListAllQuestionsByCategory), arg0, arg1)
}

// ListAllQuestionsByDifficulty mocks base method.
func (m *MockStore) ListAllQuestionsByDifficulty(arg0 context.Context, arg1 db.ListAllQuestionsByDifficultyParams) ([]db.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllQuestionsByDifficulty", arg0, arg1)
	ret0, _ := ret[0].([]db.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllQuestionsByDifficulty indicates an expected call of ListAllQuestionsByDifficulty.
func (mr *MockStoreMockRecorder) ListAllQuestionsByDifficulty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllQuestionsByDifficulty", reflect.TypeOf((*MockStore)(nil).ListAllQuestionsByDifficulty), arg0, arg1)
}

// ListAllQuestionsByType mocks base method.
func (m *MockStore) ListAllQuestionsByType(arg0 context.Context, arg1 db.ListAllQuestionsByTypeParams) ([]db.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllQuestionsByType", arg0, arg1)
	ret0, _ := ret[0].([]db.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllQuestionsByType indicates an expected call of ListAllQuestionsByType.
func (mr *MockStoreMockRecorder) ListAllQuestionsByType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllQuestionsByType", reflect.TypeOf((*MockStore)(nil).ListAllQuestionsByType), arg0, arg1)
}

// ListCategories mocks base method.
func (m *MockStore) ListCategories(arg0 context.Context, arg1 db.ListCategoriesParams) ([]db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCategories", arg0, arg1)
	ret0, _ := ret[0].([]db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCategories indicates an expected call of ListCategories.
func (mr *MockStoreMockRecorder) ListCategories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCategories", reflect.TypeOf((*MockStore)(nil).ListCategories), arg0, arg1)
}

// ListLastScoresByQuestion mocks base method.
func (m *MockStore) ListLastScoresByQuestion(arg0 context.Context, arg1 db.ListLastScoresByQuestionParams) ([]db.UserQuestionScore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLastScoresByQuestion", arg0, arg1)
	ret0, _ := ret[0].([]db.UserQuestionScore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLastScoresByQuestion indicates an expected call of ListLastScoresByQuestion.
func (mr *MockStoreMockRecorder) ListLastScoresByQuestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLastScoresByQuestion", reflect.TypeOf((*MockStore)(nil).ListLastScoresByQuestion), arg0, arg1)
}

// ListSingleQuestionScoresByUser mocks base method.
func (m *MockStore) ListSingleQuestionScoresByUser(arg0 context.Context, arg1 db.ListSingleQuestionScoresByUserParams) ([]db.UserQuestionScore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSingleQuestionScoresByUser", arg0, arg1)
	ret0, _ := ret[0].([]db.UserQuestionScore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSingleQuestionScoresByUser indicates an expected call of ListSingleQuestionScoresByUser.
func (mr *MockStoreMockRecorder) ListSingleQuestionScoresByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSingleQuestionScoresByUser", reflect.TypeOf((*MockStore)(nil).ListSingleQuestionScoresByUser), arg0, arg1)
}

// ListUsers mocks base method.
func (m *MockStore) ListUsers(arg0 context.Context, arg1 db.ListUsersParams) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0, arg1)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockStoreMockRecorder) ListUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockStore)(nil).ListUsers), arg0, arg1)
}

// UpdateAccountEmail mocks base method.
func (m *MockStore) UpdateAccountEmail(arg0 context.Context, arg1 db.UpdateAccountEmailParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccountEmail", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccountEmail indicates an expected call of UpdateAccountEmail.
func (mr *MockStoreMockRecorder) UpdateAccountEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountEmail", reflect.TypeOf((*MockStore)(nil).UpdateAccountEmail), arg0, arg1)
}

// UpdateAccountName mocks base method.
func (m *MockStore) UpdateAccountName(arg0 context.Context, arg1 db.UpdateAccountNameParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccountName", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccountName indicates an expected call of UpdateAccountName.
func (mr *MockStoreMockRecorder) UpdateAccountName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountName", reflect.TypeOf((*MockStore)(nil).UpdateAccountName), arg0, arg1)
}

// UpdateAccountOptedIn mocks base method.
func (m *MockStore) UpdateAccountOptedIn(arg0 context.Context, arg1 db.UpdateAccountOptedInParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccountOptedIn", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccountOptedIn indicates an expected call of UpdateAccountOptedIn.
func (mr *MockStoreMockRecorder) UpdateAccountOptedIn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountOptedIn", reflect.TypeOf((*MockStore)(nil).UpdateAccountOptedIn), arg0, arg1)
}

// UpdateAdminStatus mocks base method.
func (m *MockStore) UpdateAdminStatus(arg0 context.Context, arg1 db.UpdateAdminStatusParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdminStatus", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAdminStatus indicates an expected call of UpdateAdminStatus.
func (mr *MockStoreMockRecorder) UpdateAdminStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminStatus", reflect.TypeOf((*MockStore)(nil).UpdateAdminStatus), arg0, arg1)
}

// UpdateCategory mocks base method.
func (m *MockStore) UpdateCategory(arg0 context.Context, arg1 db.UpdateCategoryParams) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategory indicates an expected call of UpdateCategory.
func (mr *MockStoreMockRecorder) UpdateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockStore)(nil).UpdateCategory), arg0, arg1)
}

// UpdateLatestUserQuestionScore mocks base method.
func (m *MockStore) UpdateLatestUserQuestionScore(arg0 context.Context, arg1 db.UpdateLatestUserQuestionScoreParams) (db.UserQuestionScore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLatestUserQuestionScore", arg0, arg1)
	ret0, _ := ret[0].(db.UserQuestionScore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLatestUserQuestionScore indicates an expected call of UpdateLatestUserQuestionScore.
func (mr *MockStoreMockRecorder) UpdateLatestUserQuestionScore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLatestUserQuestionScore", reflect.TypeOf((*MockStore)(nil).UpdateLatestUserQuestionScore), arg0, arg1)
}

// UpdatePassword mocks base method.
func (m *MockStore) UpdatePassword(arg0 context.Context, arg1 db.UpdatePasswordParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockStoreMockRecorder) UpdatePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockStore)(nil).UpdatePassword), arg0, arg1)
}

// UpdateQuestion mocks base method.
func (m *MockStore) UpdateQuestion(arg0 context.Context, arg1 db.UpdateQuestionParams) (db.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuestion", arg0, arg1)
	ret0, _ := ret[0].(db.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateQuestion indicates an expected call of UpdateQuestion.
func (mr *MockStoreMockRecorder) UpdateQuestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuestion", reflect.TypeOf((*MockStore)(nil).UpdateQuestion), arg0, arg1)
}

// UpdateUserGrade mocks base method.
func (m *MockStore) UpdateUserGrade(arg0 context.Context, arg1 db.UpdateUserGradeParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserGrade", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserGrade indicates an expected call of UpdateUserGrade.
func (mr *MockStoreMockRecorder) UpdateUserGrade(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserGrade", reflect.TypeOf((*MockStore)(nil).UpdateUserGrade), arg0, arg1)
}

// UpdateUserUserQuestionScore mocks base method.
func (m *MockStore) UpdateUserUserQuestionScore(arg0 context.Context, arg1 db.UpdateUserUserQuestionScoreParams) ([]db.UserQuestionScore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserUserQuestionScore", arg0, arg1)
	ret0, _ := ret[0].([]db.UserQuestionScore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserUserQuestionScore indicates an expected call of UpdateUserUserQuestionScore.
func (mr *MockStoreMockRecorder) UpdateUserUserQuestionScore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserUserQuestionScore", reflect.TypeOf((*MockStore)(nil).UpdateUserUserQuestionScore), arg0, arg1)
}

// UpdateUsername mocks base method.
func (m *MockStore) UpdateUsername(arg0 context.Context, arg1 db.UpdateUsernameParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUsername", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUsername indicates an expected call of UpdateUsername.
func (mr *MockStoreMockRecorder) UpdateUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUsername", reflect.TypeOf((*MockStore)(nil).UpdateUsername), arg0, arg1)
}
