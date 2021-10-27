package redis

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"go-hexagonal-auth/config"

	"github.com/go-redis/redis/v8"
)

type redisServer struct {
	cfg *config.RedisServer
}

type RedisServerInterface interface {
	Connect(ctx context.Context) (*redis.Client, error)
}

func NewRedisServer(cfg *config.RedisServer) RedisServerInterface {
	return &redisServer{
		cfg: cfg,
	}
}

func (r *redisServer) Connect(ctx context.Context) (*redis.Client, error) {
	timeout := time.Duration(r.cfg.Timeout) * time.Second
	rdb := redis.NewClient(&redis.Options{
		Addr:        r.cfg.Addr,
		Password:    r.cfg.Password,
		DialTimeout: timeout,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Error("cannot connect to redis")
		return nil, err
	}
	log.Infof("success connect to redis %s", rdb)
	return rdb, nil
}
