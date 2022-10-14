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
	NewUser(name, password string) (int, error)
	NewToken(userID int, password string) (string, error)
	NewSession(userID int, token string) error
	CheckAccess(
		token string, userID int, path string, rights string,
	) (bool, error)
}

func NewStorageNUFF() (Storage, error) {
	return &storage{NewUserCuratorNU(), file.NewFileCuratorFF()}, nil
}

type storage struct {
	UserCurator
	file.FileCurator
}

func NewUserCuratorNU() UserCurator {
	return &CuratorNU{}
}

// CuratorNU nothing store info about users
type CuratorNU struct{}

func (c *CuratorNU) NewUser(name, password string) (
	int,
	error,
) {
	return 1, nil
}

func (c *CuratorNU) NewSession(userID int, token string) error {
	return nil
}

func (c *CuratorNU) NewToken(userID int, password string) (
	string, error,
) {
	return "1", nil
}

func (c *CuratorNU) CheckAccess(
	token string, userID int, path string, rights string,
) (bool, error) {
	return true, nil
}
