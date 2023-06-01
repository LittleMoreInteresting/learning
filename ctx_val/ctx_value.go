package main

import (
	"context"
	"fmt"
	"time"
)

func HandelRequest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running, parameter: ", ctx.Value(contextTxKey{code: "b"}))
			time.Sleep(2 * time.Second)
		}
	}
}

type contextTxKey struct{ code string }

func main() {
	ctx := context.WithValue(context.Background(), contextTxKey{code: "a"}, "1")
	go HandelRequest(ctx)

	time.Sleep(10 * time.Second)
}
