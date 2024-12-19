package libs

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func Redis() *redis.Client {
	godotenv.Load()
	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       redisDb,
	})
	return rdb
}

func GetFromRedis(uri string) *redis.StringCmd {
	get := Redis().Get(context.Background(), uri)

	return get
}

func SetToRedis(uri string, encoded []byte) *redis.StatusCmd {
	set := Redis().Set(context.Background(), uri, string(encoded), 0)

	return set
}
