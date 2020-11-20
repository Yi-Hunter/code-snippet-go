package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.ClusterClient

// https://github.com/go-redis/redis
// https://godoc.org/github.com/go-redis/redis
func init() {
	client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              []string{"192.168.1.1:1234"}, // ip:port
		MaxRedirects:       0,
		ReadOnly:           false,
		RouteByLatency:     false,
		RouteRandomly:      false,
		ClusterSlots:       nil,
		OnNewNode:          nil,
		OnConnect:          nil,
		Password:           "",
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        time.Second,
		WriteTimeout:       time.Second,
		PoolSize:           100,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        time.Minute,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})

	if client == nil {
		panic("redis init failed")
		return
	}

	fmt.Println("redis init success")
}

func Set(key string, value interface{}, expireTime time.Duration) (result string, err error) {
	result, err = client.Set(key, value, expireTime).Result()
	return
}

// SetNX
// if not lock, return false
func SetNX(key string, value interface{}, expireTime time.Duration) (result bool, err error) {
	result, err = client.SetNX(key, value, expireTime).Result()
	return
}

func Get(key string) (result string, err error) {
	result, err = client.Get(key).Result()
	if err != nil {
		if err != redis.Nil {
		}
	}
	return
}

func GetInt(key string) (result int, err error) {

	result, err = client.Get(key).Int()
	if err != nil {
		if err != redis.Nil {
		}
	}
	return
}

// Incr redis value++
// result is the value after value++
// if the key not exist, return 1
func Incr(key string) (result int64, err error) {

	result, err = client.Incr(key).Result()

	return
}

// Decr redis value--
// result is the value after value--
// if the key not exist, return -1
func Decr(key string) (result int64, err error) {

	result, err = client.Decr(key).Result()

	return
}

func Del(key string) (aff int64, err error) {
	aff, err = client.Del(key).Result()
	if err != nil {
		if err != redis.Nil {
		}
	}
	return
}
