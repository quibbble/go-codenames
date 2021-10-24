package go_codenames

import (
	"fmt"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
	"math/rand"
)

type state struct {
	turn    string
	teams   []string
	winners []string
	board   *board
}

func newState(teams []string, words []string) *state {
	return &state{
		turn:    teams[rand.Intn(len(teams))],
		teams:   teams,
		winners: make([]string, 0),
		board:   newBoard(teams, words),
	}
}

func (s *state) FlipCard(team string, row, column int) error {
	if len(s.winners) > 0 {
		return &bgerr.Error{
			Err:    fmt.Errorf("%s game already completed", key),
			Status: bgerr.StatusGameOver,
		}
	}
	if team != s.turn {
		return &bgerr.Error{
			Err:    fmt.Errorf("currently %s's turn", s.turn),
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
	if len(s.winners) > 0 {
		return &bgerr.Error{
			Err:    fmt.Errorf("%s game already completed", key),
			Status: bgerr.StatusGameOver,
		}
	}
	if team != s.turn {
		return &bgerr.Error{
			Err:    fmt.Errorf("currently %s's turn", s.turn),
			Status: bgerr.StatusWrongTurn,
		}
	}
	s.turn = s.otherTeam(s.turn)
	return nil
}

func (s *state) otherTeam(team string) string {
	if s.teams[0] == team {
		return s.teams[1]
	}
	return s.teams[0]
}
