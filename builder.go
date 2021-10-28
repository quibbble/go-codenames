package go_codenames

import (
	bg "github.com/quibbble/go-boardgame"
	"time"
)

const key = "codenames"

type Builder struct{}

func (b *Builder) Create(options bg.BoardGameOptions, seed ...int64) (bg.BoardGame, error) {
	if len(seed) > 0 {
		return NewCodenames(options, seed[0])
	}
	return NewCodenames(options, time.Now().Unix())
}

func (b *Builder) Key() string {
	return key
}
