package mock_test

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
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

	var seed int64 = time.Now().UnixNano()
	s.T().Logf("rand seed: %d", seed)
	rand.Seed(seed)
}

func (s *StoreSuite) TeardownTest() {
	s.mockCtl.Finish()
}

func (s *StoreSuite) TestDuplicate() {
	user1 := s.fakeUser()

	s.mockDB.EXPECT().FindUser(user1.ID).Return(user1, nil)
	s.mockDB.EXPECT().AddUser(userMatcher{user1}).Return(nil)
	newID, err := s.store.Duplicate(user1.ID)

	s.Require().NoError(err)
	s.Require().NotEqual(newID, user1.ID)
}

func (s *StoreSuite) TestDuplicateErrAdd() {
	user1 := s.fakeUser()

	testErr := errors.New("test error")
	s.mockDB.EXPECT().FindUser(user1.ID).Return(user1, nil)
	s.mockDB.EXPECT().AddUser(userMatcher{user1}).Return(testErr)
	_, err := s.store.Duplicate(user1.ID)

	s.Require().EqualError(err, testErr.Error())
}

func (s *StoreSuite) TestDuplicateErrFind() {
	user1 := s.fakeUser()

	testErr := errors.New("test error")
	s.mockDB.EXPECT().FindUser(user1.ID).Return(user1, testErr)
	_, err := s.store.Duplicate(user1.ID)

	s.Require().True(errors.Is(err, testErr))
}

func (*StoreSuite) fakeUser() mock.User {
	return mock.User{
		ID:    faker.UUIDDigit(),
		Name:  faker.Name(),
		Phone: faker.Phonenumber(),
	}
}

func TestStoreSuire(t *testing.T) {
	suite.Run(t, new(StoreSuite))
}
