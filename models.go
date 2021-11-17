package go_codenames

// Action types
const (
	ActionFlipCard = "FlipCard"
	ActionEndTurn  = "EndTurn"
)

// CodenamesMoreOptions are the additional options for creating a game of Codenames
type CodenamesMoreOptions struct {
	// Words must be 25 words in length - optional, one of Words or Seed required
	Words []string

	// Seed used to generate deterministic randomness - optional, one of Words or Seed required
	Seed int64
}

// FlipCardActionDetails is the action details for flipping a card at the desired row and column
type FlipCardActionDetails struct {
	Row, Column int
}

// CodenamesSnapshotData is the game data unique to Codenames
type CodenamesSnapshotData struct {
	Board [][]*card
}
