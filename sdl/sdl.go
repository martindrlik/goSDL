package sdl

// #include <SDL2/SDL.h>
// #cgo LDFLAGS: -lSDL2
import "C"
import "errors"

type Subsystem uint

const (
	Audio  Subsystem = C.SDL_INIT_AUDIO
	Video  Subsystem = C.SDL_INIT_VIDEO
	Events Subsystem = C.SDL_INIT_EVENTS
)

var subsystems Subsystem

func Register(subsystem Subsystem) {
	subsystems |= subsystem
}

// Init is used to initialize the SDL library. This must be
// called before using most other SDL functions.
func Init() error {
	if C.SDL_Init(C.uint(subsystems)) < 0 {
		return errors.New(Error())
	}
	return nil
}

// SDL_InitSubSystem

// Quit is used to clean up all initialized subsystems.
// You should call it upon all exit conditions.
func Quit() {
	C.SDL_Quit()
}

// SDL_QuitSubSystem
// SDL_SetMainReady
// SDL_WasInit
// SDL_WinRTRunApp
