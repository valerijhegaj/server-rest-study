package data

import (
	"errors"
	"fmt"

	"server-rest-study/pkg/file"
)

func NewStorageRAMUFF() (Storage, error) {
	UserCurator, err := NewUserCuratorRAMU()
	return &storageRAMUFF{UserCurator, file.NewFileCuratorFF()}, err
}

// storageRAMUFF RAM store users info, file system store users files
type storageRAMUFF struct {
	*CuratorRAMU
	file.FileCurator
}

func (c *storageRAMUFF) DeleteFile(userID int, name string) error {
	c.CuratorRAMU.DeleteFile(userID, name)
	return c.FileCurator.DeleteFile(userID, name)
}

func NewUserCuratorRAMU() (*CuratorRAMU, error) {
	return &CuratorRAMU{
		userID:     make(map[string]int),
		username:   make(map[int]string),
		password:   make(map[int]string),
		session:    make(map[string]int),
		access:     make(map[string]map[int]string),
		usersCount: 1,
	}, nil
}

// CuratorRAMU Ram store info about Users
type CuratorRAMU struct {
	userID   map[string]int
	username map[int]string
	password map[int]string
	session  map[string]int
	// path -> userid -> rights
	access     map[string]map[int]string
	usersCount int
}

func (c *CuratorRAMU) NewUser(username, password string) (
	int, error,
) {
	_, ok := c.userID[username]
	if ok {
		return 0, errors.New("user already exist")
	}

	c.usersCount++
	userID := c.usersCount

	c.username[userID] = username
	c.password[userID] = password
	c.userID[username] = userID

	return userID, nil
}

func (c *CuratorRAMU) NewToken(userID int, password string) (
	string, error,
) {
	if c.password[userID] != password {
		return "", errors.New("incorrect password")
	}

	token, ok := c.username[userID]
	if !ok {
		return "", errors.New("user not exist")
	}

	return token, nil
}

func (c *CuratorRAMU) NewSession(userID int, token string) error {
	_, ok := c.session[token]
	if ok {
		return errors.New("session already exist")
	}
	c.session[token] = userID
	return nil
}

func (c *CuratorRAMU) CheckAccess(
	token string, userID int, path string, askerRights string,
) (bool, error) {
	askerUserID := c.session[token]
	if askerUserID == userID {
		return true, nil
	}

	realRights := c.access[fmt.Sprintf(
		"%d/%s", userID, path,
	)][askerUserID]

	switch len(realRights) {
	case 2:
		if realRights[0] == askerRights[0] || realRights[1] == askerRights[0] {
			return true, nil
		}
	case 1:
		if realRights[0] == askerRights[0] {
			return true, nil
		}
	}
	return false, nil
}

func (c *CuratorRAMU) DeleteFile(userID int, path string) {
	delete(c.access, fmt.Sprintf("%d/%s", userID, path))
}
