package sdl

/*
#include <stdlib.h>
#include <SDL2/SDL_error.h>

void goSDL_SetError(const char *s) {
	SDL_SetError("%s", s);
}
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

// ClearError clears any previous error message.
func ClearError() {
	C.SDL_ClearError()
}

// Error retrieves a message about the last error that occurred.
func Error() string {
	cs := C.SDL_GetError()
	return C.GoString(cs)
}

// SetError sets the SDL error message.
func SetError(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	c := C.CString(s)
	C.goSDL_SetError(c)
	C.free(unsafe.Pointer(c))
}

func makeError(alt error) error {
	em := Error()
	if em == "" {
		return alt
	}
	return errors.New(em)
}
