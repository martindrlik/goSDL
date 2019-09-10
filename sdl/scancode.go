package sdl

// #include <SDL2/SDL_keycode.h>
import "C"

const (
	Scan0 = C.SDL_SCANCODE_0
	Scan1 = C.SDL_SCANCODE_1
	Scan2 = C.SDL_SCANCODE_2
	Scan3 = C.SDL_SCANCODE_3
	Scan4 = C.SDL_SCANCODE_4
	Scan5 = C.SDL_SCANCODE_5
	Scan6 = C.SDL_SCANCODE_6
	Scan7 = C.SDL_SCANCODE_7
	Scan8 = C.SDL_SCANCODE_8
	Scan9 = C.SDL_SCANCODE_9

	ScanA           = C.SDL_SCANCODE_A
	ScanACBack      = C.SDL_SCANCODE_AC_BACK
	ScanACBookmarks = C.SDL_SCANCODE_AC_BOOKMARKS
	ScanACForward   = C.SDL_SCANCODE_AC_FORWARD
	ScanACHome      = C.SDL_SCANCODE_AC_HOME
	ScanACRefresh   = C.SDL_SCANCODE_AC_REFRESH
	ScanACSearch    = C.SDL_SCANCODE_AC_SEARCH
	ScanACStop      = C.SDL_SCANCODE_AC_STOP
	ScanAgain       = C.SDL_SCANCODE_AGAIN
	ScanAltErase    = C.SDL_SCANCODE_ALTERASE

	ScanQuote       = C.SDL_SCANCODE_APOSTROPHE
	ScanApplication = C.SDL_SCANCODE_APPLICATION
	ScanAudioMute   = C.SDL_SCANCODE_AUDIOMUTE
	ScanAudioNext   = C.SDL_SCANCODE_AUDIONEXT
	ScanAudioPlay   = C.SDL_SCANCODE_AUDIOPLAY
	ScanAudioPrev   = C.SDL_SCANCODE_AUDIOPREV
	ScanAudioStop   = C.SDL_SCANCODE_AUDIOSTOP
	ScanB           = C.SDL_SCANCODE_B
	ScanBackslash   = C.SDL_SCANCODE_BACKSLASH
	ScanBackspace   = C.SDL_SCANCODE_BACKSPACE

	ScanBrightnessDown = C.SDL_SCANCODE_BRIGHTNESSDOWN
	ScanBrightnessUp   = C.SDL_SCANCODE_BRIGHTNESSUP
	ScanC              = C.SDL_SCANCODE_C
	ScanCalculator     = C.SDL_SCANCODE_CALCULATOR
	ScanCancel         = C.SDL_SCANCODE_CANCEL
	ScanCapsLock       = C.SDL_SCANCODE_CAPSLOCK
	ScanClear          = C.SDL_SCANCODE_CLEAR
	ScanClearAgain     = C.SDL_SCANCODE_CLEARAGAIN
	ScanComma          = C.SDL_SCANCODE_COMMA
	ScanComputer       = C.SDL_SCANCODE_COMPUTER

	ScanCopy             = C.SDL_SCANCODE_COPY
	ScanCrSel            = C.SDL_SCANCODE_CRSEL
	ScanCurrencySubUnit  = C.SDL_SCANCODE_CURRENCYSUBUNIT
	ScanCurrencyUnit     = C.SDL_SCANCODE_CURRENCYUNIT
	ScanCut              = C.SDL_SCANCODE_CUT
	ScanD                = C.SDL_SCANCODE_D
	ScanDecimalSeparator = C.SDL_SCANCODE_DECIMALSEPARATOR
	ScanDelete           = C.SDL_SCANCODE_DELETE
	ScanDisplaySwitch    = C.SDL_SCANCODE_DISPLAYSWITCH
	ScanDown             = C.SDL_SCANCODE_DOWN

	ScanE       = C.SDL_SCANCODE_E
	ScanEject   = C.SDL_SCANCODE_EJECT
	ScanEnd     = C.SDL_SCANCODE_END
	ScanEquals  = C.SDL_SCANCODE_EQUALS
	ScanEscape  = C.SDL_SCANCODE_ESCAPE
	ScanExecute = C.SDL_SCANCODE_EXECUTE
	ScanExSel   = C.SDL_SCANCODE_EXSEL
	ScanF       = C.SDL_SCANCODE_F
	ScanF1      = C.SDL_SCANCODE_F1
	ScanF10     = C.SDL_SCANCODE_F10

	ScanF11 = C.SDL_SCANCODE_F11
	ScanF12 = C.SDL_SCANCODE_F12
	ScanF13 = C.SDL_SCANCODE_F13
	ScanF14 = C.SDL_SCANCODE_F14
	ScanF15 = C.SDL_SCANCODE_F15
	ScanF16 = C.SDL_SCANCODE_F16
	ScanF17 = C.SDL_SCANCODE_F17
	ScanF18 = C.SDL_SCANCODE_F18
	ScanF19 = C.SDL_SCANCODE_F19
	ScanF2  = C.SDL_SCANCODE_F2

	ScanF20 = C.SDL_SCANCODE_F20
	ScanF21 = C.SDL_SCANCODE_F21
	ScanF22 = C.SDL_SCANCODE_F22
	ScanF23 = C.SDL_SCANCODE_F23
	ScanF24 = C.SDL_SCANCODE_F24
	ScanF3  = C.SDL_SCANCODE_F3
	ScanF4  = C.SDL_SCANCODE_F4
	ScanF5  = C.SDL_SCANCODE_F5
	ScanF6  = C.SDL_SCANCODE_F6
	ScanF7  = C.SDL_SCANCODE_F7

	ScanF8        = C.SDL_SCANCODE_F8
	ScanF9        = C.SDL_SCANCODE_F9
	ScanFind      = C.SDL_SCANCODE_FIND
	ScanG         = C.SDL_SCANCODE_G
	ScanBackQuote = C.SDL_SCANCODE_GRAVE
	ScanH         = C.SDL_SCANCODE_H
	ScanHelp      = C.SDL_SCANCODE_HELP
	ScanHome      = C.SDL_SCANCODE_HOME
	ScanI         = C.SDL_SCANCODE_I
	ScanInsert    = C.SDL_SCANCODE_INSERT

	ScanJ              = C.SDL_SCANCODE_J
	ScanK              = C.SDL_SCANCODE_K
	ScanKBDIllumDown   = C.SDL_SCANCODE_KBDILLUMDOWN
	ScanKBDIllumToggle = C.SDL_SCANCODE_KBDILLUMTOGGLE
	ScanKBDIllumUp     = C.SDL_SCANCODE_KBDILLUMUP
	ScanKeypad0        = C.SDL_SCANCODE_KP_0
	ScanKeypad00       = C.SDL_SCANCODE_KP_00
	ScanKeypad000      = C.SDL_SCANCODE_KP_000
	ScanKeypad1        = C.SDL_SCANCODE_KP_1
	ScanKeypad2        = C.SDL_SCANCODE_KP_2

	ScanKeypad3         = C.SDL_SCANCODE_KP_3
	ScanKeypad4         = C.SDL_SCANCODE_KP_4
	ScanKeypad5         = C.SDL_SCANCODE_KP_5
	ScanKeypad6         = C.SDL_SCANCODE_KP_6
	ScanKeypad7         = C.SDL_SCANCODE_KP_7
	ScanKeypad8         = C.SDL_SCANCODE_KP_8
	ScanKeypad9         = C.SDL_SCANCODE_KP_9
	ScanKeypadA         = C.SDL_SCANCODE_KP_A
	ScanKeypadAmpersand = C.SDL_SCANCODE_KP_AMPERSAND
	ScanKeypadAt        = C.SDL_SCANCODE_KP_AT

	ScanKeypadB            = C.SDL_SCANCODE_KP_B
	ScanKeypadBackspace    = C.SDL_SCANCODE_KP_BACKSPACE
	ScanKeypadBinary       = C.SDL_SCANCODE_KP_BINARY
	ScanKeypadC            = C.SDL_SCANCODE_KP_C
	ScanKeypadClear        = C.SDL_SCANCODE_KP_CLEAR
	ScanKeypadClearEntry   = C.SDL_SCANCODE_KP_CLEARENTRY
	ScanKeypadColon        = C.SDL_SCANCODE_KP_COLON
	ScanKeypadComma        = C.SDL_SCANCODE_KP_COMMA
	ScanKeypadD            = C.SDL_SCANCODE_KP_D
	ScanKeypadDblAmpersand = C.SDL_SCANCODE_KP_DBLAMPERSAND

	ScanKeypadDblVerticalBar = C.SDL_SCANCODE_KP_DBLVERTICALBAR
	ScanKeypadDecimal        = C.SDL_SCANCODE_KP_DECIMAL
	ScanKeypad               = C.SDL_SCANCODE_KP_DIVIDE
	ScanKeypadE              = C.SDL_SCANCODE_KP_E
	ScanKeypadEnter          = C.SDL_SCANCODE_KP_ENTER
	ScanKeypadEquals         = C.SDL_SCANCODE_KP_EQUALS
	ScanKeypadEqualsAS400    = C.SDL_SCANCODE_KP_EQUALSAS400
	ScanKeypadExclam         = C.SDL_SCANCODE_KP_EXCLAM
	ScanKeypadF              = C.SDL_SCANCODE_KP_F
	ScanKeypadGreater        = C.SDL_SCANCODE_KP_GREATER

	ScanKeypadHash        = C.SDL_SCANCODE_KP_HASH
	ScanKeypadHexadecimal = C.SDL_SCANCODE_KP_HEXADECIMAL
	ScanKeypadLeftBrace   = C.SDL_SCANCODE_KP_LEFTBRACE
	ScanKeypadLeftParen   = C.SDL_SCANCODE_KP_LEFTPAREN
	ScanKeypadLess        = C.SDL_SCANCODE_KP_LESS
	ScanKeypadMemAdd      = C.SDL_SCANCODE_KP_MEMADD
	ScanKeypadMemClear    = C.SDL_SCANCODE_KP_MEMCLEAR
	ScanKeypadMemDivide   = C.SDL_SCANCODE_KP_MEMDIVIDE
	ScanKeypadMemMultiply = C.SDL_SCANCODE_KP_MEMMULTIPLY
	ScanKeypadMemRecall   = C.SDL_SCANCODE_KP_MEMRECALL

	ScanKeypadMemStore    = C.SDL_SCANCODE_KP_MEMSTORE
	ScanKeypadMemSubtract = C.SDL_SCANCODE_KP_MEMSUBTRACT
	ScanKeypadMinus       = C.SDL_SCANCODE_KP_MINUS
	ScanKeypadMultiply    = C.SDL_SCANCODE_KP_MULTIPLY
	ScanKeypadOctal       = C.SDL_SCANCODE_KP_OCTAL
	ScanKeypadPercent     = C.SDL_SCANCODE_KP_PERCENT
	ScanKeypadPeriod      = C.SDL_SCANCODE_KP_PERIOD
	ScanKeypadPlus        = C.SDL_SCANCODE_KP_PLUS
	ScanKeypadPlusMinus   = C.SDL_SCANCODE_KP_PLUSMINUS
	ScanKeypadPower       = C.SDL_SCANCODE_KP_POWER

	ScanKeypadRightBrace  = C.SDL_SCANCODE_KP_RIGHTBRACE
	ScanKeypadRightParen  = C.SDL_SCANCODE_KP_RIGHTPAREN
	ScanKeypadSpace       = C.SDL_SCANCODE_KP_SPACE
	ScanKeypadTab         = C.SDL_SCANCODE_KP_TAB
	ScanKeypadVerticalBar = C.SDL_SCANCODE_KP_VERTICALBAR
	ScanKeypadXOR         = C.SDL_SCANCODE_KP_XOR
	ScanL                 = C.SDL_SCANCODE_L
	ScanLeftAlt           = C.SDL_SCANCODE_LALT
	ScanLeftCtrl          = C.SDL_SCANCODE_LCTRL
	ScanLeft              = C.SDL_SCANCODE_LEFT

	ScanLeftBracket = C.SDL_SCANCODE_LEFTBRACKET
	ScanLeftGUI     = C.SDL_SCANCODE_LGUI
	ScanLeftShift   = C.SDL_SCANCODE_LSHIFT
	ScanM           = C.SDL_SCANCODE_M
	ScanMail        = C.SDL_SCANCODE_MAIL
	ScanMediaSelect = C.SDL_SCANCODE_MEDIASELECT
	ScanMenu        = C.SDL_SCANCODE_MENU
	ScanMinus       = C.SDL_SCANCODE_MINUS
	ScanModeSwitch  = C.SDL_SCANCODE_MODE
	ScanMute        = C.SDL_SCANCODE_MUTE

	ScanN        = C.SDL_SCANCODE_N
	ScanNumlock  = C.SDL_SCANCODE_NUMLOCKCLEAR
	ScanO        = C.SDL_SCANCODE_O
	ScanOper     = C.SDL_SCANCODE_OPER
	ScanOut      = C.SDL_SCANCODE_OUT
	ScanP        = C.SDL_SCANCODE_P
	ScanPageDown = C.SDL_SCANCODE_PAGEDOWN
	ScanPageUp   = C.SDL_SCANCODE_PAGEUP
	ScanPaste    = C.SDL_SCANCODE_PASTE
	ScanPause    = C.SDL_SCANCODE_PAUSE

	ScanPeriod      = C.SDL_SCANCODE_PERIOD
	ScanPower       = C.SDL_SCANCODE_POWER
	ScanPrintScreen = C.SDL_SCANCODE_PRINTSCREEN
	ScanPrior       = C.SDL_SCANCODE_PRIOR
	ScanQ           = C.SDL_SCANCODE_Q
	ScanR           = C.SDL_SCANCODE_R
	ScanRightAlt    = C.SDL_SCANCODE_RALT
	ScanRightCtrl   = C.SDL_SCANCODE_RCTRL
	ScanReturn      = C.SDL_SCANCODE_RETURN
	ScanReturn      = C.SDL_SCANCODE_RETURN2

	ScanRightGUI     = C.SDL_SCANCODE_RGUI
	ScanRight        = C.SDL_SCANCODE_RIGHT
	ScanRightBracket = C.SDL_SCANCODE_RIGHTBRACKET
	ScanRightShift   = C.SDL_SCANCODE_RSHIFT
	ScanS            = C.SDL_SCANCODE_S
	ScanScrollLock   = C.SDL_SCANCODE_SCROLLLOCK
	ScanSelect       = C.SDL_SCANCODE_SELECT
	ScanSemicolon    = C.SDL_SCANCODE_SEMICOLON
	ScanSeparator    = C.SDL_SCANCODE_SEPARATOR
	Scan             = C.SDL_SCANCODE_SLASH

	ScanSleep              = C.SDL_SCANCODE_SLEEP
	ScanSpace              = C.SDL_SCANCODE_SPACE
	ScanStop               = C.SDL_SCANCODE_STOP
	ScanSysReq             = C.SDL_SCANCODE_SYSREQ
	ScanT                  = C.SDL_SCANCODE_T
	ScanTab                = C.SDL_SCANCODE_TAB
	ScanThousandsSeparator = C.SDL_SCANCODE_THOUSANDSSEPARATOR
	ScanU                  = C.SDL_SCANCODE_U
	ScanUndo               = C.SDL_SCANCODE_UNDO
	ScanUnknown            = C.SDL_SCANCODE_UNKNOWN

	ScanUp             = C.SDL_SCANCODE_UP
	ScanV              = C.SDL_SCANCODE_V
	ScanVolumeDown     = C.SDL_SCANCODE_VOLUMEDOWN
	ScanVolumeUp       = C.SDL_SCANCODE_VOLUMEUP
	ScanW              = C.SDL_SCANCODE_W
	ScanWWW            = C.SDL_SCANCODE_WWW
	ScanX              = C.SDL_SCANCODE_X
	ScanY              = C.SDL_SCANCODE_Y
	ScanZ              = C.SDL_SCANCODE_Z
	ScanInternational1 = C.SDL_SCANCODE_INTERNATIONAL1

	ScanInternational2 = C.SDL_SCANCODE_INTERNATIONAL2
	ScanInternational3 = C.SDL_SCANCODE_INTERNATIONAL3
	ScanInternational4 = C.SDL_SCANCODE_INTERNATIONAL4
	ScanInternational5 = C.SDL_SCANCODE_INTERNATIONAL5
	ScanInternational6 = C.SDL_SCANCODE_INTERNATIONAL6
	ScanInternational7 = C.SDL_SCANCODE_INTERNATIONAL7
	ScanInternational8 = C.SDL_SCANCODE_INTERNATIONAL8
	ScanInternational9 = C.SDL_SCANCODE_INTERNATIONAL9
	ScanLang1          = C.SDL_SCANCODE_LANG1
	ScanLang2          = C.SDL_SCANCODE_LANG2

	ScanLang3             = C.SDL_SCANCODE_LANG3
	ScanLang4             = C.SDL_SCANCODE_LANG4
	ScanLang5             = C.SDL_SCANCODE_LANG5
	ScanLang6             = C.SDL_SCANCODE_LANG6
	ScanLang7             = C.SDL_SCANCODE_LANG7
	ScanLang8             = C.SDL_SCANCODE_LANG8
	ScanLang9             = C.SDL_SCANCODE_LANG9
	ScanLockingCapsLock   = C.SDL_SCANCODE_LOCKINGCAPSLOCK
	ScanLockingNumLock    = C.SDL_SCANCODE_LOCKINGNUMLOCK
	ScanLockingScrollLock = C.SDL_SCANCODE_LOCKINGSCROLLLOCK

	ScanNonUSBackslash = C.SDL_SCANCODE_NONUSBACKSLASH
	ScanNonUSHash      = C.SDL_SCANCODE_NONUSHASH
)
