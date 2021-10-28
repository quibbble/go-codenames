package go_codenames

import (
	bg "github.com/quibbble/go-boardgame"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	TeamA = "TeamA"
	TeamB = "TeamB"
)

func Test_Codenames(t *testing.T) {
	codenames, err := NewCodenames(bg.BoardGameOptions{
		Teams: []string{TeamA, TeamB},
	}, time.Now().UnixNano())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	codenames.state.turn = TeamA

	// flip card at 0,0
	err = codenames.Do(bg.BoardGameAction{
		Team:       TeamA,
		ActionType: ActionFlipCard,
		MoreDetails: FlipCardActionDetails{
			Row:    0,
			Column: 0,
		},
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Equal(t, true, codenames.state.board.board[0][0].Flipped)
}
