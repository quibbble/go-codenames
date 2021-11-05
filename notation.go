package go_codenames

import (
	"fmt"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
	"strconv"
	"strings"
)

var (
	actionToNotation = map[string]string{ActionFlipCard: "f", ActionEndTurn: "e"}
	notationToAction = reverseMap(actionToNotation)
)

func (c *CodenamesOptionDetails) encode() string {
	return strings.Join(c.Words, ", ")
}

func decodeCodenamesOptionDetails(notation string) *CodenamesOptionDetails {
	return &CodenamesOptionDetails{
		Words: strings.Split(notation, ", "),
	}
}

func (f *FlipCardActionDetails) encode() []string {
	return []string{strconv.Itoa(f.Row), strconv.Itoa(f.Column)}
}

func decodeFlipCardActionDetails(notation []string) (*FlipCardActionDetails, error) {
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
		Status: bgerr.StatusGameLoadFailure,
	}
}
