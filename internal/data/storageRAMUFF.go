package data

import (
	"server-rest-study/pkg/file"
)

func NewStorageRAMUFF() (Storage, error) {
	UserCurator, err := NewUserCuratorRAMU()
	return &storageRAMUFF{UserCurator, NewCuratorFF()}, err
}

// storageRAMUFF RAM store users info, file system store users files
type storageRAMUFF struct {
	*CuratorRAMU
	file.FileCurator
}

func (c *storageRAMUFF) DeleteFile(path string) error {
	c.CuratorRAMU.DeleteFile(path)
	return c.FileCurator.DeleteFile(path)
}
