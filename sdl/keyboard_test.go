package sdl_test

import (
	"testing"

	"github.com/martindrlik/goSDL/sdl"
)

func TestKeyFromName(t *testing.T) {
	k := sdl.KeyFromName("CapsLock")
	if k != sdl.KeyCapsLock {
		t.Errorf("expected %v got %v", sdl.KeyCapsLock, k)
	}
}

func TestKeyName(t *testing.T) {
	name := sdl.KeyName(sdl.KeyCapsLock)
	if name != "CapsLock" {
		t.Errorf("expected %q got %q", "CapsLock", name)
	}
}
