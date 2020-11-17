package mock_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/otus/mock"
	"github.com/stretchr/testify/suite"
)

type userMatcher struct {
	mock.User
}

func (u userMatcher) Matches(x interface{}) bool {
	u2, ok := x.(mock.User)
	if !ok {
		return false
	}
	u2.ID = u.ID
	return u2 == u.User
}

func (u userMatcher) String() string {
	return fmt.Sprintf("is equal to %+v", u.User)
}

type StoreSuite struct {
	suite.Suite
	mockCtl *gomock.Controller
	mockDB  *MockUsersDB
	store   *mock.UserStore
}

func (s *StoreSuite) SetupTest() {
	s.mockCtl = gomock.NewController(s.T())
	s.mockDB = NewMockUsersDB(s.mockCtl)
	s.store = mock.NewUserStore(s.mockDB)
}

func (s *StoreSuite) TeardownTest() {
	s.mockCtl.Finish()
}

func (s *StoreSuite) TestDuplicate() {
	user1 := mock.User{
		ID:    "test_user_1",
		Name:  "test_name_1",
		Phone: "8-911-234-4567",
	}

	s.mockDB.EXPECT().FindUser(user1.ID).Return(user1, nil)
	s.mockDB.EXPECT().AddUser(userMatcher{user1}).Return(nil)
	newID, err := s.store.Duplicate(user1.ID)

	s.Require().NoError(err)
	s.Require().NotEqual(newID, user1.ID)
}

var errTest = errors.New("test error")

func (s *StoreSuite) TestDuplicateErr() {
	user1 := mock.User{
		ID:    "test_user_1",
		Name:  "test_name_1",
		Phone: "8-911-234-4567",
	}

	s.mockDB.EXPECT().FindUser(user1.ID).Return(user1, nil)
	s.mockDB.EXPECT().AddUser(gomock.Any()).Return(errTest)
	_, err := s.store.Duplicate(user1.ID)

	s.Require().EqualError(err, errTest.Error())
}

func TestStoreSuire(t *testing.T) {
	suite.Run(t, new(StoreSuite))
}
