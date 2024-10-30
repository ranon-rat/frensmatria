package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type ExampleGematria struct {
	Example string `json:"example"`
}

func main() {

	a := ExampleGematria{
		Example: "12341234",
	}
	b, _ := json.Marshal(a)
	dst := base64.RawStdEncoding.EncodeToString(b)
	fmt.Println(dst)
	dec, _ := base64.RawStdEncoding.DecodeString(dst)
	fmt.Println(string(dec))

}
