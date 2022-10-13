package data

import "server-rest-study/pkg/file"

func NewStorageRUFF() Storage {
	return &StorageRUFF{file.NewFileCurator()}
}

// StorageRUFF Ram store info about Users, File system store user Files
type StorageRUFF struct {
	file.FileCurator
}

func (*StorageRUFF) NewUser(name, password string) (int, error) {
	return 0, nil
}

func (*StorageRUFF) NewToken(
	userID int, password string,
) (string, error) {
	return "", nil
}

func (*StorageRUFF) NewSession(userID int, token string) error {
	return nil
}

func (*StorageRUFF) CheckAccess(
	token string, userID int, path string, rights string,
) (bool, error) {
	return true, nil
}
