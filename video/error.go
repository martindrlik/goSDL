package video

import (
	"errors"

	"github.com/martindrlik/goSDL/sdl"
)

func makeError(alt error) error {
	em := sdl.Error()
	if em == "" {
		return alt
	}
	return errors.New(em)
}
