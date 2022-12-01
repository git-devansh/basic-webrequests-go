package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://pkg.go.dev/net/http"

func main() {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	databytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}
