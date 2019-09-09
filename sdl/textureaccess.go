package sdl

// #include <SDL2/SDL_render.h>
import "C"

type TextureAccess uint

const (
	TextureAccessStatic    TextureAccess = C.SDL_TEXTUREACCESS_STATIC    // changes rarely, not lockable
	TextureAccessStreaming TextureAccess = C.SDL_TEXTUREACCESS_STREAMING // changes frequently, lockable
	TextureAccessTarget    TextureAccess = C.SDL_TEXTUREACCESS_TARGET    // can be used as a render target
)
