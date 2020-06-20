package mock

//go:generate mockgen -destination=mock_db_test.go -package=mock_test . UsersDB

type UsersDB interface {
	AddUser(user User) error
	FindUser(id string) (User, error)
}
