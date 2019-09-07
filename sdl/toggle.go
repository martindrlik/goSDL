package sdl

// #include <SDL2/SDL.h>
import "C"

type Toggle int

const (
	ToggleDisable Toggle = C.SDL_DISABLE
	ToggleEnable         = C.SDL_ENABLE
	ToggleQuery          = C.SDL_QUERY
)
