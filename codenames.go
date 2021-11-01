package go_codenames

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	bg "github.com/quibbble/go-boardgame"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
	"math/rand"
	"time"
)

const (
	minTeams = 2
	maxTeams = 2

	wordCount = 25
)

type Codenames struct {
	state   *state
	actions []*bg.BoardGameAction
	seed    int64
	options *CodenamesOptionDetails
}

func NewCodenames(options *bg.BoardGameOptions) (*Codenames, error) {
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
		words = generateWords(wordCount, rand.New(rand.NewSource(options.Seed)))
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
		seed:    options.Seed,
		options: &details,
	}, nil
}

func (c *Codenames) Do(action *bg.BoardGameAction) error {
	switch action.ActionType {
	case ActionFlipCard:
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
		c.actions = append(c.actions, action)
	case ActionEndTurn:
		if err := c.state.EndTurn(action.Team); err != nil {
			return err
		}
		c.actions = append(c.actions, action)
	case bg.ActionReset:
		seed := time.Now().UnixNano()
		c.state = newState(c.state.teams, generateWords(wordCount, rand.New(rand.NewSource(seed))))
		c.actions = make([]*bg.BoardGameAction, 0)
		c.seed = seed
	case bg.ActionUndo:
		if len(c.actions) > 0 {
			undo, _ := NewCodenames(&bg.BoardGameOptions{Teams: c.state.teams, Seed: c.seed})
			for _, a := range c.actions[:len(c.actions)-1] {
				if err := undo.Do(a); err != nil {
					return err
				}
			}
			c.state = undo.state
			c.actions = undo.actions
		} else {
			return &bgerr.Error{
				Err:    fmt.Errorf("no actions to undo"),
				Status: bgerr.StatusInvalidAction,
			}
		}
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
	var targets []*bg.BoardGameAction
	if len(c.state.winners) == 0 && (len(team) == 0 || (len(team) == 1 && team[0] == c.state.turn)) {
		targets = c.state.targets()
	}
	return &bg.BoardGameSnapshot{
		Turn:    c.state.turn,
		Teams:   c.state.teams,
		Winners: c.state.winners,
		MoreData: CodenamesSnapshotDetails{
			Board: c.state.board.board,
		},
		Targets: targets,
		Actions: c.actions,
	}, nil
}

func (c *Codenames) GetNotation() string {
	// extra colon is left for MoreOptions which may be utilized in future additions
	notation := fmt.Sprintf("%d:%d:%s:", len(c.state.teams), c.seed, c.options.encode())
	for _, action := range c.actions {
		base := fmt.Sprintf("%d,%d", indexOf(c.state.teams, action.Team), notationActionToInt[action.ActionType])
		switch action.ActionType {
		case ActionFlipCard:
			var details FlipCardActionDetails
			_ = mapstructure.Decode(action.MoreDetails, &details)
			base = fmt.Sprintf("%s,%s;", base, details.encode())
		default:
			base = fmt.Sprintf("%s;", base)
		}
		notation += base
	}
	return notation
}
