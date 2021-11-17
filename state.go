package go_codenames

import (
	"fmt"
	bg "github.com/quibbble/go-boardgame"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
)

type state struct {
	turn    string
	teams   []string
	winners []string
	board   *board
}

func newState(teams []string, words []string) *state {
	return &state{
		turn:    teams[0],
		teams:   teams,
		winners: make([]string, 0),
		board:   newBoard(teams, words),
	}
}

func (s *state) FlipCard(team string, row, column int) error {
	if team != s.turn {
		return &bgerr.Error{
			Err:    fmt.Errorf("%s cannot play on %s turn", team, s.turn),
			Status: bgerr.StatusWrongTurn,
		}
	}
	if row < 0 || row > rows || column < 0 || column > columns {
		return &bgerr.Error{
			Err:    fmt.Errorf("invalid location for row %d and col %d", row, column),
			Status: bgerr.StatusInvalidActionDetails,
		}
	}
	if s.board.board[row][column].Flipped {
		return &bgerr.Error{
			Err:    fmt.Errorf("card aready flipped at row %d and col %d", row, column),
			Status: bgerr.StatusInvalidAction,
		}
	}
	s.board.board[row][column].Flipped = true
	switch s.board.board[row][column].Team {
	case death:
		s.winners = []string{s.otherTeam(s.turn)}
	case s.turn:
		if s.board.remaining(s.turn) == 0 {
			s.winners = []string{s.turn}
		}
	default:
		s.turn = s.otherTeam(s.turn)
	}
	return nil
}

func (s *state) EndTurn(team string) error {
	if team != s.turn {
		return &bgerr.Error{
			Err:    fmt.Errorf("%s cannot play on %s turn", team, s.turn),
			Status: bgerr.StatusWrongTurn,
		}
	}
	s.turn = s.otherTeam(s.turn)
	return nil
}

func (s *state) SetWinners(winners []string) error {
	for _, winner := range winners {
		if !contains(s.teams, winner) {
			return &bgerr.Error{
				Err:    fmt.Errorf("winner not in teams"),
				Status: bgerr.StatusInvalidActionDetails,
			}
		}
	}
	s.winners = winners
	return nil
}

func (s *state) targets() []*bg.BoardGameAction {
	targets := make([]*bg.BoardGameAction, 0)
	targets = append(targets, &bg.BoardGameAction{
		Team:       s.turn,
		ActionType: ActionEndTurn,
	})
	for r, row := range s.board.board {
		for c, card := range row {
			if !card.Flipped {
				targets = append(targets, &bg.BoardGameAction{
					Team:       s.turn,
					ActionType: ActionFlipCard,
					MoreDetails: FlipCardActionDetails{
						Row:    r,
						Column: c,
					},
				})
			}
		}
	}
	return targets
}

func (s *state) otherTeam(team string) string {
	if s.teams[0] == team {
		return s.teams[1]
	}
	return s.teams[0]
}
