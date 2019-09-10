package sdl

// #include <SDL2/SDL_keycode.h>
import "C"

// Key contains key information used in key events.
type Key struct {
	Physical C.SDL_Scancode
	Virtual  C.SDL_Keycode
	Mod      uint16
	_        uint32
}
