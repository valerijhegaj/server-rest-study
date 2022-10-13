package file

import "encoding/json"

type Helper struct {
	Password string `json:"password"`
	User     string `json:"user"`
	UserID   int    `json:"user_id"`
	Token    string `json:"token"`
	FileData string `json:"file_data"`
}

func Parse(data []byte) (
	Helper,
	error,
) {
	var h Helper
	err := json.Unmarshal(data, &h)
	return h, err
}
