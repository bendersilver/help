package keydb

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var kdb *redis.Client
var Nil = redis.Nil
var ctx = context.Background()

func Init(opt *redis.Options) {
	if opt == nil {
		opt = &redis.Options{
			Addr: "localhost:6789",
			DB:   9,
		}
	}
	kdb = redis.NewClient(opt)
}

func Cli() *redis.Client {
	return kdb
}

func HGet(key, field string) *redis.StringCmd {
	return kdb.HGet(ctx, key, field)
}

// HSetStruct -
func HSetStruct(key string, dst interface{}) error {
	sf, err := newStructFields(dst)
	if err != nil {
		return err
	}
	if len(sf.RedisArgs) == 0 {
		return fmt.Errorf("wrong number of arguments for 'hset' command")
	}
	err = kdb.HSet(ctx, key, sf.RedisArgs...).Err()
	if err != nil {
		return err
	}
	if len(sf.EmptyFields) > 0 {
		return kdb.HDel(ctx, key, sf.EmptyFields...).Err()
	}
	return nil
}

// HGetStruct -
func HGetStruct(key string, dst interface{}) error {
	mp, err := kdb.HGetAll(ctx, key).Result()
	if err != nil {
		return err
	}
	sf, err := newStructFields(dst)
	if err != nil {
		return err
	}
	return sf.fromStringStringMap(mp)
}
