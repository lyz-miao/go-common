package redis

import (
    "context"
    "github.com/go-redis/redis/v8"
)

type Config struct {
    DSN string
}

func New(c *Config) *redis.Client {
    o, err := redis.ParseURL(c.DSN)
    if err != nil {
        panic(err)
    }
    r := redis.NewClient(o)

    if err = r.Ping(context.Background()).Err(); err != nil {
        panic(err)
    } else {
        return r
    }
}
