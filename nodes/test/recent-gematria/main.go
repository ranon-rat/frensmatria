package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	f, _ := os.Open("./test-file.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		base, err := url.Parse("http://localhost:8080/upload")
		if err != nil {
			return
		}

		// Query params
		params := url.Values{}
		params.Add("word-input", scanner.Text())
		base.RawQuery = params.Encode()

		_, err = http.Get(base.String())
		if err != nil {
			fmt.Println(err)
		}

	}

}
