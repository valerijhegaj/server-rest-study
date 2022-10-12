package data

import (
	"io"
	"os"
)

func NewFileCurator() fileCurator {
	return &fileCuratorPrimitive{}
}

type fileCurator interface {
	GetFile(userId int, path string) (
		io.ReadCloser,
		error,
	)
	NewFile(file io.ReadCloser, userID int, name string) error
	UpdateFile(file io.ReadCloser, userID int, name string) error
	DeleteFile(userID int, name string) error
}

func createPath(userID int, name string) string {
	return "users_data/" + string(rune(userID)) + "/" + name
}

type fileCuratorPrimitive struct {
}

func (c *fileCuratorPrimitive) GetFile(
	userID int,
	name string,
) (
	io.ReadCloser,
	error,
) {
	return os.Open(createPath(userID, name))
}

func (c *fileCuratorPrimitive) NewFile(
	reader io.ReadCloser,
	userID int,
	name string,
) error {
	writer, err := os.Create(createPath(userID, name))
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, reader)
	reader.Close()
	return err
}
func (c *fileCuratorPrimitive) UpdateFile(
	reader io.ReadCloser,
	userID int,
	path string,
) error {
	writer, err := os.Open(createPath(userID, path))
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, reader)
	reader.Close()
	return err
}

func (c *fileCuratorPrimitive) DeleteFile(
	userID int,
	path string,
) error {
	return os.RemoveAll(createPath(userID, path))
}
