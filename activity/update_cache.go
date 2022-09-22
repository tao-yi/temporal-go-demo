package activity

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type UpdateCacheArgs struct {
	Amount    float32
	AccountID int
}

func UpdateCachedActivity(ctx context.Context, args UpdateCacheArgs) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	key := fmt.Sprintf("user:%d:balance", args.AccountID)
	err := rdb.IncrByFloat(ctx, key, float64(args.Amount)).Err()
	if err != nil {
		return err
	}
	return nil
}
