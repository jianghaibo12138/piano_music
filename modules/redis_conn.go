package modules

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var Pool redis.Pool

type RedisConnection struct {
	conn redis.Conn
}

var RedisHost = "127.0.0.1:6379"
var RedisPwd = "bitorobotics"

func init() {
	Pool = redis.Pool{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", RedisHost)
		},
	}
}

func GetRedisConnect() RedisConnection {
	conn := Pool.Get()
	err := conn.Send("auth", RedisPwd)
	if err != nil {
		fmt.Println(err)
	}
	return RedisConnection{conn}
}

func (c *RedisConnection) SetDb(db int) error {
	_, err := c.conn.Do("SELECT", db)
	return err
}

func (c *RedisConnection) Set(key string, value string) error {
	_, err := c.conn.Do("SET", key, value)
	return err
}
