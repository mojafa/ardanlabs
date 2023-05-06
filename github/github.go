package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/mojafa")
	if err != nil {
		log.Fatalf("Error: %s", err)
		// log.Printf(error: %s, err)
		// os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", resp.Status)
	}
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	io.Copy(os.Stdout, resp.Body)
}
