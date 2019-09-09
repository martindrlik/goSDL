package sdl

// #include <SDL2/SDL_render.h>
import "C"
import (
	"errors"
	"unsafe"
)

// Texture contains an efficient, driver-specific representation of pixel data.
type Texture C.SDL_Texture

func (tx *Texture) cptr() *C.SDL_Texture {
	return (*C.SDL_Texture)(unsafe.Pointer(tx))
}

// CreateTexture creates a texture for a rendering context.
func (r *Renderer) CreateTexture(format PixelFormat, access TextureAccess, w, h int) (*Texture, error) {
	ct := C.SDL_CreateTexture(r.cptr(), C.uint(format), C.int(access), C.int(w), C.int(h))
	if ct == nil {
		return nil, errors.New(Error())
	}
	return (*Texture)(unsafe.Pointer(ct)), nil
}

// CreateTextureFromSurface creates a texture from an existing surface.
func (r *Renderer) CreateTextureFromSurface(s *Surface) (*Texture, error) {
	ct := C.SDL_CreateTextureFromSurface(r.cptr(), s.cptr())
	if ct == nil {
		return nil, errors.New(Error())
	}
	return (*Texture)(unsafe.Pointer(ct)), nil
}

// Destroy destroys the specified texture.
func (tx *Texture) Destroy() {
	C.SDL_DestroyTexture(tx.cptr())
}

// AlphaMod gets the additional alpha value multiplied into render copy operations.
func (tx *Texture) AlphaMod() (uint, error) {
	var ca C.Uint8
	if C.SDL_GetTextureAlphaMod(tx.cptr(), &ca) < 0 {
		return 0, errors.New(Error())
	}
	return uint(ca), nil
}

// BlendMode gets the blend mode used for texture copy operations.
func (tx *Texture) BlendMode() (BlendMode, error) {
	var cb C.SDL_BlendMode
	if C.SDL_GetTextureBlendMode(tx.cptr(), &cb) < 0 {
		return 0, errors.New(Error())
	}
	return BlendMode(cb), nil
}

// ColorMod gets the additional color value multiplied into render copy operations.
func (tx *Texture) ColorMod() (r, g, b uint8, err error) {
	var cr, cg, cb C.Uint8
	if C.SDL_GetTextureColorMod(tx.cptr(), &cr, &cg, &cb) < 0 {
		err = errors.New(Error())
		return
	}
	r = uint8(cr)
	g = uint8(cg)
	b = uint8(cb)
	return
}

// SDL_LockTexture

// Query queries the attributes of a texture.
func (tx *Texture) Query() (format PixelFormat, access TextureAccess, w, h int, err error) {
	var cf C.Uint32
	var ca, cw, ch C.int
	if C.SDL_QueryTexture(tx.cptr(), &cf, &ca, &cw, &ch) < 0 {
		err = errors.New(Error())
		return
	}
	format = PixelFormat(cf)
	access = TextureAccess(ca)
	w = int(cw)
	h = int(ch)
	return
}

// SetAlphaMod sets an additional alpha value multiplied into render copy operations.
func (tx *Texture) SetAlphaMod(alpha uint8) error {
	if C.SDL_SetTextureAlphaMod(tx.cptr(), C.Uint8(alpha)) < 0 {
		return errors.New(Error())
	}
	return nil
}

// SDL_SetTextureBlendMode
// SDL_SetTextureColorMod
// SDL_UnlockTexture
// SDL_UpdateTexture
