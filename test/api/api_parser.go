package apiParser

import (
	"fmt"
	"io"
	"net/http"

	"server-rest-study/pkg/file"
)

type User struct {
	Username, Password string
	cookies            []*http.Cookie

	client http.Client
	PORT   int
}

func (c *User) Register() (int, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://localhost:%d/api/user", c.PORT),
		file.NewReadCloserFromBytes(
			file.UnParse(
				file.Helper{
					UserName: c.Username, Password: c.Password,
				},
			),
		),
	)
	if err != nil {
		return -1, err
	}

	res, err := c.client.Do(req)
	return res.StatusCode, err
}

func (c *User) LogIn(maxAge int) (int, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://localhost:%d/api/session", c.PORT),
		file.NewReadCloserFromBytes(
			file.UnParse(
				file.Helper{
					UserName: c.Username, Password: c.Password, MaxAge: maxAge,
				},
			),
		),
	)
	if err != nil {
		return -1, err
	}

	res, err := c.client.Do(req)

	if res.StatusCode == http.StatusCreated {
		c.cookies = res.Cookies()
	}

	return res.StatusCode, err
}

func (c *User) GetFile(path string) (int, []byte, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://localhost:%d/api/files/%s", c.PORT, path),
		nil,
	)
	if err != nil {
		return -1, nil, err
	}
	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return -1, nil, err
	}

	body, err := io.ReadAll(res.Body)
	return res.StatusCode, body, err
}

func (c *User) CreateFile(path, data string) (int, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://localhost:%d/api/files/%s", c.PORT, path),
		file.NewReadCloserFromBytes(
			file.UnParse(
				file.Helper{
					FileData: data,
				},
			),
		),
	)
	if err != nil {
		return -1, err
	}
	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}

	res, err := c.client.Do(req)

	return res.StatusCode, err
}

func (c *User) UpdateFile(path, data string) (int, error) {
	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("http://localhost:%d/api/files/%s", c.PORT, path),
		file.NewReadCloserFromBytes(
			file.UnParse(
				file.Helper{
					FileData: data,
				},
			),
		),
	)
	if err != nil {
		return -1, err
	}
	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}

	res, err := c.client.Do(req)

	return res.StatusCode, err
}

func (c *User) DeleteFile(path string) (int, error) {
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("http://localhost:%d/api/files/%s", c.PORT, path),
		nil,
	)
	if err != nil {
		return -1, err
	}
	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}

	res, err := c.client.Do(req)

	return res.StatusCode, err
}

func (c *User) GiveAccess(path, username, rights string) (
	int, error,
) {
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			path = path[i+1:]
			break
		}
	}
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://localhost:%d/api/give_access", c.PORT),
		file.NewReadCloserFromBytes(
			file.UnParse(
				file.Helper{
					Path: path, UserName: username, Rights: rights,
				},
			),
		),
	)
	if err != nil {
		return -1, err
	}
	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}

	res, err := c.client.Do(req)

	return res.StatusCode, err
}
