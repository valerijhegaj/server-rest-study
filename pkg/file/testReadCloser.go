package file

import "io"

func NewReadCloser(data string) io.ReadCloser {
	return &TestStringReadCloser{data, 0}
}

type TestStringReadCloser struct {
	data string
	ptr  int
}

func (c *TestStringReadCloser) Read(b []byte) (
	int,
	error,
) {
	if len(c.data) == 0 {
		return 0, io.EOF
	}
	n := copy(b, c.data)
	c.data = c.data[n:]
	return n, nil
}

func (c *TestStringReadCloser) Close() error {
	return nil
}
