package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/todaypay/go-commons/pkg/jwtauth"
)

var redisClient *redis.Client

func main() {
	fmt.Println("Go Redis Started......")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(context.Background()).Result()
	fmt.Println(pong, err)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiTUVSQ0hBTlQiLCJpZCI6Im1fU3JsYU5pdUJRRUxQaXp3IiwiaXNfcGFydGlhbCI6ZmFsc2UsImlzcyI6IjJkcCIsInN1YiI6IjJkcCIsImV4cCI6MTg4ODYzMjc3NywibmJmIjoxNjg5MTA5Mjc3LCJpYXQiOjE2ODg2MjkxNzcsImp0aSI6IjI5NjQxMjM1LTcwOTItNDU0ZS1iYzRiLTJmNzVkZGFiZTZjNSJ9.ZRp6ZSA5juu3yT_ADu2P3w8iS_tMJ1_OcXxYU0VSzJo"
	StoreAccessToken(token, redisClient)
}

func StoreAccessToken(token string, redisClient *redis.Client) {
	t, err := jwtauth.ParseToken(token, []byte("foobarfoo"))
	if err != nil {
		fmt.Println(err)
		return
	}

	key := t.ID

	// // Store the token in the Sorted Set with expiration time
	// expirationTime := time.Now().Add(time.Minute)
	// score := float64(expirationTime.Unix())
	// err = redisClient.ZAdd(context.Background(), key, &redis.Z{
	// 	Score:  score,
	// 	Member: token,
	// }).Err()

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	fmt.Println("..................................")
	// val, err := redisClient.Get(context.Background(), key).Result()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// val := redisClient.LRange(context.Background(), key, 0, 1)
	// val, err := redisClient.GetRange(context.Background(), key, 0, -1).Result()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// redisClient.Set(context.Background(), key, token, time.Second*12)

	// // time.Sleep(time.Second * 2)
	// val, err := redisClient.Get(context.Background(), key).Result()
	// if err != nil {
	// 	fmt.Println(err)
	// 	if err == redis.Nil {
	// 		// error handling attributed to key not found (probably client error)
	// 	}
	// 	// error handling attributed to internal server error
	// }

	// fmt.Println(val)

	// REDIS HSET
	// redisClient.HSet(context.Background(), key, token, time.Second * 2)
	redisClient.HSet(context.Background(), key, redisClient.HSet(context.Background(), key, token, time.Second))
	if err != nil {
		fmt.Println("error in setting hset | err: ", err)
		return
	}

	keys, err := redisClient.Keys(context.Background(), "*").Result()
	if err != nil {
		fmt.Println("error in getting keys : ", err)
		return
	}
	// redisClient.Del(context.Background(), keys...)

	fmt.Println(keys)
	fmt.Println(len(keys))

	time.Sleep(time.Second * 3)

	hval := redisClient.HGet(context.Background(), key, token).Val()
	if err != nil {
		fmt.Println("error in hgetall | err: ", err)
		return
	}

	if hval == "" {
		fmt.Println("there is no key....")
		return
	}

	fmt.Println("hval is : ", hval)

}

/*

check only for refresh token
 - logout


*/

func StoreTokenSession(token string) {
	
}

func StoreTokenSessionMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context){
		
	}
}

