package file

import (
	"io"
	"os"
	"testing"

	"server-rest-study/test/format"
)

func Test_formatPath(t *testing.T) {
	path := "./cmd/users_data/10/name/name.txt"
	createdPath := FormatPath(10, "name/name.txt")
	if createdPath != path {
		t.Error(format.ErrorString(path, createdPath))
	}
}

func Test_separateFileAndPath(t *testing.T) {
	test := func(realPath, realName string) {
		path, name := separateFileAndPath(realPath + "/" + realName)
		if path != realPath {
			t.Error(format.ErrorString(realPath, path))
		}
		if name != realName {
			t.Error(format.ErrorString(realName, name))
		}
	}

	test("a", "c.txt")
	test("a/b", "c.txt")
	test("", "c.txt")
}

func TestFileCuratorPrimitive_NewFile(t *testing.T) {
	os.Chdir("../..")
	defer os.Chdir("pkg/file")

	test := func(userID int, name, data string) {
		fc := NewFileCurator()
		err := fc.NewFile(NewReadCloserFromString(data), userID, name)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}

		var rc io.ReadCloser
		rc, err = os.Open(FormatPath(userID, name))
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}

		buf := make([]byte, len(data)+1)
		n, err := rc.Read(buf)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if n != len(data) {
			t.Error(format.ErrorInt(len(data), n))
		}
		if string(buf[:n]) != data {
			t.Error(format.ErrorString(string(buf[:n]), data))
		}

		err = rc.Close()
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}

		deletePath := FormatPath(userID, "")
		deletePath = deletePath[:len(deletePath)-1]
		os.RemoveAll(deletePath)
	}

	test(1337, "test.txt", "some data")
	test(1337, "dir1/dir2/test.txt", "some data")
	veryBigText := "data"
	for i := 0; i < 15; i++ {
		veryBigText = veryBigText + veryBigText
	}
	test(1337, "test.txt", veryBigText)
}

func TestFileCuratorPrimitive_GetFile(t *testing.T) {
	os.Chdir("../..")
	defer os.Chdir("pkg/file")

	test := func(userID int, name, data string) {
		fc := NewFileCurator()
		err := fc.NewFile(NewReadCloserFromString(data), userID, name)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}

		rc, err := fc.GetFile(userID, name)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}

		buf := make([]byte, len(data)+1)
		n, err := rc.Read(buf)
		if n != len(data) {
			t.Error(format.ErrorInt(len(data), n))
		}
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		redData := string(buf[:n])
		if data != redData {
			t.Error(format.ErrorString(data, redData))
		}

		rc.Close()

		deletePath := FormatPath(userID, "")
		deletePath = deletePath[:len(deletePath)-1]
		os.RemoveAll(deletePath)
	}

	test(1337, "test.txt", "some data")
	test(1337, "dir1/dir2/test.txt", "some data")
	veryBigText := "data"
	for i := 0; i < 15; i++ {
		veryBigText = veryBigText + veryBigText
	}
	test(1337, "test.txt", veryBigText)
}

func TestFileCuratorPrimitive_UpdateFile(t *testing.T) {
	os.Chdir("../..")
	defer os.Chdir("pkg/file")
	test := func(userID int, name, data string) {
		fc := NewFileCurator()
		err := fc.NewFile(
			NewReadCloserFromString("old data"),
			userID,
			name,
		)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}

		err = fc.UpdateFile(NewReadCloserFromString(data), userID, name)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}

		rc, err := fc.GetFile(userID, name)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}

		buf := make([]byte, len(data)+1)
		n, err := rc.Read(buf)
		if n != len(data) {
			t.Error(format.ErrorInt(len(data), n))
		}
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		redData := string(buf[:n])
		if data != redData {
			t.Error(format.ErrorString(data, redData))
		}

		rc.Close()

		deletePath := FormatPath(userID, "")
		deletePath = deletePath[:len(deletePath)-1]
		os.RemoveAll(deletePath)
	}

	test(1337, "test.txt", "some data")
	test(1337, "dir1/dir2/test.txt", "some data")
	veryBigText := "data"
	for i := 0; i < 5; i++ {
		veryBigText = veryBigText + veryBigText
	}
	test(1337, "test.txt", veryBigText)
}

func TestFileCuratorPrimitive_DeleteFile(t *testing.T) {
	os.Chdir("../..")
	defer os.Chdir("pkg/file")
	test := func(userID int, name1, name2 string) {
		fc := NewFileCurator()
		err := fc.NewFile(
			NewReadCloserFromString("old1 data"),
			userID,
			name1,
		)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}

		err = fc.NewFile(
			NewReadCloserFromString("old2 data"),
			userID,
			name2,
		)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}

		err = fc.DeleteFile(userID, name1)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}

		var rc io.ReadCloser
		rc, err = fc.GetFile(userID, name2)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		data := "old2 data"
		buf := make([]byte, len(data)+1)
		n, err := rc.Read(buf)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if n != len(data) {
			t.Error(format.ErrorInt(n, len(data)))
		}
		redData := string(buf[:n])
		if redData != data {
			t.Error(format.ErrorString(data, redData))
		}
		rc.Close()

		deletePath := FormatPath(userID, "")
		deletePath = deletePath[:len(deletePath)-1]
		os.RemoveAll(deletePath)
	}

	test(1337, "test.txt", "test2.txt")
	test(1337, "dir1/dir2/test.txt", "dir1/dir2/test2.txt")
}
