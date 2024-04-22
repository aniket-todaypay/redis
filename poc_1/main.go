package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func main() {
	fmt.Println("Go Redis Started......")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("name", "Aniket", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiTUVSQ0hBTlQiLCJpZCI6ImFjY19xaVVrQlFydnYzaFEiLCJpc19wYXJ0aWFsIjpmYWxzZSwidXBsb2FkX2NhdGVnb3J5IjoiIiwiaXNzIjoiMmRwIiwic3ViIjoiTUVSQ0hBTlQiLCJleHAiOjE3NjAwMDIzODQsIm5iZiI6MTY5NjkzMDM4NCwiaWF0IjoxNjk2OTMwMzg0LCJqdGkiOiJmNmU5MmFjMS02MTMwLTQ3ZDktYjliZi0xNTJlZGYyY2Q3NmQifQ.X-o-iRUm5wDDGk-z3lqllp_nqPyOvjb5JZmc0x7qMhk", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
	InvalidateUser("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiTUVSQ0hBTlQiLCJpZCI6ImFjY19xaVVrQlFydnYzaFEiLCJpc19wYXJ0aWFsIjpmYWxzZSwidXBsb2FkX2NhdGVnb3J5IjoiIiwiaXNzIjoiMmRwIiwic3ViIjoiTUVSQ0hBTlQiLCJleHAiOjE3NjAwMDIzODQsIm5iZiI6MTY5NjkzMDM4NCwiaWF0IjoxNjk2OTMwMzg0LCJqdGkiOiJmNmU5MmFjMS02MTMwLTQ3ZDktYjliZi0xNTJlZGYyY2Q3NmQifQ.X-o-iRUm5wDDGk-z3lqllp_nqPyOvjb5JZmc0x7qMhk", "")
}

func InvalidateUser(token, refreshToken string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	token, err := client.Get("token").Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(token)

}
