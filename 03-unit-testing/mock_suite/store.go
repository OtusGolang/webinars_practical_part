package mock

import "fmt"

type UserStore struct {
	db UsersDB
}

func NewUserStore(db UsersDB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) Duplicate(userID string) (string, error) {
	user, err := s.db.FindUser(userID)
	if err != nil {
		return "", fmt.Errorf("failed to find user: %w", err)
	}

	user.ID = user.ID + "_"
	err = s.db.AddUser(user)
	return user.ID, err
}
