package go_codenames

import (
	"fmt"
	"strconv"
	"strings"

	bg "github.com/quibbble/go-boardgame"
	"github.com/quibbble/go-boardgame/pkg/bgn"
)

const key = "Codenames"

type Builder struct{}

func (b *Builder) Create(options *bg.BoardGameOptions) (bg.BoardGame, error) {
	return NewCodenames(options)
}

func (b *Builder) CreateWithBGN(options *bg.BoardGameOptions) (bg.BoardGameWithBGN, error) {
	return NewCodenames(options)
}

func (b *Builder) Load(game *bgn.Game) (bg.BoardGameWithBGN, error) {
	if game.Tags["Game"] != key {
		return nil, loadFailure(fmt.Errorf("game tag does not match game key"))
	}
	teamsStr, ok := game.Tags["Teams"]
	if !ok {
		return nil, loadFailure(fmt.Errorf("missing teams tag"))
	}
	teams := strings.Split(teamsStr, ", ")

	var details *CodenamesMoreOptions
	if len(game.Tags["Words"]) > 0 {
		details.Words = strings.Split(game.Tags["Words"], ", ")
	} else {
		seedStr, ok := game.Tags["Seed"]
		if !ok {
			return nil, loadFailure(fmt.Errorf("missing seed or words tag"))
		}
		s, err := strconv.Atoi(seedStr)
		if err != nil {
			return nil, loadFailure(err)
		}
		details.Seed = int64(s)
	}
	g, err := b.CreateWithBGN(&bg.BoardGameOptions{
		Teams:       teams,
		MoreOptions: details,
	})
	if err != nil {
		return nil, err
	}
	for _, action := range game.Actions {
		if action.TeamIndex >= len(teams) {
			return nil, loadFailure(fmt.Errorf("team index %d out of range", action.TeamIndex))
		}
		team := teams[action.TeamIndex]
		actionType := notationToAction[string(action.ActionKey)]
		if actionType == "" {
			return nil, loadFailure(fmt.Errorf("invalid action key %s", string(action.ActionKey)))
		}
		var details interface{}
		switch actionType {
		case ActionFlipCard:
			result, err := decodeFlipCardActionDetailsBGN(action.Details)
			if err != nil {
				return nil, err
			}
			details = result
		case bg.ActionSetWinners:
			result, err := bg.DecodeSetWinnersActionDetailsBGN(action.Details, teams)
			if err != nil {
				return nil, err
			}
			details = result
		}
		if err := g.Do(&bg.BoardGameAction{
			Team:        team,
			ActionType:  actionType,
			MoreDetails: details,
		}); err != nil {
			return nil, err
		}
	}
	return g, nil
}

func (b *Builder) Key() string {
	return key
}
