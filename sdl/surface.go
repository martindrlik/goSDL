package sdl

// #include <SDL2/SDL_render.h>
import "C"
import "unsafe"

// Surface contains a collection of pixels used in software blitting.
type Surface C.SDL_Surface

func (s *Surface) cptr() *C.SDL_Surface {
	return (*C.SDL_Surface)(unsafe.Pointer(s))
}
