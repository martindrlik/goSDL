package sdl

// #include <SDL2/SDL_render.h>
import "C"

type BlendFactor uint

const (
	BlendFactorZero             BlendFactor = C.SDL_BLENDFACTOR_ZERO                // 0, 0, 0, 0
	BlendFactorOne              BlendFactor = C.SDL_BLENDFACTOR_ONE                 // 1, 1, 1, 1
	BlendFactorSrcColor         BlendFactor = C.SDL_BLENDFACTOR_SRC_COLOR           // srcR, srcG, srcB, srcA
	BlendFactorOneMinusSrcColor BlendFactor = C.SDL_BLENDFACTOR_ONE_MINUS_SRC_COLOR // 1-srcR, 1-srcG, 1-srcB, 1-srcA
	BlendFactorSrcAlpha         BlendFactor = C.SDL_BLENDFACTOR_SRC_ALPHA           // srcA, srcA, srcA, srcA
	BlendFactorOneMinusSrcAlpha BlendFactor = C.SDL_BLENDFACTOR_ONE_MINUS_SRC_ALPHA // 1-srcA, 1-srcA, 1-srcA, 1-srcA
	BlendFactorDstColor         BlendFactor = C.SDL_BLENDFACTOR_DST_COLOR           // dstR, dstG, dstB, dstA
	BlendFactorOneMinusDstColor BlendFactor = C.SDL_BLENDFACTOR_ONE_MINUS_DST_COLOR // 1-dstR, 1-dstG, 1-dstB, 1-dstA
	BlendFactorDstAlpha         BlendFactor = C.SDL_BLENDFACTOR_DST_ALPHA           // dstA, dstA, dstA, dstA
	BlendFactorOneMinusDstAlpha BlendFactor = C.SDL_BLENDFACTOR_ONE_MINUS_DST_ALPHA // 1-dstA, 1-dstA, 1-dstA, 1-dstA
)
