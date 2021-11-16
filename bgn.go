package go_codenames

import (
	"fmt"
	bg "github.com/quibbble/go-boardgame"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
	"strconv"
	"strings"
)

var (
	actionToNotation = map[string]string{ActionFlipCard: "f", ActionEndTurn: "e", bg.ActionSetWinners: "w"}
	notationToAction = reverseMap(actionToNotation)
)

func (c *CodenamesOptionDetails) encodeBGN() string {
	return strings.Join(c.Words, ", ")
}

func decodeCodenamesOptionDetailsBGN(notation string) *CodenamesOptionDetails {
	return &CodenamesOptionDetails{
		Words: strings.Split(notation, ", "),
	}
}

func (f *FlipCardActionDetails) encodeBGN() []string {
	return []string{strconv.Itoa(f.Row), strconv.Itoa(f.Column)}
}

func decodeFlipCardActionDetailsBGN(notation []string) (*FlipCardActionDetails, error) {
	if len(notation) != 2 {
		return nil, loadFailure(fmt.Errorf("invalid flip card notation"))
	}
	row, err := strconv.Atoi(notation[0])
	if err != nil {
		return nil, loadFailure(err)
	}
	column, err := strconv.Atoi(notation[1])
	if err != nil {
		return nil, loadFailure(err)
	}
	return &FlipCardActionDetails{
		Row:    row,
		Column: column,
	}, nil
}

func loadFailure(err error) error {
	return &bgerr.Error{
		Err:    err,
		Status: bgerr.StatusBGNDecodingFailure,
	}
}
