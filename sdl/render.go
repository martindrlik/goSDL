package sdl

// #include <SDL2/SDL_render.h>
import "C"

import (
	"errors"
	"unsafe"
)

// Renderer contains a rendering state.
type Renderer C.SDL_Renderer

func (renderer *Renderer) cptr() *C.SDL_Renderer {
	return (*C.SDL_Renderer)(unsafe.Pointer(renderer))
}

// ComposeCustomBlendMode composes a custom blend mode for renderers.
func ComposeCustomBlendMode(
	srcColorFactor, dstColorFactor BlendFactor, colorOperation BlendOperation,
	srcAlphaFactor, dstAlphaFactor BlendFactor, alphaOperation BlendOperation) BlendMode {
	cm := C.SDL_ComposeCustomBlendMode(
		C.SDL_BlendFactor(srcColorFactor),
		C.SDL_BlendFactor(dstColorFactor),
		C.SDL_BlendOperation(colorOperation),
		C.SDL_BlendFactor(srcAlphaFactor),
		C.SDL_BlendFactor(dstAlphaFactor),
		C.SDL_BlendOperation(alphaOperation))
	return BlendMode(cm)
}

// CreateRenderer creates a 2D rendering context for a window.
func (w *Window) CreateRenderer(index int, flags uint) (*Renderer, error) {
	cr := C.SDL_CreateRenderer(w.cptr(), C.int(index), C.uint(flags))
	if cr == nil {
		return nil, errors.New(Error())
	}
	return (*Renderer)(unsafe.Pointer(cr)), nil
}

// SDL_CreateSoftwareRenderer
// SDL_CreateTexture
// SDL_CreateTextureFromSurface
// SDL_CreateWindowAndRenderer

// Destroy destroys the rendering context for
// a window and free associated textures.
func (renderer *Renderer) Destroy() {
	C.SDL_DestroyRenderer(renderer.cptr())
}

// SDL_DestroyTexture
// SDL_GL_BindTexture
// SDL_GL_UnbindTexture
// SDL_GetNumRenderDrivers
// SDL_GetRenderDrawBlendMode
// SDL_GetRenderDrawColor
// SDL_GetRenderDriverInfo
// SDL_GetRenderTarget
// SDL_GetRenderer
// SDL_GetRendererInfo
// SDL_GetRendererOutputSize
// SDL_GetTextureAlphaMod
// SDL_GetTextureBlendMode
// SDL_GetTextureColorMod
// SDL_LockTexture
// SDL_QueryTexture

// RenderClear clears the current rendering
// target with the drawing color.
func (renderer *Renderer) RenderClear() error {
	if C.SDL_RenderClear(renderer.cptr()) < 0 {
		return errors.New(Error())
	}
	return nil
}

// SDL_RenderCopy
// SDL_RenderCopyEx

// RenderDrawLine draws a line on the current rendering target.
func (renderer *Renderer) RenderDrawLine(x1, y1, x2, y2 int) error {
	if C.SDL_RenderDrawLine(
		renderer.cptr(),
		C.int(x1),
		C.int(y1),
		C.int(x2),
		C.int(y2)) < 0 {
		return errors.New(Error())
	}
	return nil
}

// SDL_RenderDrawLines

// RenderDrawPoint draws a point on the current rendering target.
func (renderer *Renderer) RenderDrawPoint(x, y int) error {
	if C.SDL_RenderDrawPoint(
		renderer.cptr(),
		C.int(x),
		C.int(y)) < 0 {
		return errors.New(Error())
	}
	return nil
}

// RenderDrawPoints draws multiple points on the current rendering target.
func (renderer *Renderer) RenderDrawPoints(points []struct{ X, Y int32 }) error {
	if C.SDL_RenderDrawPoints(
		renderer.cptr(),
		(*C.SDL_Point)(unsafe.Pointer(&points[0])),
		C.int(len(points))) < 0 {
		return errors.New(Error())
	}
	return nil
}

// SDL_RenderDrawRect
// SDL_RenderDrawRects
// SDL_RenderFillRect
// SDL_RenderFillRects
// SDL_RenderGetClipRect
// SDL_RenderGetIntegerScale
// SDL_RenderGetLogicalSize
// SDL_RenderGetScale
// SDL_RenderGetViewport
// SDL_RenderIsClipEnabled

// RenderPresent updates the screen with any rendering
// performed since the previous call.
func (renderer *Renderer) RenderPresent() {
	C.SDL_RenderPresent(renderer.cptr())
}

// SDL_RenderReadPixels
// SDL_RenderSetClipRect
// SDL_RenderSetIntegerScale
// SDL_RenderSetLogicalSize
// SDL_RenderSetScale
// SDL_RenderSetViewport
// SDL_RenderTargetSupported
// SDL_SetRenderDrawBlendMode

// SetRenderDrawColor sets the color used for drawing
// operations (Rect, Line and Clear).
func (renderer *Renderer) SetRenderDrawColor(r, g, b, a uint8) error {
	if C.SDL_SetRenderDrawColor(
		renderer.cptr(),
		C.Uint8(r),
		C.Uint8(g),
		C.Uint8(b),
		C.Uint8(a)) < 0 {
		return errors.New(Error())
	}
	return nil
}

// SDL_SetRenderTarget
// SDL_SetTextureAlphaMod
// SDL_SetTextureBlendMode
// SDL_SetTextureColorMod
// SDL_UnlockTexture
// SDL_UpdateTexture
// SDL_UpdateYUVTexture
