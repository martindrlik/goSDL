package sdl

// #include <SDL2/SDL_render.h>
import "C"

const (
	RendererSoftware      = C.SDL_RENDERER_SOFTWARE      // the renderer is a software fallback
	RendererAccelerated   = C.SDL_RENDERER_ACCELERATED   // the renderer uses hardware acceleration
	RendererPresentVSync  = C.SDL_RENDERER_PRESENTVSYNC  // present is synchronized with the refresh rate
	RendererTargetTexture = C.SDL_RENDERER_TARGETTEXTURE // the renderer supports rendering to texture
)
