package metrics

import (
	"backend/database/redis"
	"backend/model"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type MetricsService struct {
}

func (ms *MetricsService) AddMatch(match model.Match) {
	ctx := context.Background()
	rdb := redis.GetConnection()

	err := rdb.Set(ctx, "match-"+time.Now().String(), getMarshaled(match), 0).Err()
	if err != nil {
		panic(err)
	}
}

func (ms *MetricsService) AddPlayer(player model.Player) {
	ctx := context.Background()
	rdb := redis.GetConnection()

	err := rdb.Set(ctx, "player-"+time.Now().String(), getMarshaled(player), 0).Err()
	if err != nil {
		panic(err)
	}
}

func (ms *MetricsService) GetLastCreatedMatches() []model.Match {
	ctx := context.Background()
	rdb := redis.GetConnection()

	matches := GetMatchesFromRedis(rdb, ctx)

	return matches
}

func (ms *MetricsService) GetLastJoinedPlayers() []model.Player {
	ctx := context.Background()
	rdb := redis.GetConnection()

	players := GetPlayersFromRedis(rdb, ctx)

	return players
}

func getMarshaled(obj any) any {
	marshaled, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return marshaled
}

func getPatterns() (string, string, string) {
	currentTime := time.Now()
	oneHsBefore := currentTime.Add(-1 * time.Hour)
	twoHsBefore := currentTime.Add(-2 * time.Hour)

	pattern1 := getPattern(currentTime)
	pattern2 := getPattern(oneHsBefore)
	pattern3 := getPattern(twoHsBefore)

	return pattern1, pattern2, pattern3
}

func getPattern(time time.Time) string {
	hour := time.Hour()
	strHour := fmt.Sprintf("%d", hour)

	if hour < 10 {
		strHour = "0" + strHour
	}
	return fmt.Sprintf("%d-%d-%d %s:*", time.Year(), time.Month(), time.Day(), strHour)
}

func NewMetricsService() MetricsService {
	return MetricsService{}
}
