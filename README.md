# Go-codenames

Go-codenames is a [Go](https://golang.org) implementation of the board game [Codenames](https://boardgamegeek.com/boardgame/178900/codenames). Please note that this repo only includes game logic and a basic API to interact with the game but does NOT include any form of GUI.

Check out [quibbble.com](https://quibbble.com/codenames) if you wish to view and play a live version of this game which utilizes this project along with a separate custom UI.

## Usage

To play a game create a new Codenames instance:
```go
game, err := NewCodenames(bg.BoardGameOptions{
    Teams: []string{"TeamA", "TeamB"} // must contain exactly 2 teams
})
```

To flip a card do the following action:
```go
err := game.Do(bg.BoardGameAction{
    Team: "TeamA",
    ActionType: "FlipCard",
    MoreDetails: FlipCardActionDetails{
        Row: 0,    // rows 0 to 4 are valid
        Column: 0, // columns 0 to 4 are valid
    },
})
```

To end your turn do the following action:
```go
err := game.Do(bg.BoardGameAction{
    Team: "TeamA",
    ActionType: "EndTurn",
})
```

To get the current state of the game call the following:
```go
snapshot, err := game.GetSnapshot("TeamA")
```