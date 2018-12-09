package video

// #include <SDL2/SDL_video.h>
// #include <SDL2/SDL_render.h>
// #cgo LDFLAGS: -lSDL2
import "C"
import (
	"unsafe"

	"github.com/martindrlik/sdlbindings/v2/sdl"
)

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
	// high-DPI mode if supported (>= SDL 2.0.1).
	AllowHighDPI = C.SDL_WINDOW_ALLOW_HIGHDPI
	// MouseCapture is flag for that window has mouse
	// captured (unrelated to INPUT_GRABBED, >= SDL 2.0.4).
	MouseCapture = C.SDL_WINDOW_MOUSE_CAPTURE
	// AlwaysOnTop is flag for that window should always be
	// above others (X11 only, >= SDL 2.0.5).
	AlwaysOnTop = C.SDL_WINDOW_ALWAYS_ON_TOP
	// SkipTaskbar is flag for that window should not be added to
	// the taskbar (X11 only, >= SDL 2.0.5).
	SkipTaskbar = C.SDL_WINDOW_SKIP_TASKBAR
	// Utility is flag for that window should be treated as
	// a utility window (X11 only, >= SDL 2.0.5).
	Utility = C.SDL_WINDOW_UTILITY
	// Tooltip is flag for that window should be treated as
	// a tooltip (X11 only, >= SDL 2.0.5).
	Tooltip = C.SDL_WINDOW_TOOLTIP
	// PopupMenu is flag for that window should be treated as
	// a popup menu (X11 only, >= SDL 2.0.5).
	PopupMenu = C.SDL_WINDOW_POPUP_MENU
)

func init() {
	sdl.Register(sdl.Video)
}

type Window C.SDL_Window

func (window *Window) cptr() *C.SDL_Window {
	return (*C.SDL_Window)(unsafe.Pointer(window))
}

// SDL_CreateWindow

// CreateWindowAndRenderer creates a window and default renderer.
func CreateWindowAndRenderer(width, height int, flags uint32) (*Window, *Renderer, error) {
	var (
		window   *C.SDL_Window
		renderer *C.SDL_Renderer
	)
	if C.SDL_CreateWindowAndRenderer(
		C.int(width),
		C.int(height),
		C.Uint32(flags),
		&window,
		&renderer) != 0 {
		return nil, nil, sdl.Error()
	}
	return (*Window)(unsafe.Pointer(window)), (*Renderer)(unsafe.Pointer(renderer)), nil
}

// SDL_CreateWindowFrom

// Destroy destroys a window.
func (window *Window) Destroy() {
	C.SDL_DestroyWindow(window.cptr())
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
// SDL_GetDisplayName
// SDL_GetDisplayUsableBounds
// SDL_GetGrabbedWindow
// SDL_GetNumDisplayModes
// SDL_GetNumVideoDisplays
// SDL_GetNumVideoDrivers
// SDL_GetVideoDriver
// SDL_GetWindowBordersSize
// SDL_GetWindowBrightness
// SDL_GetWindowData
// SDL_GetWindowDisplayIndex
// SDL_GetWindowDisplayMode
// SDL_GetWindowFlags
// SDL_GetWindowFromID
// SDL_GetWindowGammaRamp
// SDL_GetWindowGrab
// SDL_GetWindowID
// SDL_GetWindowMaximumSize
// SDL_GetWindowMinimumSize
// SDL_GetWindowOpacity
// SDL_GetWindowPixelFormat
// SDL_GetWindowPosition
// SDL_GetWindowSize
// SDL_GetWindowSurface

// Title gets the title of a window.
func (window *Window) Title() string {
	return C.GoString(C.SDL_GetWindowTitle(window.cptr()))
}

// SDL_GetWindowWMInfo

// Hide hides a window.
func (window *Window) Hide() {
	C.SDL_HideWindow(window.cptr())
}

// SDL_IsScreenSaverEnabled

// Maximize makes a window as large as possible.
func (window *Window) Maximize() {
	C.SDL_MaximizeWindow(window.cptr())
}

// Minimize minimizes a window to an iconic representation.
func (window *Window) Minimize() {
	C.SDL_MinimizeWindow(window.cptr())
}

// Raise raises a window above other windows and set the input focus.
func (window *Window) Raise() {
	C.SDL_RaiseWindow(window.cptr())
}

// Restore restores the size and position of a minimized
// or maximized window.
func (window *Window) Restore() {
	C.SDL_RestoreWindow(window.cptr())
}

// SDL_SetWindowBordered
// SDL_SetWindowBrightness
// SDL_SetWindowData
// SDL_SetWindowDisplayMode
// SDL_SetWindowFullscreen
// SDL_SetWindowGammaRamp
// SDL_SetWindowGrab
// SDL_SetWindowHitTest
// SDL_SetWindowIcon
// SDL_SetWindowInputFocus
// SDL_SetWindowMaximumSize
// SDL_SetWindowMinimumSize
// SDL_SetWindowModalFor
// SDL_SetWindowOpacity
// SDL_SetWindowPosition
// SDL_SetWindowResizable
// SDL_SetWindowSize

// SetTitle sets the title of a window.
func (window *Window) SetTitle(title string) {
	C.SDL_SetWindowTitle(window.cptr(), C.CString(title))
}

// SDL_ShowMessageBox
// SDL_ShowSimpleMessageBox

// Show shows a window.
func (window *Window) Show() {
	C.SDL_ShowWindow(window.cptr())
}

// SDL_UpdateWindowSurface
// SDL_UpdateWindowSurfaceRects
// SDL_VideoInit
// SDL_VideoQuit
