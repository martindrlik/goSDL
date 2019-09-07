package clipboard

// #include <SDL2/SDL_clipboard.h>
// #cgo LDFLAGS: -lSDL2
import "C"

import (
	"errors"
	"unsafe"

	"github.com/martindrlik/goSDL/sdl"
)

// Text gets UTF-8 text from the clipboard.
func Text() (string, error) {
	cs := C.SDL_GetClipboardText()
	if cs == nil {
		return "", errors.New(sdl.Error())
	}
	s := C.GoString(cs)
	C.SDL_free(unsafe.Pointer(cs))
	return s, nil
}

// HasText gets a flag indicating whether the clipboard
// exists and contains a text string that is non-empty.
func HasText() bool {
	return C.SDL_HasClipboardText() == C.SDL_TRUE
}

// SetText puts UTF-8 text into the clipboard.
func SetText(text string) error {
	ct := C.CString(text)
	if C.SDL_SetClipboardText(ct) < 0 {
		return errors.New(sdl.Error())
	}
	return nil
}
