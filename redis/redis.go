package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	Redis := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"10.0.6.48:6380"},
	})

	// res := Redis.Do("set", "key", "value")
	// fmt.Println(res)

	// val, err := Redis.Get("keyll").Result()
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// fmt.Println(val)

	err := Redis.Set("key", "value", 0).Err() // key value ttl
	if err != nil {
		panic(err)
	}

	val, err := Redis.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := Redis.Get("missing_key").Result()
	if err == redis.Nil {
		fmt.Println("missing_key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("missing_key", val2)
	}
}