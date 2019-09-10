package sdl

// #include <SDL2/SDL_keycode.h>
import "C"

const (
	Key0 = C.SDLK_0
	Key1 = C.SDLK_1
	Key2 = C.SDLK_2
	Key3 = C.SDLK_3
	Key4 = C.SDLK_4
	Key5 = C.SDLK_5
	Key6 = C.SDLK_6
	Key7 = C.SDLK_7
	Key8 = C.SDLK_8
	Key9 = C.SDLK_9

	KeyA           = C.SDLK_a
	KeyACBack      = C.SDLK_AC_BACK
	KeyACBookmarks = C.SDLK_AC_BOOKMARKS
	KeyACForward   = C.SDLK_AC_FORWARD
	KeyACHome      = C.SDLK_AC_HOME
	KeyACRefresh   = C.SDLK_AC_REFRESH
	KeyACSearch    = C.SDLK_AC_SEARCH
	KeyACStop      = C.SDLK_AC_STOP
	KeyAgain       = C.SDLK_AGAIN
	KeyAltErase    = C.SDLK_ALTERASE

	KeyQuote       = C.SDLK_QUOTE
	KeyApplication = C.SDLK_APPLICATION
	KeyAudioMute   = C.SDLK_AUDIOMUTE
	KeyAudioNext   = C.SDLK_AUDIONEXT
	KeyAudioPlay   = C.SDLK_AUDIOPLAY
	KeyAudioPrev   = C.SDLK_AUDIOPREV
	KeyAudioStop   = C.SDLK_AUDIOSTOP
	KeyB           = C.SDLK_b
	KeyBackslash   = C.SDLK_BACKSLASH
	KeyBackspace   = C.SDLK_BACKSPACE

	KeyBrightnessDown = C.SDLK_BRIGHTNESSDOWN
	KeyBrightnessUp   = C.SDLK_BRIGHTNESSUP
	KeyC              = C.SDLK_c
	KeyCalculator     = C.SDLK_CALCULATOR
	KeyCancel         = C.SDLK_CANCEL
	KeyCapsLock       = C.SDLK_CAPSLOCK
	KeyClear          = C.SDLK_CLEAR
	KeyClearAgain     = C.SDLK_CLEARAGAIN
	KeyComma          = C.SDLK_COMMA
	KeyComputer       = C.SDLK_COMPUTER

	KeyCopy             = C.SDLK_COPY
	KeyCrSel            = C.SDLK_CRSEL
	KeyCurrencySubUnit  = C.SDLK_CURRENCYSUBUNIT
	KeyCurrencyUnit     = C.SDLK_CURRENCYUNIT
	KeyCut              = C.SDLK_CUT
	KeyD                = C.SDLK_d
	KeyDecimalSeparator = C.SDLK_DECIMALSEPARATOR
	KeyDelete           = C.SDLK_DELETE
	KeyDisplaySwitch    = C.SDLK_DISPLAYSWITCH
	KeyDown             = C.SDLK_DOWN

	KeyE       = C.SDLK_e
	KeyEject   = C.SDLK_EJECT
	KeyEnd     = C.SDLK_END
	KeyEquals  = C.SDLK_EQUALS
	KeyEscape  = C.SDLK_ESCAPE
	KeyExecute = C.SDLK_EXECUTE
	KeyExSel   = C.SDLK_EXSEL
	KeyF       = C.SDLK_f
	KeyF1      = C.SDLK_F1
	KeyF10     = C.SDLK_F10

	KeyF11 = C.SDLK_F11
	KeyF12 = C.SDLK_F12
	KeyF13 = C.SDLK_F13
	KeyF14 = C.SDLK_F14
	KeyF15 = C.SDLK_F15
	KeyF16 = C.SDLK_F16
	KeyF17 = C.SDLK_F17
	KeyF18 = C.SDLK_F18
	KeyF19 = C.SDLK_F19
	KeyF2  = C.SDLK_F2

	KeyF20 = C.SDLK_F20
	KeyF21 = C.SDLK_F21
	KeyF22 = C.SDLK_F22
	KeyF23 = C.SDLK_F23
	KeyF24 = C.SDLK_F24
	KeyF3  = C.SDLK_F3
	KeyF4  = C.SDLK_F4
	KeyF5  = C.SDLK_F5
	KeyF6  = C.SDLK_F6
	KeyF7  = C.SDLK_F7

	KeyF8        = C.SDLK_F8
	KeyF9        = C.SDLK_F9
	KeyFind      = C.SDLK_FIND
	KeyG         = C.SDLK_g
	KeyBackQuote = C.SDLK_BACKQUOTE
	KeyH         = C.SDLK_h
	KeyHelp      = C.SDLK_HELP
	KeyHome      = C.SDLK_HOME
	KeyI         = C.SDLK_i
	KeyInsert    = C.SDLK_INSERT

	KeyJ              = C.SDLK_j
	KeyK              = C.SDLK_k
	KeyKBDIllumDown   = C.SDLK_KBDILLUMDOWN
	KeyKBDIllumToggle = C.SDLK_KBDILLUMTOGGLE
	KeyKBDIllumUp     = C.SDLK_KBDILLUMUP
	Keypad0           = C.SDLK_KP_0
	Keypad00          = C.SDLK_KP_00
	Keypad000         = C.SDLK_KP_000
	Keypad1           = C.SDLK_KP_1
	Keypad2           = C.SDLK_KP_2

	Keypad3         = C.SDLK_KP_3
	Keypad4         = C.SDLK_KP_4
	Keypad5         = C.SDLK_KP_5
	Keypad6         = C.SDLK_KP_6
	Keypad7         = C.SDLK_KP_7
	Keypad8         = C.SDLK_KP_8
	Keypad9         = C.SDLK_KP_9
	KeypadA         = C.SDLK_KP_A
	KeypadAmpersand = C.SDLK_KP_AMPERSAND
	KeypadAt        = C.SDLK_KP_AT

	KeypadB            = C.SDLK_KP_B
	KeypadBackspace    = C.SDLK_KP_BACKSPACE
	KeypadBinary       = C.SDLK_KP_BINARY
	KeypadC            = C.SDLK_KP_C
	KeypadClear        = C.SDLK_KP_CLEAR
	KeypadClearEntry   = C.SDLK_KP_CLEARENTRY
	KeypadColon        = C.SDLK_KP_COLON
	KeypadComma        = C.SDLK_KP_COMMA
	KeypadD            = C.SDLK_KP_D
	KeypadDblAmpersand = C.SDLK_KP_DBLAMPERSAND

	KeypadDblVerticalBar = C.SDLK_KP_DBLVERTICALBAR
	KeypadDecimal        = C.SDLK_KP_DECIMAL
	Keypad               = C.SDLK_KP_DIVIDE
	KeypadE              = C.SDLK_KP_E
	KeypadEnter          = C.SDLK_KP_ENTER
	KeypadEquals         = C.SDLK_KP_EQUALS
	KeypadEqualsAS400    = C.SDLK_KP_EQUALSAS400
	KeypadExclam         = C.SDLK_KP_EXCLAM
	KeypadF              = C.SDLK_KP_F
	KeypadGreater        = C.SDLK_KP_GREATER

	KeypadHash        = C.SDLK_KP_HASH
	KeypadHexadecimal = C.SDLK_KP_HEXADECIMAL
	KeypadLeftBrace   = C.SDLK_KP_LEFTBRACE
	KeypadLeftParen   = C.SDLK_KP_LEFTPAREN
	KeypadLess        = C.SDLK_KP_LESS
	KeypadMemAdd      = C.SDLK_KP_MEMADD
	KeypadMemClear    = C.SDLK_KP_MEMCLEAR
	KeypadMemDivide   = C.SDLK_KP_MEMDIVIDE
	KeypadMemMultiply = C.SDLK_KP_MEMMULTIPLY
	KeypadMemRecall   = C.SDLK_KP_MEMRECALL

	KeypadMemStore    = C.SDLK_KP_MEMSTORE
	KeypadMemSubtract = C.SDLK_KP_MEMSUBTRACT
	KeypadMinus       = C.SDLK_KP_MINUS
	KeypadMultiply    = C.SDLK_KP_MULTIPLY
	KeypadOctal       = C.SDLK_KP_OCTAL
	KeypadPercent     = C.SDLK_KP_PERCENT
	KeypadPeriod      = C.SDLK_KP_PERIOD
	KeypadPlus        = C.SDLK_KP_PLUS
	KeypadPlusMinus   = C.SDLK_KP_PLUSMINUS
	KeypadPower       = C.SDLK_KP_POWER

	KeypadRightBrace  = C.SDLK_KP_RIGHTBRACE
	KeypadRightParen  = C.SDLK_KP_RIGHTPAREN
	KeypadSpace       = C.SDLK_KP_SPACE
	KeypadTab         = C.SDLK_KP_TAB
	KeypadVerticalBar = C.SDLK_KP_VERTICALBAR
	KeypadXOR         = C.SDLK_KP_XOR
	KeyL              = C.SDLK_l
	KeyLeftAlt        = C.SDLK_LALT
	KeyLeftCtrl       = C.SDLK_LCTRL
	KeyLeft           = C.SDLK_LEFT

	KeyLeftBracket = C.SDLK_LEFTBRACKET
	KeyLeftGUI     = C.SDLK_LGUI
	KeyLeftShift   = C.SDLK_LSHIFT
	KeyM           = C.SDLK_m
	KeyMail        = C.SDLK_MAIL
	KeyMediaSelect = C.SDLK_MEDIASELECT
	KeyMenu        = C.SDLK_MENU
	KeyMinus       = C.SDLK_MINUS
	KeyModeSwitch  = C.SDLK_MODE
	KeyMute        = C.SDLK_MUTE

	KeyN        = C.SDLK_n
	KeyNumlock  = C.SDLK_NUMLOCKCLEAR
	KeyO        = C.SDLK_o
	KeyOper     = C.SDLK_OPER
	KeyOut      = C.SDLK_OUT
	KeyP        = C.SDLK_p
	KeyPageDown = C.SDLK_PAGEDOWN
	KeyPageUp   = C.SDLK_PAGEUP
	KeyPaste    = C.SDLK_PASTE
	KeyPause    = C.SDLK_PAUSE

	KeyPeriod      = C.SDLK_PERIOD
	KeyPower       = C.SDLK_POWER
	KeyPrintScreen = C.SDLK_PRINTSCREEN
	KeyPrior       = C.SDLK_PRIOR
	KeyQ           = C.SDLK_q
	KeyR           = C.SDLK_r
	KeyRightAlt    = C.SDLK_RALT
	KeyRightCtrl   = C.SDLK_RCTRL
	KeyReturn      = C.SDLK_RETURN
	KeyReturn      = C.SDLK_RETURN2

	KeyRightGUI     = C.SDLK_RGUI
	KeyRight        = C.SDLK_RIGHT
	KeyRightBracket = C.SDLK_RIGHTBRACKET
	KeyRightShift   = C.SDLK_RSHIFT
	KeyS            = C.SDLK_s
	KeyScrollLock   = C.SDLK_SCROLLLOCK
	KeySelect       = C.SDLK_SELECT
	KeySemicolon    = C.SDLK_SEMICOLON
	KeySeparator    = C.SDLK_SEPARATOR
	Key             = C.SDLK_SLASH

	KeySleep              = C.SDLK_SLEEP
	KeySpace              = C.SDLK_SPACE
	KeyStop               = C.SDLK_STOP
	KeySysReq             = C.SDLK_SYSREQ
	KeyT                  = C.SDLK_t
	KeyTab                = C.SDLK_TAB
	KeyThousandsSeparator = C.SDLK_THOUSANDSSEPARATOR
	KeyU                  = C.SDLK_u
	KeyUndo               = C.SDLK_UNDO
	KeyUnknown            = C.SDLK_UNKNOWN

	KeyUp         = C.SDLK_UP
	KeyV          = C.SDLK_v
	KeyVolumeDown = C.SDLK_VOLUMEDOWN
	KeyVolumeUp   = C.SDLK_VOLUMEUP
	KeyW          = C.SDLK_w
	KeyWWW        = C.SDLK_WWW
	KeyX          = C.SDLK_x
	KeyY          = C.SDLK_y
	KeyZ          = C.SDLK_z
	KeyAmpersand  = C.SDLK_AMPERSAND

	KeyMultiply  = C.SDLK_ASTERISK
	KeyAt        = C.SDLK_AT
	KeyPower     = C.SDLK_CARET
	KeyColon     = C.SDLK_COLON
	KeyDollar    = C.SDLK_DOLLAR
	KeyExclam    = C.SDLK_EXCLAIM
	KeyGreater   = C.SDLK_GREATER
	KeyHash      = C.SDLK_HASH
	KeyLeftParen = C.SDLK_LEFTPAREN
	KeyLess      = C.SDLK_LESS

	KeyPercent    = C.SDLK_PERCENT
	KeyPlus       = C.SDLK_PLUS
	KeyQuestion   = C.SDLK_QUESTION
	Key           = C.SDLK_QUOTEDBL
	KeyRightParen = C.SDLK_RIGHTPAREN
	KeyUnderscore = C.SDLK_UNDERSCORE
)
