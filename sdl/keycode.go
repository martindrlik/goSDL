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
	KeyA = C.SDLK_a
	// KeyACBack "AC Back" (the Back key (application control keypad))
	KeyACBack = C.SDLK_AC_BACK
	// KeyACBookmarks "AC Bookmarks" (the Bookmarks key (application control keypad))
	KeyACBookmarks = C.SDLK_AC_BOOKMARKS
	// KeyACForward "AC Forward" (the Forward key (application control keypad))
	KeyACForward = C.SDLK_AC_FORWARD
	// KeyACHome "AC Home" (the Home key (application control keypad))
	KeyACHome = C.SDLK_AC_HOME
	// KeyACRefresh "AC Refresh" (the Refresh key (application control keypad))
	KeyACRefresh = C.SDLK_AC_REFRESH
	// KeyACSearch "AC Search" (the Search key (application control keypad))
	KeyACSearch = C.SDLK_AC_SEARCH
	// KeyACStop "AC Stop" (the Stop key (application control keypad))
	KeyACStop = C.SDLK_AC_STOP
	// KeyAgain "Again" (the Again key (Redo))
	KeyAgain = C.SDLK_AGAIN
	// KeyAltErase "AltErase" (Erase-Eaze)
	KeyAltErase = C.SDLK_ALTERASE

	KeyQuote = C.SDLK_QUOTE
	// KeyApplication "Application" (the Application / Compose / Context Menu (Windows) key)
	KeyApplication = C.SDLK_APPLICATION
	// KeyAudioMute "AudioMute" (the Mute volume key)
	KeyAudioMute = C.SDLK_AUDIOMUTE
	// KeyAudioNext "AudioNext" (the Next Track media key)
	KeyAudioNext = C.SDLK_AUDIONEXT
	// KeyAudioPlay "AudioPlay" (the Play media key)
	KeyAudioPlay = C.SDLK_AUDIOPLAY
	// KeyAudioPrev "AudioPrev" (the Previous Track media key)
	KeyAudioPrev = C.SDLK_AUDIOPREV
	// KeyAudioStop "AudioStop" (the Stop media key)
	KeyAudioStop = C.SDLK_AUDIOSTOP

	KeyB = C.SDLK_b
	// KeyBackslash "\" (Located at the lower left of the return key on ISO keyboards and at the right end of the QWERTY row on ANSI keyboards. Produces REVERSE SOLIDUS (backslash) and VERTICAL LINE in a US layout, REVERSE SOLIDUS and VERTICAL LINE in a UK Mac layout, NUMBER SIGN and TILDE in a UK Windows layout, DOLLAR SIGN and POUND SIGN in a Swiss German layout, NUMBER SIGN and APOSTROPHE in a German layout, GRAVE ACCENT and POUND SIGN in a French Mac layout, and ASTERISK and MICRO SIGN in a French Windows layout.)
	KeyBackslash = C.SDLK_BACKSLASH

	KeyBackspace = C.SDLK_BACKSPACE
	// KeyBrightnessDown "BrightnessDown" (the Brightness Down key)
	KeyBrightnessDown = C.SDLK_BRIGHTNESSDOWN
	// KeyBrightnessUp "BrightnessUp" (the Brightness Up key)
	KeyBrightnessUp = C.SDLK_BRIGHTNESSUP

	KeyC = C.SDLK_c
	// KeyCalculator "Calculator" (the Calculator key)
	KeyCalculator = C.SDLK_CALCULATOR

	KeyCancel     = C.SDLK_CANCEL
	KeyCapsLock   = C.SDLK_CAPSLOCK
	KeyClear      = C.SDLK_CLEAR
	KeyClearAgain = C.SDLK_CLEARAGAIN
	KeyComma      = C.SDLK_COMMA
	// KeyComputer "Computer" (the My Computer key)
	KeyComputer = C.SDLK_COMPUTER

	KeyCopy  = C.SDLK_COPY
	KeyCrSel = C.SDLK_CRSEL
	// KeyCurrencySubUnit "CurrencySubUnit" (the Currency Subunit key)
	KeyCurrencySubUnit = C.SDLK_CURRENCYSUBUNIT
	// KeyCurrencyUnit "CurrencyUnit" (the Currency Unit key)
	KeyCurrencyUnit = C.SDLK_CURRENCYUNIT

	KeyCut = C.SDLK_CUT
	KeyD   = C.SDLK_d
	// KeyDecimalSeparator "DecimalSeparator" (the Decimal Separator key)
	KeyDecimalSeparator = C.SDLK_DECIMALSEPARATOR

	KeyDelete = C.SDLK_DELETE
	// KeyDisplaySwitch "DisplaySwitch" (display mirroring/dual display switch, video mode switch)
	KeyDisplaySwitch = C.SDLK_DISPLAYSWITCH
	// KeyDown "Down" (the Down arrow key (navigation keypad))
	KeyDown = C.SDLK_DOWN

	KeyE = C.SDLK_e
	// KeyEject "Eject" (the Eject key)
	KeyEject = C.SDLK_EJECT

	KeyEnd    = C.SDLK_END
	KeyEquals = C.SDLK_EQUALS
	// KeyEscape "Escape" (the Esc key)
	KeyEscape = C.SDLK_ESCAPE

	KeyExecute = C.SDLK_EXECUTE
	KeyExSel   = C.SDLK_EXSEL
	KeyF       = C.SDLK_f
	KeyF1      = C.SDLK_F1
	KeyF10     = C.SDLK_F10
	KeyF11     = C.SDLK_F11
	KeyF12     = C.SDLK_F12
	KeyF13     = C.SDLK_F13
	KeyF14     = C.SDLK_F14
	KeyF15     = C.SDLK_F15
	KeyF16     = C.SDLK_F16
	KeyF17     = C.SDLK_F17
	KeyF18     = C.SDLK_F18
	KeyF19     = C.SDLK_F19
	KeyF2      = C.SDLK_F2
	KeyF20     = C.SDLK_F20
	KeyF21     = C.SDLK_F21
	KeyF22     = C.SDLK_F22
	KeyF23     = C.SDLK_F23
	KeyF24     = C.SDLK_F24
	KeyF3      = C.SDLK_F3
	KeyF4      = C.SDLK_F4
	KeyF5      = C.SDLK_F5
	KeyF6      = C.SDLK_F6
	KeyF7      = C.SDLK_F7
	KeyF8      = C.SDLK_F8
	KeyF9      = C.SDLK_F9
	KeyFind    = C.SDLK_FIND
	KeyG       = C.SDLK_g
	// KeyBackQuote "`" (Located in the top left corner (on both ANSI and ISO keyboards). Produces GRAVE ACCENT and TILDE in a US Windows layout and in US and UK Mac layouts on ANSI keyboards, GRAVE ACCENT and NOT SIGN in a UK Windows layout, SECTION SIGN and PLUS-MINUS SIGN in US and UK Mac layouts on ISO keyboards, SECTION SIGN and DEGREE SIGN in a Swiss German layout (Mac: only on ISO keyboards), CIRCUMFLEX ACCENT and DEGREE SIGN in a German layout (Mac: only on ISO keyboards), SUPERSCRIPT TWO and TILDE in a French Windows layout, COMMERCIAL AT and NUMBER SIGN in a French Mac layout on ISO keyboards, and LESS-THAN SIGN and GREATER-THAN SIGN in a Swiss German, German, or French Mac layout on ANSI keyboards.)
	KeyBackQuote = C.SDLK_BACKQUOTE

	KeyH    = C.SDLK_h
	KeyHelp = C.SDLK_HELP
	KeyHome = C.SDLK_HOME
	KeyI    = C.SDLK_i
	// KeyInsert "Insert" (insert on PC, help on some Mac keyboards (but does send code 73, not 117))
	KeyInsert = C.SDLK_INSERT

	KeyJ = C.SDLK_j
	KeyK = C.SDLK_k
	// KeyKBDIllumDown "KBDIllumDown" (the Keyboard Illumination Down key)
	KeyKBDIllumDown = C.SDLK_KBDILLUMDOWN
	// KeyKBDIllumToggle "KBDIllumToggle" (the Keyboard Illumination Toggle key)
	KeyKBDIllumToggle = C.SDLK_KBDILLUMTOGGLE
	// KeyKBDIllumUp "KBDIllumUp" (the Keyboard Illumination Up key)
	KeyKBDIllumUp = C.SDLK_KBDILLUMUP
	// Keypad0 "Keypad 0" (the 0 key (numeric keypad))
	Keypad0 = C.SDLK_KP_0
	// Keypad00 "Keypad 00" (the 00 key (numeric keypad))
	Keypad00 = C.SDLK_KP_00
	// Keypad000 "Keypad 000" (the 000 key (numeric keypad))
	Keypad000 = C.SDLK_KP_000
	// Keypad1 "Keypad 1" (the 1 key (numeric keypad))
	Keypad1 = C.SDLK_KP_1
	// Keypad2 "Keypad 2" (the 2 key (numeric keypad))
	Keypad2 = C.SDLK_KP_2
	// Keypad3 "Keypad 3" (the 3 key (numeric keypad))
	Keypad3 = C.SDLK_KP_3
	// Keypad4 "Keypad 4" (the 4 key (numeric keypad))
	Keypad4 = C.SDLK_KP_4
	// Keypad5 "Keypad 5" (the 5 key (numeric keypad))
	Keypad5 = C.SDLK_KP_5
	// Keypad6 "Keypad 6" (the 6 key (numeric keypad))
	Keypad6 = C.SDLK_KP_6
	// Keypad7 "Keypad 7" (the 7 key (numeric keypad))
	Keypad7 = C.SDLK_KP_7
	// Keypad8 "Keypad 8" (the 8 key (numeric keypad))
	Keypad8 = C.SDLK_KP_8
	// Keypad9 "Keypad 9" (the 9 key (numeric keypad))
	Keypad9 = C.SDLK_KP_9
	// KeypadA "Keypad A" (the A key (numeric keypad))
	KeypadA = C.SDLK_KP_A
	// KeypadAmpersand "Keypad &" (the & key (numeric keypad))
	KeypadAmpersand = C.SDLK_KP_AMPERSAND
	// KeypadAt "Keypad @" (the @ key (numeric keypad))
	KeypadAt = C.SDLK_KP_AT
	// KeypadB "Keypad B" (the B key (numeric keypad))
	KeypadB = C.SDLK_KP_B
	// KeypadBackspace "Keypad Backspace" (the Backspace key (numeric keypad))
	KeypadBackspace = C.SDLK_KP_BACKSPACE
	// KeypadBinary "Keypad Binary" (the Binary key (numeric keypad))
	KeypadBinary = C.SDLK_KP_BINARY
	// KeypadC "Keypad C" (the C key (numeric keypad))
	KeypadC = C.SDLK_KP_C
	// KeypadClear "Keypad Clear" (the Clear key (numeric keypad))
	KeypadClear = C.SDLK_KP_CLEAR
	// KeypadClearEntry "Keypad ClearEntry" (the Clear Entry key (numeric keypad))
	KeypadClearEntry = C.SDLK_KP_CLEARENTRY
	// KeypadColon "Keypad :" (the : key (numeric keypad))
	KeypadColon = C.SDLK_KP_COLON
	// KeypadComma "Keypad ," (the Comma key (numeric keypad))
	KeypadComma = C.SDLK_KP_COMMA
	// KeypadD "Keypad D" (the D key (numeric keypad))
	KeypadD = C.SDLK_KP_D
	// KeypadDblAmpersand "Keypad &&" (the && key (numeric keypad))
	KeypadDblAmpersand = C.SDLK_KP_DBLAMPERSAND
	// KeypadDblVerticalBar "Keypad ||" (the || key (numeric keypad))
	KeypadDblVerticalBar = C.SDLK_KP_DBLVERTICALBAR
	// KeypadDecimal "Keypad Decimal" (the Decimal key (numeric keypad))
	KeypadDecimal = C.SDLK_KP_DECIMAL
	// Keypad "Keypad /" (the / key (numeric keypad))
	Keypad = C.SDLK_KP_DIVIDE
	// KeypadE "Keypad E" (the E key (numeric keypad))
	KeypadE = C.SDLK_KP_E
	// KeypadEnter "Keypad Enter" (the Enter key (numeric keypad))
	KeypadEnter = C.SDLK_KP_ENTER
	// KeypadEquals "Keypad =" (the = key (numeric keypad))
	KeypadEquals = C.SDLK_KP_EQUALS
	// KeypadEqualsAS400 "Keypad = (AS400)" (the Equals AS400 key (numeric keypad))
	KeypadEqualsAS400 = C.SDLK_KP_EQUALSAS400
	// KeypadExclam "Keypad !" (the ! key (numeric keypad))
	KeypadExclam = C.SDLK_KP_EXCLAM
	// KeypadF "Keypad F" (the F key (numeric keypad))
	KeypadF = C.SDLK_KP_F
	// KeypadGreater "Keypad >" (the Greater key (numeric keypad))
	KeypadGreater = C.SDLK_KP_GREATER
	// KeypadHash "Keypad #" (the # key (numeric keypad))
	KeypadHash = C.SDLK_KP_HASH
	// KeypadHexadecimal "Keypad Hexadecimal" (the Hexadecimal key (numeric keypad))
	KeypadHexadecimal = C.SDLK_KP_HEXADECIMAL
	// KeypadLeftBrace "Keypad {" (the Left Brace key (numeric keypad))
	KeypadLeftBrace = C.SDLK_KP_LEFTBRACE
	// KeypadLeftParen "Keypad (" (the Left Parenthesis key (numeric keypad))
	KeypadLeftParen = C.SDLK_KP_LEFTPAREN
	// KeypadLess "Keypad <" (the Less key (numeric keypad))
	KeypadLess = C.SDLK_KP_LESS
	// KeypadMemAdd "Keypad MemAdd" (the Mem Add key (numeric keypad))
	KeypadMemAdd = C.SDLK_KP_MEMADD
	// KeypadMemClear "Keypad MemClear" (the Mem Clear key (numeric keypad))
	KeypadMemClear = C.SDLK_KP_MEMCLEAR
	// KeypadMemDivide "Keypad MemDivide" (the Mem Divide key (numeric keypad))
	KeypadMemDivide = C.SDLK_KP_MEMDIVIDE
	// KeypadMemMultiply "Keypad MemMultiply" (the Mem Multiply key (numeric keypad))
	KeypadMemMultiply = C.SDLK_KP_MEMMULTIPLY
	// KeypadMemRecall "Keypad MemRecall" (the Mem Recall key (numeric keypad))
	KeypadMemRecall = C.SDLK_KP_MEMRECALL
	// KeypadMemStore "Keypad MemStore" (the Mem Store key (numeric keypad))
	KeypadMemStore = C.SDLK_KP_MEMSTORE
	// KeypadMemSubtract "Keypad MemSubtract" (the Mem Subtract key (numeric keypad))
	KeypadMemSubtract = C.SDLK_KP_MEMSUBTRACT
	// KeypadMinus "Keypad -" (the - key (numeric keypad))
	KeypadMinus = C.SDLK_KP_MINUS
	// KeypadMultiply "Keypad *" (the * key (numeric keypad))
	KeypadMultiply = C.SDLK_KP_MULTIPLY
	// KeypadOctal "Keypad Octal" (the Octal key (numeric keypad))
	KeypadOctal = C.SDLK_KP_OCTAL
	// KeypadPercent "Keypad %" (the Percent key (numeric keypad))
	KeypadPercent = C.SDLK_KP_PERCENT
	// KeypadPeriod "Keypad ." (the . key (numeric keypad))
	KeypadPeriod = C.SDLK_KP_PERIOD
	// KeypadPlus "Keypad +" (the + key (numeric keypad))
	KeypadPlus = C.SDLK_KP_PLUS
	// KeypadPlusMinus "Keypad +/-" (the +/- key (numeric keypad))
	KeypadPlusMinus = C.SDLK_KP_PLUSMINUS
	// KeypadPower "Keypad ^" (the Power key (numeric keypad))
	KeypadPower = C.SDLK_KP_POWER
	// KeypadRightBrace "Keypad }" (the Right Brace key (numeric keypad))
	KeypadRightBrace = C.SDLK_KP_RIGHTBRACE
	// KeypadRightParen "Keypad )" (the Right Parenthesis key (numeric keypad))
	KeypadRightParen = C.SDLK_KP_RIGHTPAREN
	// KeypadSpace "Keypad Space" (the Space key (numeric keypad))
	KeypadSpace = C.SDLK_KP_SPACE
	// KeypadTab "Keypad Tab" (the Tab key (numeric keypad))
	KeypadTab = C.SDLK_KP_TAB
	// KeypadVerticalBar "Keypad |" (the | key (numeric keypad))
	KeypadVerticalBar = C.SDLK_KP_VERTICALBAR
	// KeypadXOR "Keypad XOR" (the XOR key (numeric keypad))
	KeypadXOR = C.SDLK_KP_XOR

	KeyL = C.SDLK_l
	// KeyLeftAlt "Left Alt" (alt, option)
	KeyLeftAlt = C.SDLK_LALT

	KeyLeftCtrl = C.SDLK_LCTRL
	// KeyLeft "Left" (the Left arrow key (navigation keypad))
	KeyLeft = C.SDLK_LEFT

	KeyLeftBracket = C.SDLK_LEFTBRACKET
	// KeyLeftGUI "Left GUI" (windows, command (apple), meta)
	KeyLeftGUI = C.SDLK_LGUI

	KeyLeftShift = C.SDLK_LSHIFT
	KeyM         = C.SDLK_m
	// KeyMail "Mail" (the Mail/eMail key)
	KeyMail = C.SDLK_MAIL
	// KeyMediaSelect "MediaSelect" (the Media Select key)
	KeyMediaSelect = C.SDLK_MEDIASELECT

	KeyMenu  = C.SDLK_MENU
	KeyMinus = C.SDLK_MINUS
	// KeyModeSwitch "ModeSwitch" (I'm not sure if this is really not covered by any of the above, but since there's a special KMOD_MODE for it I'm adding it here)
	KeyModeSwitch = C.SDLK_MODE

	KeyMute = C.SDLK_MUTE
	KeyN    = C.SDLK_n
	// KeyNumlock "Numlock" (the Num Lock key (PC) / the Clear key (Mac))
	KeyNumlock = C.SDLK_NUMLOCKCLEAR

	KeyO        = C.SDLK_o
	KeyOper     = C.SDLK_OPER
	KeyOut      = C.SDLK_OUT
	KeyP        = C.SDLK_p
	KeyPageDown = C.SDLK_PAGEDOWN
	KeyPageUp   = C.SDLK_PAGEUP
	KeyPaste    = C.SDLK_PASTE
	// KeyPause "Pause" (the Pause / Break key)
	KeyPause = C.SDLK_PAUSE

	KeyPeriod = C.SDLK_PERIOD
	// KeyPower "Power" (The USB document says this is a status flag, not a physical key - but some Mac keyboards do have a power key.)
	KeyPower = C.SDLK_POWER

	KeyPrintScreen = C.SDLK_PRINTSCREEN
	KeyPrior       = C.SDLK_PRIOR
	KeyQ           = C.SDLK_q
	KeyR           = C.SDLK_r
	// KeyRightAlt "Right Alt" (alt gr, option)
	KeyRightAlt = C.SDLK_RALT

	KeyRightCtrl = C.SDLK_RCTRL
	// KeyReturn "Return" (the Enter key (main keyboard))
	KeyReturn = C.SDLK_RETURN

	KeyReturn = C.SDLK_RETURN2
	// KeyRightGUI "Right GUI" (windows, command (apple), meta)
	KeyRightGUI = C.SDLK_RGUI
	// KeyRight "Right" (the Right arrow key (navigation keypad))
	KeyRight = C.SDLK_RIGHT

	KeyRightBracket = C.SDLK_RIGHTBRACKET
	KeyRightShift   = C.SDLK_RSHIFT
	KeyS            = C.SDLK_s
	KeyScrollLock   = C.SDLK_SCROLLLOCK
	KeySelect       = C.SDLK_SELECT
	KeySemicolon    = C.SDLK_SEMICOLON
	KeySeparator    = C.SDLK_SEPARATOR
	Key             = C.SDLK_SLASH
	// KeySleep "Sleep" (the Sleep key)
	KeySleep = C.SDLK_SLEEP
	// KeySpace "Space" (the Space Bar key(s))
	KeySpace = C.SDLK_SPACE

	KeyStop = C.SDLK_STOP
	// KeySysReq "SysReq" (the SysReq key)
	KeySysReq = C.SDLK_SYSREQ

	KeyT = C.SDLK_t
	// KeyTab "Tab" (the Tab key)
	KeyTab = C.SDLK_TAB
	// KeyThousandsSeparator "ThousandsSeparator" (the Thousands Separator key)
	KeyThousandsSeparator = C.SDLK_THOUSANDSSEPARATOR

	KeyU    = C.SDLK_u
	KeyUndo = C.SDLK_UNDO
	// KeyUnknown "" (no name, empty string)
	KeyUnknown = C.SDLK_UNKNOWN
	// KeyUp "Up" (the Up arrow key (navigation keypad))
	KeyUp = C.SDLK_UP

	KeyV          = C.SDLK_v
	KeyVolumeDown = C.SDLK_VOLUMEDOWN
	KeyVolumeUp   = C.SDLK_VOLUMEUP
	KeyW          = C.SDLK_w
	// KeyWWW "WWW" (the WWW/World Wide Web key)
	KeyWWW = C.SDLK_WWW

	KeyX          = C.SDLK_x
	KeyY          = C.SDLK_y
	KeyZ          = C.SDLK_z
	KeyAmpersand  = C.SDLK_AMPERSAND
	KeyMultiply   = C.SDLK_ASTERISK
	KeyAt         = C.SDLK_AT
	KeyPower      = C.SDLK_CARET
	KeyColon      = C.SDLK_COLON
	KeyDollar     = C.SDLK_DOLLAR
	KeyExclam     = C.SDLK_EXCLAIM
	KeyGreater    = C.SDLK_GREATER
	KeyHash       = C.SDLK_HASH
	KeyLeftParen  = C.SDLK_LEFTPAREN
	KeyLess       = C.SDLK_LESS
	KeyPercent    = C.SDLK_PERCENT
	KeyPlus       = C.SDLK_PLUS
	KeyQuestion   = C.SDLK_QUESTION
	Key           = C.SDLK_QUOTEDBL
	KeyRightParen = C.SDLK_RIGHTPAREN
	KeyUnderscore = C.SDLK_UNDERSCORE
)
