package data

func InitStorage() error {
	return nil
}

func GetStorage() Storage {
	return NewTestStorage()
}

type Storage interface {
	NewUser(name, password string) error
	NewSession(userID int, token string) error

	fileCurator
}

func NewTestStorage() *TestStorage {
	return &TestStorage{NewFileCurator()}
}

type TestStorage struct {
	fileCurator
}

func (c *TestStorage) NewUser(name, password string) error {
	return nil
}

func (c *TestStorage) NewSession(userID int, token string) error {
	return nil
}
