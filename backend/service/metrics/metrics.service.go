package metrics

import (
	"backend/database/redis"
	"backend/model"
	"context"
	"encoding/json"
	"time"
)

func AddMatch(match model.Match) {
	ctx := context.Background()
	rdb := redis.GetConnection()

	err := rdb.Set(ctx, "match-"+time.Now().String(), getMarshaled(match), 0).Err()
	if err != nil {
		panic(err)
	}
}

func AddPlayer(player model.Player) {
	ctx := context.Background()
	rdb := redis.GetConnection()

	err := rdb.Set(ctx, "player-"+time.Now().String(), getMarshaled(player), 0).Err()
	if err != nil {
		panic(err)
	}
}

func getMarshaled(obj any) any {
	marshaled, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return marshaled
}
