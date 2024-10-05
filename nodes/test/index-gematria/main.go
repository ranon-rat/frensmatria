package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	str := "hola"
	GetCombinations("", str)

}

// this is just for making it easier to obtain results

func GetCombinations(result, str string) {
	if str == "" {
		base, err := url.Parse("http://localhost:8080/upload")
		if err != nil {
			return
		}

		// Query params
		params := url.Values{}
		params.Add("word-input", result)
		base.RawQuery = params.Encode()
		_, err = http.Get(base.String())
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	for i := range str {
		GetCombinations(result+string(str[i]), str[:i]+str[i+1:])

	}

}
