package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type key int

const requestIDKey = key(42)

// Prinln vvsdv
func Prinln(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIDKey).(int64)
	if !ok {
		log.Println("Could not find request ID in context")
		return
	}
	log.Printf("[%d] %s", id, msg)
}

// Decorate ssdvd
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id)
		f(w, r.WithContext(ctx))
	}
}
