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
	NewFile(file io.ReadCloser, userID int, path string) error
	UpdateFile(file io.ReadCloser, userID int, path string) error
	DeleteFile(userID int, path string) error
}

func FormatPath(userID int, path string) string {
	return fmt.Sprintf("./cmd/users_data/%d/%s", userID, path)
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

func CreateDirsToFile(path string) error {
	path, _ = separateFileAndPath(path)
	return os.MkdirAll(path, 0700)
}

type fileCuratorPrimitive struct {
}

func (c *fileCuratorPrimitive) GetFile(userID int, path string) (
	io.ReadCloser,
	error,
) {
	return os.Open(FormatPath(userID, path))
}

func (c *fileCuratorPrimitive) NewFile(
	reader io.ReadCloser, userID int, path string,
) error {
	path = FormatPath(userID, path)
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
