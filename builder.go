package go_codenames

import (
	bg "github.com/quibbble/go-boardgame"
)

const key = "codenames"

type Builder struct{}

func (b *Builder) Create(options bg.BoardGameOptions) (bg.BoardGame, error) {
	return NewCodenames(options)
}

func (b *Builder) Key() string {
	return key
}
