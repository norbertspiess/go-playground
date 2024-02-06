package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Request failed", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
