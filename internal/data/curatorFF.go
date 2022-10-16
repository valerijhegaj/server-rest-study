package data

import (
	"io"

	"server-rest-study/pkg/file"
)

func NewCuratorFF() file.FileCurator {
	return &CuratorFF{file.NewFileCurator()}
}

type CuratorFF struct {
	file.FileCurator
}

func (c *CuratorFF) transToSystemPath(webPath string) string {
	return "cmd/users_data/" + webPath
}

func (c *CuratorFF) GetFile(path string) (io.ReadCloser, error) {
	return c.FileCurator.GetFile(c.transToSystemPath(path))
}

func (c *CuratorFF) NewFile(file io.ReadCloser, path string) error {
	return c.FileCurator.NewFile(file, c.transToSystemPath(path))
}
func (c *CuratorFF) UpdateFile(
	file io.ReadCloser, path string,
) error {
	return c.FileCurator.UpdateFile(file, c.transToSystemPath(path))
}
func (c *CuratorFF) DeleteFile(path string) error {
	return c.FileCurator.DeleteFile(c.transToSystemPath(path))
}
