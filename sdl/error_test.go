package sdl_test

import (
	"testing"

	"github.com/martindrlik/goSDL/sdl"
)

func TestError(t *testing.T) {
	sdl.SetError("some %s", "error")
	t.Run("Error", func(t *testing.T) {
		expected := "some error"
		actual := sdl.Error()
		if expected != actual {
			t.Fatalf("expected %q got %q", expected, actual)
		}
	})
	t.Run("ClearError", func(t *testing.T) {
		sdl.ClearError()
		actual := sdl.Error()
		if actual != "" {
			t.Errorf("expected \"\" after ClearError got %q", actual)
		}
	})
}
