package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type key string

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, key("foo"), "bar")

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}
	io.Copy(os.Stdout, res.Body)
}
