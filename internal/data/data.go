package data

import "server-rest-study/pkg/file"

func InitStorage() error {
	return nil
}

func GetStorage() Storage {
	return NewTestStorage()
}

type Storage interface {
	NewUser(name, password string) (
		int,
		error,
	)
	NewToken(userID int, password string) (
		string,
		error,
	)
	NewSession(userID int, token string) error

	// in future will be update to check user access
	// by creating adapter
	file.FileCurator
}

func NewTestStorage() *TestStorage {
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

func (c *TestStorage) NewToken(
	userID int,
	password string,
) (
	string,
	error,
) {
	return "1", nil
}
