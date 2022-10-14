package file

import "encoding/json"

type Helper struct {
	Password string `json:"password"`
	UserName string `json:"username"`
	UserID   int    `json:"user_id"`
	Token    string `json:"token"`
	FileData string `json:"file_data"`
	Path     string `json:"path"`
	Rights   string `json:"rights"`
}

func Parse(data []byte) (
	Helper,
	error,
) {
	var h Helper
	err := json.Unmarshal(data, &h)
	return h, err
}
