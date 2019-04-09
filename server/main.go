package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ealipio/app-with-gorm/log"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, int(42), int64(100))

	log.Prinln(ctx, "Handler started")
	defer log.Prinln(ctx, "handler ended")

	fmt.Printf("value for foo is %v", ctx.Value("foo"))

	select {
	case <-time.After(5 * time.Second):
		fmt.Println(w, "Hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Prinln(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// time.Sleep(5 * time.Second)
		// fmt.Print(w, "Hello")

	}
}
