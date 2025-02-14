package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var eventKey string = "events"

func GetAllEvents(ctx context.Context, rdb *redis.Client) (hit bool, err error) {
	return false, nil
}

func GetEvent(ectx context.Context, rdb *redis.Client, eventId string) (event []EventResp, hit bool, err error) {
	return nil, false, nil
}

func SetEvent(ctx context.Context, rdb *redis.Client, eventId string) error {
	return nil
}

func EditEvent(ctx context.Context, rdb *redis.Client, eventId string) error {

	return nil
}

func DeleteEvent(ctx context.Context, rdb *redis.Client, eventId string) error {
	return nil
}
