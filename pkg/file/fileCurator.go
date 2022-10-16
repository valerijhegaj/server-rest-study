package file

import (
	"io"
	"os"
)

func NewFileCurator() FileCurator {
	return &fileCuratorPrimitive{}
}

type FileCurator interface {
	GetFile(path string) (
		io.ReadCloser,
		error,
	)
	NewFile(file io.ReadCloser, path string) error
	UpdateFile(file io.ReadCloser, path string) error
	DeleteFile(path string) error
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

func (c *fileCuratorPrimitive) GetFile(path string) (
	io.ReadCloser,
	error,
) {
	return os.Open(path)
}

func (c *fileCuratorPrimitive) NewFile(
	reader io.ReadCloser, path string,
) error {
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
	reader io.ReadCloser, path string,
) error {
	return c.NewFile(reader, path)
}

func (c *fileCuratorPrimitive) DeleteFile(
	path string,
) error {
	return os.RemoveAll(path)
}
