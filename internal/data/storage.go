package data

import "server-rest-study/pkg/file"

var GlobalStorage Storage

func InitStorage() error {
	var err error
	GlobalStorage, err = NewStorageRAMUFF()
	return err
}

func GetStorage() Storage {
	return GlobalStorage
}

type Storage interface {
	UserCurator
	file.FileCurator
}

type UserCurator interface {
	NewUser(username, password string) (int, error)
	NewToken(username string, password string) (string, error)
	NewSession(username string, token string) error
	CheckAccess(
		token string, userID int, path string, rights string,
	) (bool, error)
	SetRights(token string, userID int, path, rights string) error
}

const (
	NotAuthorized = "not exist session"
)
