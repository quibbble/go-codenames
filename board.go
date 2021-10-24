package go_codenames

import (
	wr "github.com/mroth/weightedrand"
)

const (
	rows    = 5
	columns = 5

	none  = "none"
	death = "death"
)

type board struct {
	board [][]*card
}

func newBoard(teams, words []string) *board {
	keys := []string{teams[0], teams[1], none, death}
	remaining := map[string]int{keys[0]: 9, keys[1]: 8, keys[2]: 7, keys[3]: 1}
	var b = make([][]*card, rows)
	count := 0
	for row := 0; row < rows; row++ {
		b[row] = make([]*card, columns)
		for col := 0; col < columns; col++ {
			word := words[count]
			chooser, _ := wr.NewChooser(
				wr.Choice{Item: keys[0], Weight: uint(remaining[keys[0]])},
				wr.Choice{Item: keys[1], Weight: uint(remaining[keys[1]])},
				wr.Choice{Item: keys[2], Weight: uint(remaining[keys[2]])},
				wr.Choice{Item: keys[3], Weight: uint(remaining[keys[3]])},
			)
			choice := chooser.Pick().(string)
			remaining[choice]--
			c := &card{
				Word:    word,
				Team:    choice,
				Flipped: false,
			}
			b[row][col] = c
			count++
		}
	}
	return &board{
		board: b,
	}
}

func (b *board) remaining(team string) int {
	count := 0
	for _, row := range b.board {
		for _, card := range row {
			if !card.Flipped && card.Team == team {
				count++
			}
		}
	}
	return count
}
