package sdl

// #include <SDL2/SDL_mouse.h>
import "C"
import (
	"errors"
)

func cbool(b bool) C.SDL_bool {
	if b {
		return C.SDL_TRUE
	}
	return C.SDL_FALSE
}

// CaptureMouse captures the mouse and to track input outside an SDL window.
func CaptureMouse(enabled bool) error {
	cb := cbool(enabled)
	if C.SDL_CaptureMouse(cb) == -1 {
		return errors.New(Error())
	}
	return nil
}

// SDL_CreateColorCursor
// SDL_CreateCursor
// SDL_CreateSystemCursor
// SDL_FreeCursor
// SDL_GetCursor
// SDL_GetDefaultCursor
// SDL_GetGlobalMouseState
// SDL_GetMouseFocus
// SDL_GetMouseState
// SDL_GetRelativeMouseMode
// SDL_GetRelativeMouseState
// SDL_SetCursor
// SDL_SetRelativeMouseMode

// ShowCursor toggles whether or not the cursor is shown.
func ShowCursor(toggle int) (int, error) {
	ct := C.int(toggle)
	cu := C.SDL_ShowCursor(ct)
	if cu < 0 {
		return 0, errors.New(Error())
	}
	return int(cu), nil
}

// SDL_WarpMouseGlobal
// SDL_WarpMouseInWindow
