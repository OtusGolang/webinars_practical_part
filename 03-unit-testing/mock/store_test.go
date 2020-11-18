package mock_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/otus/mock"
	"github.com/stretchr/testify/require"
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

func TestDuplicate(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockDB := NewMockUsersDB(mockCtl)
	store := mock.NewUserStore(mockDB)

	user1 := mock.User{
		ID:    "test_user_1",
		Name:  "test_name_1",
		Phone: "8-911-234-4567",
	}

	mockDB.EXPECT().FindUser(user1.ID).Return(user1, nil)
	mockDB.EXPECT().AddUser(userMatcher{user1}).Return(nil)
	newID, err := store.Duplicate(user1.ID)

	require.NoError(t, err)
	require.NotEqual(t, newID, user1.ID)
}

var errAddUser = errors.New("test error")

func TestDuplicateErr(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockDB := NewMockUsersDB(mockCtl)
	store := mock.NewUserStore(mockDB)

	user1 := mock.User{
		ID:    "test_user_1",
		Name:  "test_name_1",
		Phone: "8-911-234-4567",
	}

	mockDB.EXPECT().FindUser(user1.ID).Return(user1, nil)
	mockDB.EXPECT().AddUser(gomock.Any()).Return(errAddUser)
	_, err := store.Duplicate(user1.ID)

	require.EqualError(t, err, errAddUser.Error())
}
