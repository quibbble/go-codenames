package go_codenames

// Action types
const (
	FlipCard = "FlipCard"
	EndTurn  = "EndTurn"
	Reset    = "Reset"
)

// CodenamesOptionDetails allows custom words to be used in the game
type CodenamesOptionDetails struct {
	Words []string // must be 25 words in length
}

// FlipCardActionDetails is the action details for flipping a card at the desired row and column
type FlipCardActionDetails struct {
	Row, Column int
}

// CodenamesSnapshotDetails are the details unique to codenames
type CodenamesSnapshotDetails struct {
	Board [][]*card
}
