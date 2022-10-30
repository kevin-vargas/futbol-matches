package metrics

import (
	"backend/model"
	"context"
	"encoding/json"
	redis2 "github.com/go-redis/redis/v8"
)

func GetPlayersFromRedis(rdb *redis2.Client, ctx context.Context) []model.Player {
	var cursor uint64

	pattern1, pattern2, pattern3 := getPatterns()

	it1 := rdb.Scan(ctx, cursor, "player-"+pattern1+"*", 0).Iterator()
	it2 := rdb.Scan(ctx, cursor, "player-"+pattern2+"*", 0).Iterator()
	it3 := rdb.Scan(ctx, cursor, "player-"+pattern3+"*", 0).Iterator()

	return append(getPlayersFromIterator(it1, ctx), append(getPlayersFromIterator(it2, ctx),
		getPlayersFromIterator(it3, ctx)...)...)
}

func getPlayersFromIterator(it *redis2.ScanIterator, ctx context.Context) []model.Player {
	players := []model.Player{}
	for it.Next(ctx) {
		objBytes := []byte(it.Val())
		player := model.Player{}
		_ = json.Unmarshal(objBytes, &player)
		players = append(players, player)
	}
	return players
}
