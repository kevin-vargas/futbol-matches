package match

import (
	"backend/model"
	mr "backend/repository/match"
)

type MatchService struct {
	repository mr.MatchRepository
}

func (ms *MatchService) GetAllMatches() []model.Match {
	return ms.repository.GetAllMatches()
}

func (ms *MatchService) GetMatch(id string) model.Match {
	return ms.repository.GetMatch(id)
}

func (ms *MatchService) CreateMatch(match model.Match) (string, error) {
	if match.StartingPlayers == nil {
		match.StartingPlayers = make([]model.Player, 0, 0)
	}
	if match.SubstitutePlayer == nil {
		match.SubstitutePlayer = make([]model.Player, 0, 0)
	}
	return ms.repository.CreateMatch(match)
}

func (ms *MatchService) UpdateMatch(id string, updates model.Match) error {
	currentMatch := ms.repository.GetMatch(id)
	currentMatch = setNewValues(currentMatch, updates)

	return ms.repository.UpdateMatch(currentMatch)
}

func (ms *MatchService) DeleteMatch(id string) error {
	return ms.repository.DeleteMatch(id)
}

func (ms *MatchService) AddPlayer(matchId string, player model.Player) bool {
	match := ms.repository.GetMatch(matchId)
	if len(match.StartingPlayers)+len(match.SubstitutePlayer) == match.MaxPlayers {
		return false
	} else {
		if len(match.StartingPlayers) == match.Format*2 {
			match.SubstitutePlayer = append(match.SubstitutePlayer, player)
		} else {
			match.StartingPlayers = append(match.StartingPlayers, player)
		}
		ms.repository.UpdateMatch(match)
		return true
	}
}

func setNewValues(current model.Match, updates model.Match) model.Match {
	if updates.Owner != "" && (updates.Owner != current.Owner) {
		current.Owner = updates.Owner
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

func NewMatchService(matchRepo mr.MatchRepository) MatchService {
	return MatchService{
		repository: matchRepo,
	}
}
