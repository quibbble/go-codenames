package go_codenames

// Action types
const (
	ActionFlipCard = "FlipCard"
	ActionEndTurn  = "EndTurn"
)

// CodenamesOptionDetails allows custom words to be used in the game
type CodenamesOptionDetails struct {
	Words []string // must be 25 words in length
}

// FlipCardActionDetails is the action details for flipping a card at the desired row and column
type FlipCardActionDetails struct {
	Row, Column int
}

// CodenamesSnapshotData is the game data unique to Codenames
type CodenamesSnapshotData struct {
	Board [][]*card
}
