package sdl

// #include <SDL2/SDL_error.h>
// #cgo LDFLAGS: -lSDL2
import "C"
import "errors"

var (
	// ErrFailure indicates that there is an error,
	// but SDL_GetError returned no error.
	ErrFailure = errors.New("failure")
)

// ClearError clears any previous error message.
func ClearError() {
	C.SDL_ClearError()
}

// Error retrieves a message about the last
// error that occurred.
func Error(alternativeError error) error {
	cs := C.SDL_GetError()
	s := C.GoString(cs)
	if s == "" {
		return alternativeError
	}
	return errors.New(s)
}
