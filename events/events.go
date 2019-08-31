package events

// #include <SDL2/SDL_events.h>
// #cgo LDFLAGS: -lSDL2
import "C"

import (
	"unsafe"

	"github.com/martindrlik/goSDL/sdl"
)

const (
	// 	Application events

	Quit = C.SDL_QUIT // user-requested quit; see Remarks for details

	// Android, iOS and WinRT events; see Remarks for details

	AppTerminating         = C.SDL_APP_TERMINATING         // OS is terminating the application
	AppLowMemory           = C.SDL_APP_LOWMEMORY           // OS is low on memory; free some
	AppWillEnterBackground = C.SDL_APP_WILLENTERBACKGROUND // application is entering background
	AppDidEnterBackground  = C.SDL_APP_DIDENTERBACKGROUND  // application entered background
	AppWillEnterForeground = C.SDL_APP_WILLENTERFOREGROUND // application is entering foreground
	AppDidEnterForeground  = C.SDL_APP_DIDENTERFOREGROUND  // application entered foreground

	// Window events

	WindowEvent = C.SDL_WINDOWEVENT // window state change
	SysWMEvent  = C.SDL_SYSWMEVENT  // system specific event

	// Keyboard events

	KeyDown       = C.SDL_KEYDOWN       // key pressed
	KeyUp         = C.SDL_KEYUP         // key released
	TextEditing   = C.SDL_TEXTEDITING   // keyboard text editing (composition)
	TextInput     = C.SDL_TEXTINPUT     // keyboard text input
	KeymapChanged = C.SDL_KEYMAPCHANGED // keymap changed due to a system event such as an input language or keyboard layout change (>= SDL 2.0.4)

	// Mouse events

	MouseMotion     = C.SDL_MOUSEMOTION     // mouse moved
	MouseButtonDown = C.SDL_MOUSEBUTTONDOWN // mouse button pressed
	MouseButtonUp   = C.SDL_MOUSEBUTTONUP   // mouse button released
	MouseWheel      = C.SDL_MOUSEWHEEL      // mouse wheel motion

	// Joystick events

	JoyAxisMotion    = C.SDL_JOYAXISMOTION    // joystick axis motion
	JoyBallMotion    = C.SDL_JOYBALLMOTION    // joystick trackball motion
	JoyHatmotion     = C.SDL_JOYHATMOTION     // joystick hat position change
	JoyButtonDown    = C.SDL_JOYBUTTONDOWN    // joystick button pressed
	JoyButtonUp      = C.SDL_JOYBUTTONUP      // joystick button released
	JoyDeviceAdded   = C.SDL_JOYDEVICEADDED   // joystick connected
	JoyDeviceRemoved = C.SDL_JOYDEVICEREMOVED // joystick disconnected

	// Controller events

	ControllerAxisMotion     = C.SDL_CONTROLLERAXISMOTION     // controller axis motion
	ControllerButtonDown     = C.SDL_CONTROLLERBUTTONDOWN     // controller button pressed
	ControllerButtonUp       = C.SDL_CONTROLLERBUTTONUP       // controller button released
	ControllerDeviceAdded    = C.SDL_CONTROLLERDEVICEADDED    // controller connected
	ControllerDeviceRemoved  = C.SDL_CONTROLLERDEVICEREMOVED  // controller disconnected
	ControllerDeviceRemapped = C.SDL_CONTROLLERDEVICEREMAPPED // controller mapping updated

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

	ClipboardUpdate = C.SDL_CLIPBOARDUPDATE // the clipboard changed

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

func init() {
	sdl.Register(sdl.Events)
}

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
// SDL_HasEvent
// SDL_HasEvents
// SDL_LoadDollarTemplates
// SDL_PeepEvents

// PollEvent is used to poll for currently
// pending events.
func PollEvent(ev *Event) int {
	return int(C.SDL_PollEvent((*C.SDL_Event)(unsafe.Pointer(ev))))
}

// SDL_PumpEvents
// SDL_PushEvent
// SDL_QuitRequested
// SDL_RecordGesture
// SDL_RegisterEvents
// SDL_SaveAllDollarTemplates
// SDL_SaveDollarTemplate
// SDL_SetEventFilter
// SDL_WaitEvent
// SDL_WaitEventTimeout
