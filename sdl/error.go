package sdl

// #include <SDL2/SDL_error.h>
import "C"

// ClearError clears any previous error message.
func ClearError() {
	C.SDL_ClearError()
}

// Error retrieves a message about the last error that occurred.
func Error() string {
	cs := C.SDL_GetError()
	return C.GoString(cs)
}
