package data

import (
	"errors"
	"fmt"
)

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

// CuratorRAMU Ram store info about Users unsafe when multithread work
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

	userID := c.usersCount
	c.usersCount++

	c.username[userID] = username
	c.password[userID] = password
	c.userID[username] = userID

	return userID, nil
}

func (c *CuratorRAMU) NewToken(username, password string) (
	string, error,
) {
	userID, ok := c.userID[username]
	if !ok {
		return "", errors.New("user not exist")
	}
	if c.password[userID] != password {
		return "", errors.New("incorrect password")
	}
	token := username

	return token, nil
}

func (c *CuratorRAMU) NewSession(username, token string) error {
	_, ok := c.session[token]
	if ok {
		return errors.New("session already exist")
	}
	c.session[token] = c.userID[username]
	return nil
}

func (c *CuratorRAMU) CheckAccess(
	token string, path string, askerRights string,
) (bool, error) {
	askerUserID := c.session[token]
	ownerID := c.userID[c.getOwner(path)]
	if ownerID == askerUserID {
		return true, nil
	}

	realRights := c.access[path][askerUserID]

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

func (c *CuratorRAMU) DeleteFile(path string) {
	delete(c.access, path)
}

func (c *CuratorRAMU) SetRights(
	token string, username string, path, rights string,
) error {
	owner, ok := c.session[token]
	if !ok {
		return errors.New(NotAuthorized)
	}
	path = fmt.Sprintf("%s/%s", c.username[owner], path)
	_, ok = c.access[path]
	if !ok {
		c.access[path] = make(map[int]string)
	}
	c.access[path][c.userID[username]] = rights
	return nil
}

func (c *CuratorRAMU) getOwner(path string) string {
	ptr := len(path)
	for i := 0; i < ptr; i++ {
		if path[i] == '/' {
			ptr = i
		}
	}
	return path[:ptr]
}
