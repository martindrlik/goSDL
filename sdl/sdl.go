package sdl

// #include <SDL2/SDL.h>
// #cgo LDFLAGS: -lSDL2
import "C"
import "errors"

const (
	Timer          = C.SDL_INIT_TIMER          // timer subsystem
	Audio          = C.SDL_INIT_AUDIO          // audio subsystem
	Video          = C.SDL_INIT_VIDEO          // video subsystem; automatically initializes the events subsystem
	Joystick       = C.SDL_INIT_JOYSTICK       // joystick subsystem; automatically initializes the events subsystem
	Haptic         = C.SDL_INIT_HAPTIC         // haptic (force feedback) subsystem
	GameController = C.SDL_INIT_GAMECONTROLLER // controller subsystem; automatically initializes the joystick subsystem
	Events         = C.SDL_INIT_EVENTS         // events subsystem
	Everything     = C.SDL_INIT_EVERYTHING     // all of the above subsystems
	NoParachute    = C.SDL_INIT_NOPARACHUTE    // compatibility; this flag is ignored
)

// Init is used to initialize the SDL library. This must be
// called before using most other SDL functions.
func Init(flags uint) error {
	cf := C.uint(flags)
	if C.SDL_Init(cf) < 0 {
		return errors.New(Error())
	}
	return nil
}

// InitSubSystem initializes specific SDL subsystems.
func InitSubSystem(flags uint) error {
	cf := C.uint(flags)
	if C.SDL_InitSubSystem(cf) < 0 {
		return errors.New(Error())
	}
	return nil
}

// Quit is used to clean up all initialized subsystems.
// You should call it upon all exit conditions.
func Quit() { C.SDL_Quit() }

// QuitSubSystem shuts down specific SDL subsystems.
func QuitSubSystem(flags uint) {
	cf := C.uint(flags)
	C.SDL_QuitSubSystem(cf)
}

// SetMainReady circumvents failure of SDL_Init() when
// not using SDL_main() as an entry point.
func SetMainReady() { C.SDL_SetMainReady() }

// WasInit gets a mask of the specified subsystems
// which have previously been initialized.
func WasInit(flags uint) uint {
	cf := C.uint(flags)
	cm := C.SDL_WasInit(cf)
	return uint(cm)
}

// SDL_WinRTRunApp
