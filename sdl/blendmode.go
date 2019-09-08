package sdl

// #include <SDL2/SDL_render.h>
import "C"

type BlendMode uint

const (
	BlendModeNone  BlendMode = C.SDL_BLENDMODE_NONE  // no blending dstRGBA = srcRGBA
	BlendModeBlend BlendMode = C.SDL_BLENDMODE_BLEND // alpha blending dstRGB = (srcRGB * srcA) + (dstRGB * (1-srcA)) dstA = srcA + (dstA * (1-srcA))
	BlendModeAdd   BlendMode = C.SDL_BLENDMODE_ADD   // additive blending dstRGB = (srcRGB * srcA) + dstRGB dstA = dstA
	BlendModeMod   BlendMode = C.SDL_BLENDMODE_MOD   // color modulate dstRGB = srcRGB * dstRGB
)
