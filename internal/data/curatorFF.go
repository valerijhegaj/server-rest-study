package data

import (
	"io"

	"server-rest-study/pkg/file"
)

func NewCuratorFF() file.FileCurator {
	return &CuratorFF{file.NewFileCuratorFF()}
}

type CuratorFF struct {
	file.FileCurator
}

func (c *CuratorFF) GetFile(userId int, path string) (io.ReadCloser, error) {
	return c.FileCurator.GetFile(userId, path)
}
func (c *CuratorFF) NewFile(file io.ReadCloser, userID int, name string) error {
	return c.FileCurator.NewFile(file, userID, name)
}
func (c *CuratorFF) UpdateFile(file io.ReadCloser, userID int, name string) error {
	return c.FileCurator.UpdateFile(file, userID, name)
}
func (c *CuratorFF) DeleteFile(userID int, name string) error {
	return c.FileCurator.DeleteFile(userID, name)
}
