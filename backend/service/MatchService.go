package service

import (
	"backend/model"
	"backend/repository"
)

type MatchService struct {
	repository repository.MatchRepository
}

func (ms *MatchService) GetAllMatches() []model.Match {
	return ms.repository.GetAllMatches()
}

func (ms *MatchService) GetMatch(id string) model.Match {
	return ms.repository.GetMatch(id)
}

func (ms *MatchService) CreateMatch(match model.Match) model.Match {
	if match.StartingPlayers == nil {
		match.StartingPlayers = make([]string, 0, 0)
	}
	if match.SubstitutePlayer == nil {
		match.SubstitutePlayer = make([]string, 0, 0)
	}
	return ms.repository.CreateMatch(match)
}

func (ms *MatchService) UpdateMatch(id string, updates model.Match) model.Match {
	currentMatch := ms.repository.GetMatch(id)
	currentMatch = setNewValues(currentMatch, updates)

	return ms.repository.UpdateMatch(currentMatch)
}

func (ms *MatchService) DeleteMatch(id string) {
	ms.repository.DeleteMatch(id)
}

func setNewValues(current model.Match, updates model.Match) model.Match {
	if updates.Description != "" && (updates.Description != current.Description) {
		current.Description = updates.Description
	}

	if updates.Finished != current.Finished {
		current.Finished = updates.Finished
	}

	if updates.Time != "" && (updates.Time != current.Time) {
		current.Time = updates.Time
	}

	if updates.Place != "" && (updates.Place != current.Place) {
		current.Place = updates.Place
	}

	if updates.Place != "" && (updates.Place != current.Place) {
		current.Place = updates.Place
	}

	if updates.Format != 0 && (updates.Format != current.Format) {
		current.Format = updates.Format
	}

	if updates.MaxPlayers != 0 && (updates.MaxPlayers != current.MaxPlayers) {
		current.MaxPlayers = updates.MaxPlayers
	}

	if updates.StartingPlayers != nil && (len(updates.StartingPlayers) != len(current.StartingPlayers)) {
		current.StartingPlayers = updates.StartingPlayers
	}

	if updates.SubstitutePlayer != nil && (len(updates.SubstitutePlayer) != len(current.SubstitutePlayer)) {
		current.SubstitutePlayer = updates.SubstitutePlayer
	}

	return current
}

func NewMatchService() MatchService {
	return MatchService{}
}
