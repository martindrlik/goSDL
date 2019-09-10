package sdl

// #include <SDL2/SDL_events.h>
import "C"

import (
	"unsafe"
)

const (
	// 	Application events

	EvQuit = C.SDL_QUIT // user-requested quit; see Remarks for details

	// Android, iOS and WinRT events; see Remarks for details

	EvAppTerminating         = C.SDL_APP_TERMINATING         // OS is terminating the application
	EvAppLowMemory           = C.SDL_APP_LOWMEMORY           // OS is low on memory; free some
	EvAppWillEnterBackground = C.SDL_APP_WILLENTERBACKGROUND // application is entering background
	EvAppDidEnterBackground  = C.SDL_APP_DIDENTERBACKGROUND  // application entered background
	EvAppWillEnterForeground = C.SDL_APP_WILLENTERFOREGROUND // application is entering foreground
	EvAppDidEnterForeground  = C.SDL_APP_DIDENTERFOREGROUND  // application entered foreground

	// Window events

	EvWindowEvent = C.SDL_WINDOWEVENT // window state change
	EvSysWMEvent  = C.SDL_SYSWMEVENT  // system specific event

	// Keyboard events

	EvKeyDown       = C.SDL_KEYDOWN       // key pressed
	EvKeyUp         = C.SDL_KEYUP         // key released
	EvTextEditing   = C.SDL_TEXTEDITING   // keyboard text editing (composition)
	EvTextInput     = C.SDL_TEXTINPUT     // keyboard text input
	EvKeymapChanged = C.SDL_KEYMAPCHANGED // keymap changed due to a system event such as an input language or keyboard layout change (>= SDL 2.0.4)

	// Mouse events

	EvMouseMotion     = C.SDL_MOUSEMOTION     // mouse moved
	EvMouseButtonDown = C.SDL_MOUSEBUTTONDOWN // mouse button pressed
	EvMouseButtonUp   = C.SDL_MOUSEBUTTONUP   // mouse button released
	EvMouseWheel      = C.SDL_MOUSEWHEEL      // mouse wheel motion

	// Joystick events

	EvJoyAxisMotion    = C.SDL_JOYAXISMOTION    // joystick axis motion
	EvJoyBallMotion    = C.SDL_JOYBALLMOTION    // joystick trackball motion
	EvJoyHatmotion     = C.SDL_JOYHATMOTION     // joystick hat position change
	EvJoyButtonDown    = C.SDL_JOYBUTTONDOWN    // joystick button pressed
	EvJoyButtonUp      = C.SDL_JOYBUTTONUP      // joystick button released
	EvJoyDeviceAdded   = C.SDL_JOYDEVICEADDED   // joystick connected
	EvJoyDeviceRemoved = C.SDL_JOYDEVICEREMOVED // joystick disconnected

	// Controller events

	EvControllerAxisMotion     = C.SDL_CONTROLLERAXISMOTION     // controller axis motion
	EvControllerButtonDown     = C.SDL_CONTROLLERBUTTONDOWN     // controller button pressed
	EvControllerButtonUp       = C.SDL_CONTROLLERBUTTONUP       // controller button released
	EvControllerDeviceAdded    = C.SDL_CONTROLLERDEVICEADDED    // controller connected
	EvControllerDeviceRemoved  = C.SDL_CONTROLLERDEVICEREMOVED  // controller disconnected
	EvControllerDeviceRemapped = C.SDL_CONTROLLERDEVICEREMAPPED // controller mapping updated

	// Touch events

	// SDL_FINGERDOWN
	// user has touched input device
	// SDL_FINGERUP
	// user stopped touching input device
	// SDL_FINGERMOTION
	// user is dragging finger on input device
	// Gesture events
	// SDL_DOLLARGESTURE
	// SDL_DOLLARRECORD
	// SDL_MULTIGESTURE

	// Clipboard events

	EvClipboardUpdate = C.SDL_CLIPBOARDUPDATE // the clipboard changed

// Drag and drop events
// SDL_DROPFILE
// the system requests a file open
// SDL_DROPTEXT
// text/plain drag-and-drop event
// SDL_DROPBEGIN
// a new set of drops is beginning (>= SDL 2.0.5)
// SDL_DROPCOMPLETE
// current set of drops is now complete (>= SDL 2.0.5)
// Audio hotplug events
// SDL_AUDIODEVICEADDED
// a new audio device is available (>= SDL 2.0.4)
// SDL_AUDIODEVICEREMOVED
// an audio device has been removed (>= SDL 2.0.4)
// Render events
// SDL_RENDER_TARGETS_RESET
// the render targets have been reset and their contents need to be updated (>= SDL 2.0.2)
// SDL_RENDER_DEVICE_RESET
// the device has been reset and all textures need to be recreated (>= SDL 2.0.4)
// These are for your use, and should be allocated with SDL_RegisterEvents()
// SDL_USEREVENT
// a user-specified event
// SDL_LASTEVENT
// only for bounding internal arrays
)

// Event is a union that contains structures
// for the different event types.
type Event struct {
	Type uint32
	_    [52]byte
}

// SDL_AddEventWatch
// SDL_DelEventWatch
// SDL_EventState
// SDL_FilterEvents
// SDL_FlushEvent
// SDL_FlushEvents
// SDL_GetEventFilter
// SDL_GetEventState
// SDL_GetNumTouchDevices
// SDL_GetNumTouchFingers
// SDL_GetTouchDevice
// SDL_GetTouchFinger

// HasEvent checks for the existence of certain event type in the event queue.
func HasEvent(t int) bool {
	return C.SDL_HasEvent(C.Uint32(t)) == C.SDL_TRUE
}

// HasEvents checks for the existence of a range of event types in the event queue.
func HasEvents(min, max int) bool {
	return C.SDL_HasEvents(C.Uint32(min), C.Uint32(max)) == C.SDL_TRUE
}

// SDL_LoadDollarTemplates
// SDL_PeepEvents

// PollEvent is used to poll for currently
// pending events.
func PollEvent(ev *Event) int {
	return int(C.SDL_PollEvent((*C.SDL_Event)(unsafe.Pointer(ev))))
}

// PumpEvents pumps the event loop, gathering events from the input devices.
func PumpEvents() {
	C.SDL_PumpEvents()
}

// SDL_PushEvent
// SDL_QuitRequested
// SDL_RecordGesture
// SDL_RegisterEvents
// SDL_SaveAllDollarTemplates
// SDL_SaveDollarTemplate
// SDL_SetEventFilter
// SDL_WaitEvent
// SDL_WaitEventTimeout
