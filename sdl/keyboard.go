package sdl

// #include <SDL2/SDL_keyboard.h>
import "C"

import (
	"image"
	"reflect"
	"unsafe"
)

// KeyFromName gets a key code from a human-readable name.
func KeyFromName(name string) int {
	cs := C.CString(name)
	kc := C.SDL_GetKeyFromName(cs)
	C.free(unsafe.Pointer(cs))
	return int(kc)
}

// KeyFromScancode gets the key code corresponding to the given
// scancode according to the current keyboard layout.
func KeyFromScancode(scancode int) int {
	sc := C.SDL_Scancode(scancode)
	kc := C.SDL_GetKeyFromScancode(sc)
	return int(kc)
}

// KeyName gets a human-readable name for a key.
func KeyName(keycode int) string {
	kc := C.SDL_Keycode(keycode)
	cs := C.SDL_GetKeyName(kc)
	return C.GoString(cs)
}

// KeyboardState gets a snapshot of the current state of the keyboard.
func KeyboardState() []uint8 {
	var num C.int
	start := C.SDL_GetKeyboardState(&num)
	s := reflect.SliceHeader{
		Cap:  int(num),
		Data: uintptr(unsafe.Pointer(start)),
		Len:  int(num),
	}
	return *(*[]uint8)(unsafe.Pointer(&s))
}

// ModState gets the current key modifier state for the keyboard.
func ModState() int {
	ms := C.SDL_GetModState()
	return int(ms)
}

// ScancodeFromKey gets the scancode corresponding to the given
// key code according to the current keyboard layout.
func ScancodeFromKey(keycode int) int {
	kc := C.SDL_Keycode(keycode)
	sc := C.SDL_GetScancodeFromKey(kc)
	return int(sc)
}

// ScancodeFromName gets a scancode from a human-readable name.
func ScancodeFromName(name string) int {
	cs := C.CString(name)
	sc := C.SDL_GetScancodeFromName(cs)
	C.free(unsafe.Pointer(cs))
	return int(sc)
}

// ScancodeName gets a human-readable name for a scancode.
func ScancodeName(scancode int) string {
	sc := C.SDL_Scancode(scancode)
	cs := C.SDL_GetScancodeName(sc)
	return C.GoString(cs)
}

// HasScreenKeyboardSupport checks whether the platform
// has some screen keyboard support.
func HasScreenKeyboardSupport() bool {
	f := C.SDL_HasScreenKeyboardSupport()
	return f == C.SDL_TRUE
}

// IsTextInputActive checks whether or not Unicode
// text input events are enabled.
func IsTextInputActive() bool {
	f := C.SDL_IsTextInputActive()
	return f == C.SDL_TRUE
}

// SetModState sets the current key modifier
// state for the keyboard.
func SetModState(modstate int) {
	ms := C.SDL_Keymod(modstate)
	C.SDL_SetModState(ms)
}

// SetTextInputRect sets the rectangle used to type
// Unicode text inputs.
func SetTextInputRect(rect image.Rectangle) {
	cr := (*C.SDL_Rect)(unsafe.Pointer(&rect))
	C.SDL_SetTextInputRect(cr)
}

// StartTextInput starts accepting Unicode text
// input events.
func StartTextInput() {
	C.SDL_StartTextInput()
}

// StopTextInput stops receiving any text input events.
func StopTextInput() {
	C.SDL_StopTextInput()
}

// KeyboardFocus gets the window which currently has keyboard focus.
func KeyboardFocus() *Window {
	cw := C.SDL_GetKeyboardFocus()
	return (*Window)(unsafe.Pointer(cw))
}

// IsScreenKeyboardShown checks whether the screen keyboard
// is shown for given window.
func (window *Window) IsScreenKeyboardShown() bool {
	f := C.SDL_IsScreenKeyboardShown(window.cptr())
	return f == C.SDL_TRUE
}
