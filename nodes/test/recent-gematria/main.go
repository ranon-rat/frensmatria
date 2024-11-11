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
		text := scanner.Text()

		fmt.Println(text)
		base, err := url.Parse("http://localhost:8080/upload")
		if err != nil {
			return
		}

		// Query params
		params := url.Values{}
		params.Add("word-input", text)
		base.RawQuery = params.Encode()

		http.Get(base.String())

	}

}
