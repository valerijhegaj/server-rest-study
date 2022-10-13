package file

import (
	"io"
	"testing"

	"server-rest-study/test/format"
)

func TestTestStringReadCloser_Read(t *testing.T) {
	testData := "adfasdfasfd"
	bufSize := 6

	rc := NewReadCloser(testData)
	buf := make([]byte, bufSize)

	n, err := rc.Read(buf)
	if n != bufSize {
		t.Error(format.ErrorInt(bufSize, n))
	}
	if err != nil {
		t.Error(format.ErrorString("without errors", err.Error()))
	}
	readData := string(buf[:n])
	if readData != testData[:bufSize] {
		t.Error(format.ErrorString(testData, readData))
	}

	n, err = rc.Read(buf)
	if n != 5 {
		t.Error(format.ErrorInt(5, n))
	}
	if err != nil {
		t.Error(format.ErrorString("without errors", err.Error()))
	}
	readData = string(buf[:n])
	if readData != testData[bufSize:] {
		t.Error(format.ErrorString(testData[bufSize:], readData))
	}

	n, err = rc.Read(buf)
	if err != io.EOF {
		t.Error(format.ErrorString(io.EOF.Error(), err.Error()))
	}
}
