package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a, _ := json.Marshal(
		struct {
			Name string
		}{"name"},
	)
	fmt.Println(string(a))
}
