package sdl

// #include <SDL2/SDL_video.h>
// #include <SDL2/SDL_render.h>
// #include <SDL2/SDL_messagebox.h>
import "C"
import (
	"errors"
	"unsafe"
)

type Window C.SDL_Window

const (
	// Fullscreen is flag for fullscreen window.
	Fullscreen = C.SDL_WINDOW_FULLSCREEN
	// FullscreenDesktop is flag for fullscreen window
	// at the current desktop resolution.
	FullscreenDesktop = C.SDL_WINDOW_FULLSCREEN_DESKTOP

	// SDL_WINDOW_OPENGL window usable with OpenGL context

	// Shown is flag for window is visible.
	Shown = C.SDL_WINDOW_SHOWN
	// Hidden is flag for window is not visible.
	Hidden = C.SDL_WINDOW_HIDDEN
	// Borderless is flag for no window decoration.
	Borderless = C.SDL_WINDOW_BORDERLESS
	// Resizable is flag for that window can be resized.
	Resizable = C.SDL_WINDOW_RESIZABLE
	// Minimized is flag for window is minimized.
	Minimized = C.SDL_WINDOW_MINIMIZED
	// Maximized is flag for window is maximized.
	Maximized = C.SDL_WINDOW_MAXIMIZED
	// InputGrabbed is flag for window has grabbed input focus.
	InputGrabbed = C.SDL_WINDOW_INPUT_GRABBED
	// InputFocus is flag for window has input focus.
	InputFocus = C.SDL_WINDOW_INPUT_FOCUS
	// MouseFocus is flag for window has mouse focus.
	MouseFocus = C.SDL_WINDOW_MOUSE_FOCUS

	// SDL_WINDOW_FOREIGN window not created by SDL

	// AllowHighDPI is flag for that window should be created in
	// high-DPI mode if supported (>WindowFlag = SDL 2.0.1).
	AllowHighDPI = C.SDL_WINDOW_ALLOW_HIGHDPI
	// MouseCapture is flag for that window has mouse
	// captured (unrelated to INPUT_GRABBED, >WindowFlag = SDL 2.0.4).
	MouseCapture = C.SDL_WINDOW_MOUSE_CAPTURE
	// AlwaysOnTop is flag for that window should always be
	// above others (X11 only, >WindowFlag = SDL 2.0.5).
	AlwaysOnTop = C.SDL_WINDOW_ALWAYS_ON_TOP
	// SkipTaskbar is flag for that window should not be added to
	// the taskbar (X11 only, >WindowFlag = SDL 2.0.5).
	SkipTaskbar = C.SDL_WINDOW_SKIP_TASKBAR
	// Utility is flag for that window should be treated as
	// a utility window (X11 only, >WindowFlag = SDL 2.0.5).
	Utility = C.SDL_WINDOW_UTILITY
	// Tooltip is flag for that window should be treated as
	// a tooltip (X11 only, >WindowFlag = SDL 2.0.5).
	Tooltip = C.SDL_WINDOW_TOOLTIP
	// PopupMenu is flag for that window should be treated as
	// a popup menu (X11 only, >WindowFlag = SDL 2.0.5).
	PopupMenu = C.SDL_WINDOW_POPUP_MENU
)

const (
	// MsgError is flag for displaying error dialog.
	MsgError = C.SDL_MESSAGEBOX_ERROR
	// MsgWarning is flag for displaying warning dialog.
	MsgWarning = C.SDL_MESSAGEBOX_WARNING
	// MsgInformation is flag for displaying informational dialog.
	MsgInformation = C.SDL_MESSAGEBOX_INFORMATION
)

var (
	ErrNoWindowForID       = errors.New("no window associated with given id")
	ErrInvalidDisplayIndex = errors.New("invalid display index or failure")
)

func (w *Window) cptr() *C.SDL_Window {
	return (*C.SDL_Window)(unsafe.Pointer(w))
}

// CreateWindow creates a window with the specified
// position, dimensions, and flags.
func CreateWindow(title string, x, y, width, height int, flags uint32) (*Window, error) {
	cw := C.SDL_CreateWindow(
		C.CString(title),
		C.int(x),
		C.int(y),
		C.int(width),
		C.int(height),
		C.Uint32(flags))
	if cw == nil {
		return nil, errors.New(Error())
	}
	w := (*Window)(unsafe.Pointer(cw))
	return w, nil
}

// CreateWindowAndRenderer creates a window and default renderer.
func CreateWindowAndRenderer(width, height int, flags uint32) (*Window, *Renderer, error) {
	var (
		cw *C.SDL_Window
		cr *C.SDL_Renderer
	)
	if C.SDL_CreateWindowAndRenderer(
		C.int(width),
		C.int(height),
		C.Uint32(flags),
		&cw,
		&cr) != 0 {
		return nil, nil, errors.New(Error())
	}
	w := (*Window)(unsafe.Pointer(cw))
	r := (*Renderer)(unsafe.Pointer(cr))
	return w, r, nil
}

// SDL_CreateWindowFrom

// Destroy destroys a window.
func (w *Window) Destroy() {
	C.SDL_DestroyWindow(w.cptr())
}

// DisableScreenSaver prevents the screen from
// being blanked by a screen saver.
//
// If you disable the screensaver, it is automatically
// re-enabled when SDL quits.
func DisableScreenSaver() {
	C.SDL_DisableScreenSaver()
}

// EnableScreenSaver allows the screen to be
// blanked by a screen saver.
func EnableScreenSaver() {
	C.SDL_EnableScreenSaver()
}

// SDL_GL_CreateContext
// SDL_GL_DeleteContext
// SDL_GL_ExtensionSupported
// SDL_GL_GetAttribute
// SDL_GL_GetCurrentContext
// SDL_GL_GetCurrentWindow
// SDL_GL_GetDrawableSize
// SDL_GL_GetProcAddress
// SDL_GL_GetSwapInterval
// SDL_GL_LoadLibrary
// SDL_GL_MakeCurrent
// SDL_GL_ResetAttributes
// SDL_GL_SetAttribute
// SDL_GL_SetSwapInterval
// SDL_GL_SwapWindow
// SDL_GL_UnloadLibrary
// SDL_GetClosestDisplayMode
// SDL_GetCurrentDisplayMode
// SDL_GetCurrentVideoDriver
// SDL_GetDesktopDisplayMode
// SDL_GetDisplayBounds
// SDL_GetDisplayDPI
// SDL_GetDisplayMode

// DisplayName gets the name of a display in UTF-8 encoding.
func DisplayName(displayIndex int) (string, error) {
	cs := C.SDL_GetDisplayName(C.int(displayIndex))
	if cs == nil {
		return "", makeError(ErrInvalidDisplayIndex)
	}
	return C.GoString(cs), nil
}

// SDL_GetDisplayUsableBounds
// SDL_GetGrabbedWindow

// NumDisplayModes gets the number of available display modes.
func NumDisplayModes(displayIndex int) (int, error) {
	return intError(C.SDL_GetNumDisplayModes(C.int(displayIndex)))
}

// NumDisplays gets the number of available video displays.
func NumDisplays() (int, error) {
	return intError(C.SDL_GetNumVideoDisplays())
}

// NumDrivers gets the number of video drivers compiled into SDL.
func NumDrivers() (int, error) {
	return intError(C.SDL_GetNumVideoDrivers())
}

func intError(result C.int) (int, error) {
	if result < 0 {
		return 0, errors.New(Error())
	}
	return int(result), nil
}

// SDL_GetVideoDriver
// SDL_GetWindowBordersSize

// Brightness gets the brightness (gamma multiplier) for the display that owns this window.
func (w *Window) Brightness() float32 {
	return float32(C.SDL_GetWindowBrightness(w.cptr()))
}

// SDL_GetWindowData
// SDL_GetWindowDisplayIndex
// SDL_GetWindowDisplayMode
// SDL_GetWindowFlags

// WindowFromID gets a window from a stored ID.
func WindowFromID(id uint32) (*Window, error) {
	cid := C.Uint32(id)
	w := C.SDL_GetWindowFromID(cid)
	if w == nil {
		return nil, makeError(ErrNoWindowForID)
	}
	return (*Window)(unsafe.Pointer(w)), nil
}

// SDL_GetWindowGammaRamp
// SDL_GetWindowGrab

// ID gets the numeric ID of a window, for logging purposes.
func (w *Window) ID() (uint32, error) {
	cid := C.SDL_GetWindowID(w.cptr())
	if cid == 0 {
		return 0, errors.New(Error())
	}
	return uint32(cid), nil
}

// MaximumSize gets the maximum size of a window's client area.
func (w *Window) MaximumSize() (maximumWidth, maximumHeight int) {
	var x, y C.int
	C.SDL_GetWindowMaximumSize(w.cptr(), &x, &y)
	return int(x), int(y)
}

// MinimumSize gets the minimum size of a window's client area.
func (w *Window) MinimumSize() (minimumWidth, minimumHeight int) {
	var x, y C.int
	C.SDL_GetWindowMinimumSize(w.cptr(), &x, &y)
	return int(x), int(y)
}

// SDL_GetWindowOpacity
// SDL_GetWindowPixelFormat

// Position gets the position of a window.
func (w *Window) Position() (x, y int) {
	var px, py C.int
	C.SDL_GetWindowPosition(w.cptr(), &px, &py)
	return int(px), int(py)
}

// Size gets the size of a window's client area.
func (w *Window) Size() (width, height int) {
	var x, y C.int
	C.SDL_GetWindowSize(w.cptr(), &x, &y)
	return int(x), int(y)
}

// SDL_GetWindowSurface

// Title gets the title of a window.
func (w *Window) Title() string {
	return C.GoString(C.SDL_GetWindowTitle(w.cptr()))
}

// SDL_GetWindowWMInfo

// Hide hides a window.
func (w *Window) Hide() {
	C.SDL_HideWindow(w.cptr())
}

// SDL_IsScreenSaverEnabled

// Maximize makes a window as large as possible.
func (w *Window) Maximize() {
	C.SDL_MaximizeWindow(w.cptr())
}

// Minimize minimizes a window to an iconic representation.
func (w *Window) Minimize() {
	C.SDL_MinimizeWindow(w.cptr())
}

// Raise raises a window above other windows and set the input focus.
func (w *Window) Raise() {
	C.SDL_RaiseWindow(w.cptr())
}

// Restore restores the size and position of a minimized
// or maximized window.
func (w *Window) Restore() {
	C.SDL_RestoreWindow(w.cptr())
}

// SDL_SetWindowBordered

// SetBrightness sets the brightness (gamma multiplier) for the display that owns this window.
// Brightness f value should be set where 0.0 is completely dark and 1.0 is normal brightness.
func (w *Window) SetBrightness(f float32) error {
	if C.SDL_SetWindowBrightness(w.cptr(), C.float(f)) >= 0 {
		return nil
	}
	return errors.New(Error())
}

// SDL_SetWindowData
// SDL_SetWindowDisplayMode

// SetFullscreen sets a window's fullscreen state.
//
// Flags may be Fullscreen, for "real" fullscreen with a videomode change;
// FullscreenDesktop for "fake" fullscreen that takes the size of the desktop;
// and 0 for windowed mode.
func (w *Window) SetFullscreen(flags uint32) error {
	if C.SDL_SetWindowFullscreen(
		w.cptr(),
		C.Uint32(flags)) < 0 {
		return errors.New(Error())
	}
	return nil
}

// SDL_SetWindowGammaRamp
// SDL_SetWindowGrab
// SDL_SetWindowHitTest
// SDL_SetWindowIcon
// SDL_SetWindowInputFocus

// SetMaximumSize sets the maximum size of a window's client area.
func (w *Window) SetMaximumSize(maximumWidth, maximumHeight int) {
	C.SDL_SetWindowMaximumSize(w.cptr(), C.int(maximumWidth), C.int(maximumHeight))
}

// SetMinimumSize set the minimum size of a window's client area.
func (w *Window) SetMinimumSize(minimumWidth, minimumHeight int) {
	C.SDL_SetWindowMinimumSize(w.cptr(), C.int(minimumWidth), C.int(minimumHeight))
}

// SDL_SetWindowModalFor
// SDL_SetWindowOpacity

// SetPosition sets the position of a window.
func (w *Window) SetPosition(x, y int) {
	C.SDL_SetWindowPosition(w.cptr(), C.int(x), C.int(y))
}

// SetResizable sets the user-resizable state of a window.
func (w *Window) SetResizable(isResizable bool) {
	if isResizable {
		C.SDL_SetWindowResizable(w.cptr(), C.SDL_TRUE)
	} else {
		C.SDL_SetWindowResizable(w.cptr(), C.SDL_FALSE)
	}
}

// SetSize sets the size of a window's client area.
func (w *Window) SetSize(width, height int) {
	C.SDL_SetWindowSize(w.cptr(), C.int(width), C.int(height))
}

// SetTitle sets the title of a window.
func (w *Window) SetTitle(title string) {
	C.SDL_SetWindowTitle(w.cptr(), C.CString(title))
}

// SDL_ShowMessageBox

// ShowSimpleMessageBox displays a simple modal message box.
func ShowSimpleMessageBox(flags uint32, title, message string, parent *Window) error {
	if C.SDL_ShowSimpleMessageBox(
		C.Uint32(flags),
		C.CString(title),
		C.CString(message),
		parent.cptr()) >= 0 {
		return nil
	}
	return errors.New(Error())
}

// Show shows a window.
func (w *Window) Show() {
	C.SDL_ShowWindow(w.cptr())
}

// SDL_UpdateWindowSurface
// SDL_UpdateWindowSurfaceRects
// SDL_VideoInit
// SDL_VideoQuit
