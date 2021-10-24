package go_codenames

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	bg "github.com/quibbble/go-boardgame"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
)

const (
	minTeams = 2
	maxTeams = 2

	wordCount = 25
)

type Codenames struct {
	state   *state
	actions []*bg.BoardGameAction
}

func NewCodenames(options bg.BoardGameOptions) (*Codenames, error) {
	if len(options.Teams) < minTeams {
		return nil, &bgerr.Error{
			Err:    fmt.Errorf("at least %d teams required to create a game of %s", minTeams, key),
			Status: bgerr.StatusTooFewTeams,
		}
	} else if len(options.Teams) > maxTeams {
		return nil, &bgerr.Error{
			Err:    fmt.Errorf("at most %d teams allowed to create a game of %s", maxTeams, key),
			Status: bgerr.StatusTooManyTeams,
		}
	}
	var details CodenamesOptionDetails
	if err := mapstructure.Decode(options.MoreOptions, &details); err != nil {
		return nil, &bgerr.Error{
			Err:    err,
			Status: bgerr.StatusInvalidOption,
		}
	}
	words := details.Words
	if len(details.Words) == 0 {
		words = generateWords(wordCount)
	}
	if len(words) != wordCount {
		return nil, &bgerr.Error{
			Err:    fmt.Errorf("must provide %d words or leave 'words' empty to create a game of %s", wordCount, key),
			Status: bgerr.StatusInvalidOption,
		}
	}
	return &Codenames{
		state:   newState(options.Teams, words),
		actions: make([]*bg.BoardGameAction, 0),
	}, nil
}

func (c *Codenames) Do(action bg.BoardGameAction) error {
	switch action.ActionType {
	case FlipCard:
		var details FlipCardActionDetails
		if err := mapstructure.Decode(action.MoreDetails, &details); err != nil {
			return &bgerr.Error{
				Err:    err,
				Status: bgerr.StatusInvalidActionDetails,
			}
		}
		if err := c.state.FlipCard(action.Team, details.Row, details.Column); err != nil {
			return err
		}
		c.actions = append(c.actions, &action)
	case EndTurn:
		if err := c.state.EndTurn(action.Team); err != nil {
			return err
		}
		c.actions = append(c.actions, &action)
	case Reset:
		c.state = newState(c.state.teams, generateWords(wordCount))
		c.actions = make([]*bg.BoardGameAction, 0)
	default:
		return &bgerr.Error{
			Err:    fmt.Errorf("cannot process action type %s", action.ActionType),
			Status: bgerr.StatusUnknownActionType,
		}
	}
	return nil
}

func (c *Codenames) GetSnapshot(team ...string) (*bg.BoardGameSnapshot, error) {
	if len(team) > 1 {
		return nil, &bgerr.Error{
			Err:    fmt.Errorf("get snapshot requires zero or one team"),
			Status: bgerr.StatusTooManyTeams,
		}
	}
	return &bg.BoardGameSnapshot{
		Turn:    c.state.turn,
		Teams:   c.state.teams,
		Winners: c.state.winners,
		MoreData: CodenamesSnapshotDetails{
			Board: c.state.board.board,
		},
		Actions: c.actions,
	}, nil
}
