package sdl

// #include <SDL2/SDL.h>
import "C"

const (
	Disable = C.SDL_DISABLE
	Enable  = C.SDL_ENABLE
	Query   = C.SDL_QUERY
)
