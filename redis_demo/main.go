package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

const script_del = "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end"

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.72.128:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	/*result, err := rdb.Set(ctx, "r1", "111111", 60*time.Second).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)*/
	get, err := rdb.Get(ctx, "r1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(get)

	eval, err := rdb.Eval(ctx, script_del, []string{"r1"}, "6").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(eval)
}
