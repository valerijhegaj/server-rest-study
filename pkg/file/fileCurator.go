package file

import (
	"fmt"
	"io"
	"os"
)

func NewFileCuratorFF() FileCurator {
	return &fileCuratorPrimitive{}
}

type FileCurator interface {
	GetFile(userId int, path string) (
		io.ReadCloser,
		error,
	)
	NewFile(file io.ReadCloser, userID int, name string) error
	UpdateFile(file io.ReadCloser, userID int, name string) error
	DeleteFile(userID int, name string) error
}

func FormatPath(userID int, name string) string {
	return fmt.Sprintf("./cmd/users_data/%d/%s", userID, name)
}

func separateFileAndPath(path string) (string, string) {
	firstSlash := 0
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			firstSlash = i
			break
		}
	}
	return path[:firstSlash], path[firstSlash+1:]
}

func CreateDirsToFile(name string) error {
	path, _ := separateFileAndPath(name)
	return os.MkdirAll(path, 0700)
}

type fileCuratorPrimitive struct {
}

func (c *fileCuratorPrimitive) GetFile(userID int, name string) (
	io.ReadCloser,
	error,
) {
	return os.Open(FormatPath(userID, name))
}

func (c *fileCuratorPrimitive) NewFile(
	reader io.ReadCloser, userID int, name string,
) error {
	path := FormatPath(userID, name)
	err := CreateDirsToFile(path)
	var writer io.WriteCloser
	writer, err = os.Create(path)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, reader)
	reader.Close()
	writer.Close()

	return err
}
func (c *fileCuratorPrimitive) UpdateFile(
	reader io.ReadCloser, userID int, path string,
) error {
	return c.NewFile(reader, userID, path)
}

func (c *fileCuratorPrimitive) DeleteFile(
	userID int, path string,
) error {
	return os.RemoveAll(FormatPath(userID, path))
}
