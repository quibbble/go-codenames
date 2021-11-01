package go_codenames

import (
	"fmt"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
	"strconv"
	"strings"
)

// Notation - "'number of teams':'seed':'MoreOptions':'team index','action type number','details','details';..."

var (
	notationActionToInt = map[string]int{ActionFlipCard: 0, ActionEndTurn: 1}
	notationIntToAction = map[string]string{"0": ActionFlipCard, "1": ActionEndTurn}
)

func (c *CodenamesOptionDetails) encode() string {
	notation := ""
	for _, word := range c.Words {
		notation += fmt.Sprintf("%s,", word)
	}
	if len(notation) > 0 {
		return notation[:len(notation)-1]
	}
	return notation
}

func decodeNotationCodenamesOptionDetails(notation string) *CodenamesOptionDetails {
	return &CodenamesOptionDetails{
		Words: strings.Split(notation, ","),
	}
}

func (f *FlipCardActionDetails) encode() string {
	return fmt.Sprintf("%d,%d", f.Row, f.Column)
}

func decodeNotationFlipCardActionDetails(notation string) (*FlipCardActionDetails, error) {
	split := strings.Split(notation, ",")
	row, err := strconv.Atoi(split[0])
	if err != nil {
		return nil, loadFailure(err)
	}
	column, err := strconv.Atoi(split[1])
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
