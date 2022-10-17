package repository

import (
	"backend/model"
	"github.com/google/uuid"
)

type MatchRepository struct {
	matches []model.Match
}

func (mr *MatchRepository) GetAllMatches() []model.Match {
	if mr.matches == nil {
		return make([]model.Match, 0, 0)
	}
	return mr.matches
}

func (mr *MatchRepository) GetMatch(id string) model.Match {
	if mr.matches != nil {
		for _, match := range mr.matches {
			if match.Id == id {
				return match
			}
		}
	}
	return model.Match{}
}

func (mr *MatchRepository) CreateMatch(match model.Match) model.Match {
	if mr.matches == nil {
		mr.matches = make([]model.Match, 0, 0)
	}

	match.Id = uuid.New().String()
	mr.matches = append(mr.matches, match)
	return match
}

func (mr *MatchRepository) UpdateMatch(match model.Match) model.Match {
	mr.DeleteMatch(match.Id)
	mr.matches = append(mr.matches, match)
	return match
}

func (mr *MatchRepository) DeleteMatch(id string) {
	index := getIndexOfElement(*mr, id)
	mr.matches = append(mr.matches[:index], mr.matches[index+1:]...)
}

func getIndexOfElement(mr MatchRepository, id string) int {
	for index, match := range mr.matches {
		if match.Id == id {
			return index
		}
	}
	return -1
}
