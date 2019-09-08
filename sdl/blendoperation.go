package sdl

// #include <SDL2/SDL_render.h>
import "C"

type BlendOperation uint

const (
	BlendOpAdd         BlendOperation = C.SDL_BLENDOPERATION_ADD          // additive operation dst + src
	BlendOpSubtract    BlendOperation = C.SDL_BLENDOPERATION_SUBTRACT     // subtractive operation dst - src
	BlendOpRevSubtract BlendOperation = C.SDL_BLENDOPERATION_REV_SUBTRACT // reversed subtractive operation src - dst
	BlendOpMinimum     BlendOperation = C.SDL_BLENDOPERATION_MINIMUM      // minimum operation min(dst, src)
	BlendOpMaximum     BlendOperation = C.SDL_BLENDOPERATION_MAXIMUM      // maximum operation max(dst, src)
)
