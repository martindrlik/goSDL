package sdl

// #include <SDL2/SDL_keycode.h>
import "C"

const (
	KmodNone   = C.KMOD_NONE   // 0 (no modifier is applicable)
	KmodLShift = C.KMOD_LSHIFT // the left Shift key is down
	KmodRShift = C.KMOD_RSHIFT // the right Shift key is down
	KmodLCtrl  = C.KMOD_LCTRL  // the left Ctrl (Control) key is down
	KmodRCtrl  = C.KMOD_RCTRL  // the right Ctrl (Control) key is down
	KmodLAlt   = C.KMOD_LALT   // the left Alt key is down
	KmodRAlt   = C.KMOD_RALT   // the right Alt key is down
	KmodLGUI   = C.KMOD_LGUI   // the left GUI key (often the Windows key) is down
	KmodRGUI   = C.KMOD_RGUI   // the right GUI key (often the Windows key) is down
	KmodNum    = C.KMOD_NUM    // the Num Lock key (may be located on an extended keypad) is down
	KmodCaps   = C.KMOD_CAPS   // the Caps Lock key is down
	KmodMode   = C.KMOD_MODE   // the AltGr key is down
	KmodCtrl   = C.KMOD_CTRL   // (KmodLCtrl|KmodRCtrl)
	KmodShift  = C.KMOD_SHIFT  // (KmodLShift|KmodRShift)
	KmodAlt    = C.KMOD_ALT    // (KmodLAlt|KmodRAlt)
	KmodGUI    = C.KMOD_GUI    // (KmodLGUI|KmodRGUI)
)
