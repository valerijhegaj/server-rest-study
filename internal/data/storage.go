package data

import "server-rest-study/pkg/file"

var GlobalStorage Storage

func InitStorage() error {
	GlobalStorage = NewTestStorage()
	return nil
}

func GetStorage() Storage {
	return GlobalStorage
}

type Storage interface {
	NewUser(name, password string) (int, error)
	NewToken(userID int, password string) (string, error)
	NewSession(userID int, token string) error
	CheckAccess(
		token string, userID int, path string, rights string,
	) (bool, error)

	file.FileCurator
}

func NewTestStorage() Storage {
	return &TestStorage{file.NewFileCurator()}
}

type TestStorage struct {
	file.FileCurator
}

func (c *TestStorage) NewUser(name, password string) (
	int,
	error,
) {
	return 1, nil
}

func (c *TestStorage) NewSession(userID int, token string) error {
	return nil
}

func (c *TestStorage) NewToken(userID int, password string) (
	string, error,
) {
	return "1", nil
}

func (c *TestStorage) CheckAccess(
	token string, userID int, path string, rights string,
) (bool, error) {
	return true, nil
}
