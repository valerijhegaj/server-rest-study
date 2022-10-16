package file

import (
	"encoding/json"
)

type Helper struct {
	Password string `json:"password,omitempty"`
	UserName string `json:"username,omitempty"`
	FileData string `json:"file_data,omitempty"`
	Path     string `json:"path,omitempty"`
	Rights   string `json:"rights,omitempty"`
	MaxAge   int    `json:"max_age,omitempty"`
}

func Parse(data []byte) (
	Helper,
	error,
) {
	var h Helper
	err := json.Unmarshal(data, &h)
	return h, err
}

func UnParse(data Helper) []byte {
	rawData, _ := json.Marshal(data)
	return rawData
}
