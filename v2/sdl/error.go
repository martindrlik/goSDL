package sdl

// #include <SDL2/SDL_error.h>
// #cgo LDFLAGS: -lSDL2
import "C"
import "errors"

// ClearError clears any previous error message.
func ClearError() {
	C.SDL_ClearError()
}

// Error retrieves a message about the last
// error that occurred.
func Error() error {
	cs := C.SDL_GetError()
	s := C.GoString(cs)
	if s == "" {
		return nil
	}
	return errors.New(s)
}
