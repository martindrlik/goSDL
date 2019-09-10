package sdl_test

import (
	"os"
	"testing"

	"github.com/martindrlik/goSDL/sdl"
)

func TestMain(m *testing.M) {
	sdl.Init(sdl.Everything)
	defer sdl.Quit()
	os.Exit(m.Run())
}
