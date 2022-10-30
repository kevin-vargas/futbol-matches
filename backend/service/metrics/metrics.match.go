package metrics

import (
	"backend/model"
	"context"
	"encoding/json"
	redis2 "github.com/go-redis/redis/v8"
)

func GetMatchesFromRedis(rdb *redis2.Client, ctx context.Context) []model.Match {
	var cursor uint64

	pattern1, pattern2, pattern3 := getPatterns()

	it1 := rdb.Scan(ctx, cursor, "match-"+pattern1+"*", 0).Iterator()
	it2 := rdb.Scan(ctx, cursor, "match-"+pattern2+"*", 0).Iterator()
	it3 := rdb.Scan(ctx, cursor, "match-"+pattern3+"*", 0).Iterator()

	return append(getMatchesFromIterator(it1, ctx), append(getMatchesFromIterator(it2, ctx),
		getMatchesFromIterator(it3, ctx)...)...)
}

func getMatchesFromIterator(it *redis2.ScanIterator, ctx context.Context) []model.Match {
	matches := []model.Match{}
	for it.Next(ctx) {
		objBytes := []byte(it.Val())
		match := model.Match{}
		_ = json.Unmarshal(objBytes, &match)
		matches = append(matches, match)
	}
	return matches
}
