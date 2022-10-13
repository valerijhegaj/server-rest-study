package data

import "server-rest-study/pkg/file"

func InitStorage() error {
	return nil
}

func GetStorage() Storage {
	return NewTestStorage()
}

type Storage interface {
	NewUser(name, password string) error
	NewSession(userID int, token string) error

	file.FileCurator
}

func NewTestStorage() *TestStorage {
	return &TestStorage{file.NewFileCurator()}
}

type TestStorage struct {
	file.FileCurator
}

func (c *TestStorage) NewUser(name, password string) error {
	return nil
}

func (c *TestStorage) NewSession(userID int, token string) error {
	return nil
}
