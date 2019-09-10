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
	ScanA = C.SDL_SCANCODE_A
	// ScanACBack "AC Back" (the Back key (application control keypad))
	ScanACBack = C.SDL_SCANCODE_AC_BACK
	// ScanACBookmarks "AC Bookmarks" (the Bookmarks key (application control keypad))
	ScanACBookmarks = C.SDL_SCANCODE_AC_BOOKMARKS
	// ScanACForward "AC Forward" (the Forward key (application control keypad))
	ScanACForward = C.SDL_SCANCODE_AC_FORWARD
	// ScanACHome "AC Home" (the Home key (application control keypad))
	ScanACHome = C.SDL_SCANCODE_AC_HOME
	// ScanACRefresh "AC Refresh" (the Refresh key (application control keypad))
	ScanACRefresh = C.SDL_SCANCODE_AC_REFRESH
	// ScanACSearch "AC Search" (the Search key (application control keypad))
	ScanACSearch = C.SDL_SCANCODE_AC_SEARCH
	// ScanACStop "AC Stop" (the Stop key (application control keypad))
	ScanACStop = C.SDL_SCANCODE_AC_STOP
	// ScanAgain "Again" (the Again key (Redo))
	ScanAgain = C.SDL_SCANCODE_AGAIN
	// ScanAltErase "AltErase" (Erase-Eaze)
	ScanAltErase = C.SDL_SCANCODE_ALTERASE

	ScanQuote = C.SDL_SCANCODE_APOSTROPHE
	// ScanApplication "Application" (the Application / Compose / Context Menu (Windows) key)
	ScanApplication = C.SDL_SCANCODE_APPLICATION
	// ScanAudioMute "AudioMute" (the Mute volume key)
	ScanAudioMute = C.SDL_SCANCODE_AUDIOMUTE
	// ScanAudioNext "AudioNext" (the Next Track media key)
	ScanAudioNext = C.SDL_SCANCODE_AUDIONEXT
	// ScanAudioPlay "AudioPlay" (the Play media key)
	ScanAudioPlay = C.SDL_SCANCODE_AUDIOPLAY
	// ScanAudioPrev "AudioPrev" (the Previous Track media key)
	ScanAudioPrev = C.SDL_SCANCODE_AUDIOPREV
	// ScanAudioStop "AudioStop" (the Stop media key)
	ScanAudioStop = C.SDL_SCANCODE_AUDIOSTOP

	ScanB = C.SDL_SCANCODE_B
	// ScanBackslash "\" (Located at the lower left of the return key on ISO keyboards and at the right end of the QWERTY row on ANSI keyboards. Produces REVERSE SOLIDUS (backslash) and VERTICAL LINE in a US layout, REVERSE SOLIDUS and VERTICAL LINE in a UK Mac layout, NUMBER SIGN and TILDE in a UK Windows layout, DOLLAR SIGN and POUND SIGN in a Swiss German layout, NUMBER SIGN and APOSTROPHE in a German layout, GRAVE ACCENT and POUND SIGN in a French Mac layout, and ASTERISK and MICRO SIGN in a French Windows layout.)
	ScanBackslash = C.SDL_SCANCODE_BACKSLASH

	ScanBackspace = C.SDL_SCANCODE_BACKSPACE
	// ScanBrightnessDown "BrightnessDown" (the Brightness Down key)
	ScanBrightnessDown = C.SDL_SCANCODE_BRIGHTNESSDOWN
	// ScanBrightnessUp "BrightnessUp" (the Brightness Up key)
	ScanBrightnessUp = C.SDL_SCANCODE_BRIGHTNESSUP

	ScanC = C.SDL_SCANCODE_C
	// ScanCalculator "Calculator" (the Calculator key)
	ScanCalculator = C.SDL_SCANCODE_CALCULATOR

	ScanCancel     = C.SDL_SCANCODE_CANCEL
	ScanCapsLock   = C.SDL_SCANCODE_CAPSLOCK
	ScanClear      = C.SDL_SCANCODE_CLEAR
	ScanClearAgain = C.SDL_SCANCODE_CLEARAGAIN
	ScanComma      = C.SDL_SCANCODE_COMMA
	// ScanComputer "Computer" (the My Computer key)
	ScanComputer = C.SDL_SCANCODE_COMPUTER

	ScanCopy  = C.SDL_SCANCODE_COPY
	ScanCrSel = C.SDL_SCANCODE_CRSEL
	// ScanCurrencySubUnit "CurrencySubUnit" (the Currency Subunit key)
	ScanCurrencySubUnit = C.SDL_SCANCODE_CURRENCYSUBUNIT
	// ScanCurrencyUnit "CurrencyUnit" (the Currency Unit key)
	ScanCurrencyUnit = C.SDL_SCANCODE_CURRENCYUNIT

	ScanCut = C.SDL_SCANCODE_CUT
	ScanD   = C.SDL_SCANCODE_D
	// ScanDecimalSeparator "DecimalSeparator" (the Decimal Separator key)
	ScanDecimalSeparator = C.SDL_SCANCODE_DECIMALSEPARATOR

	ScanDelete = C.SDL_SCANCODE_DELETE
	// ScanDisplaySwitch "DisplaySwitch" (display mirroring/dual display switch, video mode switch)
	ScanDisplaySwitch = C.SDL_SCANCODE_DISPLAYSWITCH
	// ScanDown "Down" (the Down arrow key (navigation keypad))
	ScanDown = C.SDL_SCANCODE_DOWN

	ScanE = C.SDL_SCANCODE_E
	// ScanEject "Eject" (the Eject key)
	ScanEject = C.SDL_SCANCODE_EJECT

	ScanEnd    = C.SDL_SCANCODE_END
	ScanEquals = C.SDL_SCANCODE_EQUALS
	// ScanEscape "Escape" (the Esc key)
	ScanEscape = C.SDL_SCANCODE_ESCAPE

	ScanExecute = C.SDL_SCANCODE_EXECUTE
	ScanExSel   = C.SDL_SCANCODE_EXSEL
	ScanF       = C.SDL_SCANCODE_F
	ScanF1      = C.SDL_SCANCODE_F1
	ScanF10     = C.SDL_SCANCODE_F10
	ScanF11     = C.SDL_SCANCODE_F11
	ScanF12     = C.SDL_SCANCODE_F12
	ScanF13     = C.SDL_SCANCODE_F13
	ScanF14     = C.SDL_SCANCODE_F14
	ScanF15     = C.SDL_SCANCODE_F15
	ScanF16     = C.SDL_SCANCODE_F16
	ScanF17     = C.SDL_SCANCODE_F17
	ScanF18     = C.SDL_SCANCODE_F18
	ScanF19     = C.SDL_SCANCODE_F19
	ScanF2      = C.SDL_SCANCODE_F2
	ScanF20     = C.SDL_SCANCODE_F20
	ScanF21     = C.SDL_SCANCODE_F21
	ScanF22     = C.SDL_SCANCODE_F22
	ScanF23     = C.SDL_SCANCODE_F23
	ScanF24     = C.SDL_SCANCODE_F24
	ScanF3      = C.SDL_SCANCODE_F3
	ScanF4      = C.SDL_SCANCODE_F4
	ScanF5      = C.SDL_SCANCODE_F5
	ScanF6      = C.SDL_SCANCODE_F6
	ScanF7      = C.SDL_SCANCODE_F7
	ScanF8      = C.SDL_SCANCODE_F8
	ScanF9      = C.SDL_SCANCODE_F9
	ScanFind    = C.SDL_SCANCODE_FIND
	ScanG       = C.SDL_SCANCODE_G
	// ScanBackQuote "`" (Located in the top left corner (on both ANSI and ISO keyboards). Produces GRAVE ACCENT and TILDE in a US Windows layout and in US and UK Mac layouts on ANSI keyboards, GRAVE ACCENT and NOT SIGN in a UK Windows layout, SECTION SIGN and PLUS-MINUS SIGN in US and UK Mac layouts on ISO keyboards, SECTION SIGN and DEGREE SIGN in a Swiss German layout (Mac: only on ISO keyboards), CIRCUMFLEX ACCENT and DEGREE SIGN in a German layout (Mac: only on ISO keyboards), SUPERSCRIPT TWO and TILDE in a French Windows layout, COMMERCIAL AT and NUMBER SIGN in a French Mac layout on ISO keyboards, and LESS-THAN SIGN and GREATER-THAN SIGN in a Swiss German, German, or French Mac layout on ANSI keyboards.)
	ScanBackQuote = C.SDL_SCANCODE_GRAVE

	ScanH    = C.SDL_SCANCODE_H
	ScanHelp = C.SDL_SCANCODE_HELP
	ScanHome = C.SDL_SCANCODE_HOME
	ScanI    = C.SDL_SCANCODE_I
	// ScanInsert "Insert" (insert on PC, help on some Mac keyboards (but does send code 73, not 117))
	ScanInsert = C.SDL_SCANCODE_INSERT

	ScanJ = C.SDL_SCANCODE_J
	ScanK = C.SDL_SCANCODE_K
	// ScanKBDIllumDown "KBDIllumDown" (the Keyboard Illumination Down key)
	ScanKBDIllumDown = C.SDL_SCANCODE_KBDILLUMDOWN
	// ScanKBDIllumToggle "KBDIllumToggle" (the Keyboard Illumination Toggle key)
	ScanKBDIllumToggle = C.SDL_SCANCODE_KBDILLUMTOGGLE
	// ScanKBDIllumUp "KBDIllumUp" (the Keyboard Illumination Up key)
	ScanKBDIllumUp = C.SDL_SCANCODE_KBDILLUMUP
	// ScanKeypad0 "Keypad 0" (the 0 key (numeric keypad))
	ScanKeypad0 = C.SDL_SCANCODE_KP_0
	// ScanKeypad00 "Keypad 00" (the 00 key (numeric keypad))
	ScanKeypad00 = C.SDL_SCANCODE_KP_00
	// ScanKeypad000 "Keypad 000" (the 000 key (numeric keypad))
	ScanKeypad000 = C.SDL_SCANCODE_KP_000
	// ScanKeypad1 "Keypad 1" (the 1 key (numeric keypad))
	ScanKeypad1 = C.SDL_SCANCODE_KP_1
	// ScanKeypad2 "Keypad 2" (the 2 key (numeric keypad))
	ScanKeypad2 = C.SDL_SCANCODE_KP_2
	// ScanKeypad3 "Keypad 3" (the 3 key (numeric keypad))
	ScanKeypad3 = C.SDL_SCANCODE_KP_3
	// ScanKeypad4 "Keypad 4" (the 4 key (numeric keypad))
	ScanKeypad4 = C.SDL_SCANCODE_KP_4
	// ScanKeypad5 "Keypad 5" (the 5 key (numeric keypad))
	ScanKeypad5 = C.SDL_SCANCODE_KP_5
	// ScanKeypad6 "Keypad 6" (the 6 key (numeric keypad))
	ScanKeypad6 = C.SDL_SCANCODE_KP_6
	// ScanKeypad7 "Keypad 7" (the 7 key (numeric keypad))
	ScanKeypad7 = C.SDL_SCANCODE_KP_7
	// ScanKeypad8 "Keypad 8" (the 8 key (numeric keypad))
	ScanKeypad8 = C.SDL_SCANCODE_KP_8
	// ScanKeypad9 "Keypad 9" (the 9 key (numeric keypad))
	ScanKeypad9 = C.SDL_SCANCODE_KP_9
	// ScanKeypadA "Keypad A" (the A key (numeric keypad))
	ScanKeypadA = C.SDL_SCANCODE_KP_A
	// ScanKeypadAmpersand "Keypad &" (the & key (numeric keypad))
	ScanKeypadAmpersand = C.SDL_SCANCODE_KP_AMPERSAND
	// ScanKeypadAt "Keypad @" (the @ key (numeric keypad))
	ScanKeypadAt = C.SDL_SCANCODE_KP_AT
	// ScanKeypadB "Keypad B" (the B key (numeric keypad))
	ScanKeypadB = C.SDL_SCANCODE_KP_B
	// ScanKeypadBackspace "Keypad Backspace" (the Backspace key (numeric keypad))
	ScanKeypadBackspace = C.SDL_SCANCODE_KP_BACKSPACE
	// ScanKeypadBinary "Keypad Binary" (the Binary key (numeric keypad))
	ScanKeypadBinary = C.SDL_SCANCODE_KP_BINARY
	// ScanKeypadC "Keypad C" (the C key (numeric keypad))
	ScanKeypadC = C.SDL_SCANCODE_KP_C
	// ScanKeypadClear "Keypad Clear" (the Clear key (numeric keypad))
	ScanKeypadClear = C.SDL_SCANCODE_KP_CLEAR
	// ScanKeypadClearEntry "Keypad ClearEntry" (the Clear Entry key (numeric keypad))
	ScanKeypadClearEntry = C.SDL_SCANCODE_KP_CLEARENTRY
	// ScanKeypadColon "Keypad :" (the : key (numeric keypad))
	ScanKeypadColon = C.SDL_SCANCODE_KP_COLON
	// ScanKeypadComma "Keypad ," (the Comma key (numeric keypad))
	ScanKeypadComma = C.SDL_SCANCODE_KP_COMMA
	// ScanKeypadD "Keypad D" (the D key (numeric keypad))
	ScanKeypadD = C.SDL_SCANCODE_KP_D
	// ScanKeypadDblAmpersand "Keypad &&" (the && key (numeric keypad))
	ScanKeypadDblAmpersand = C.SDL_SCANCODE_KP_DBLAMPERSAND
	// ScanKeypadDblVerticalBar "Keypad ||" (the || key (numeric keypad))
	ScanKeypadDblVerticalBar = C.SDL_SCANCODE_KP_DBLVERTICALBAR
	// ScanKeypadDecimal "Keypad Decimal" (the Decimal key (numeric keypad))
	ScanKeypadDecimal = C.SDL_SCANCODE_KP_DECIMAL
	// ScanKeypad "Keypad /" (the / key (numeric keypad))
	ScanKeypad = C.SDL_SCANCODE_KP_DIVIDE
	// ScanKeypadE "Keypad E" (the E key (numeric keypad))
	ScanKeypadE = C.SDL_SCANCODE_KP_E
	// ScanKeypadEnter "Keypad Enter" (the Enter key (numeric keypad))
	ScanKeypadEnter = C.SDL_SCANCODE_KP_ENTER
	// ScanKeypadEquals "Keypad =" (the = key (numeric keypad))
	ScanKeypadEquals = C.SDL_SCANCODE_KP_EQUALS
	// ScanKeypadEqualsAS400 "Keypad = (AS400)" (the Equals AS400 key (numeric keypad))
	ScanKeypadEqualsAS400 = C.SDL_SCANCODE_KP_EQUALSAS400
	// ScanKeypadExclam "Keypad !" (the ! key (numeric keypad))
	ScanKeypadExclam = C.SDL_SCANCODE_KP_EXCLAM
	// ScanKeypadF "Keypad F" (the F key (numeric keypad))
	ScanKeypadF = C.SDL_SCANCODE_KP_F
	// ScanKeypadGreater "Keypad >" (the Greater key (numeric keypad))
	ScanKeypadGreater = C.SDL_SCANCODE_KP_GREATER
	// ScanKeypadHash "Keypad #" (the # key (numeric keypad))
	ScanKeypadHash = C.SDL_SCANCODE_KP_HASH
	// ScanKeypadHexadecimal "Keypad Hexadecimal" (the Hexadecimal key (numeric keypad))
	ScanKeypadHexadecimal = C.SDL_SCANCODE_KP_HEXADECIMAL
	// ScanKeypadLeftBrace "Keypad {" (the Left Brace key (numeric keypad))
	ScanKeypadLeftBrace = C.SDL_SCANCODE_KP_LEFTBRACE
	// ScanKeypadLeftParen "Keypad (" (the Left Parenthesis key (numeric keypad))
	ScanKeypadLeftParen = C.SDL_SCANCODE_KP_LEFTPAREN
	// ScanKeypadLess "Keypad <" (the Less key (numeric keypad))
	ScanKeypadLess = C.SDL_SCANCODE_KP_LESS
	// ScanKeypadMemAdd "Keypad MemAdd" (the Mem Add key (numeric keypad))
	ScanKeypadMemAdd = C.SDL_SCANCODE_KP_MEMADD
	// ScanKeypadMemClear "Keypad MemClear" (the Mem Clear key (numeric keypad))
	ScanKeypadMemClear = C.SDL_SCANCODE_KP_MEMCLEAR
	// ScanKeypadMemDivide "Keypad MemDivide" (the Mem Divide key (numeric keypad))
	ScanKeypadMemDivide = C.SDL_SCANCODE_KP_MEMDIVIDE
	// ScanKeypadMemMultiply "Keypad MemMultiply" (the Mem Multiply key (numeric keypad))
	ScanKeypadMemMultiply = C.SDL_SCANCODE_KP_MEMMULTIPLY
	// ScanKeypadMemRecall "Keypad MemRecall" (the Mem Recall key (numeric keypad))
	ScanKeypadMemRecall = C.SDL_SCANCODE_KP_MEMRECALL
	// ScanKeypadMemStore "Keypad MemStore" (the Mem Store key (numeric keypad))
	ScanKeypadMemStore = C.SDL_SCANCODE_KP_MEMSTORE
	// ScanKeypadMemSubtract "Keypad MemSubtract" (the Mem Subtract key (numeric keypad))
	ScanKeypadMemSubtract = C.SDL_SCANCODE_KP_MEMSUBTRACT
	// ScanKeypadMinus "Keypad -" (the - key (numeric keypad))
	ScanKeypadMinus = C.SDL_SCANCODE_KP_MINUS
	// ScanKeypadMultiply "Keypad *" (the * key (numeric keypad))
	ScanKeypadMultiply = C.SDL_SCANCODE_KP_MULTIPLY
	// ScanKeypadOctal "Keypad Octal" (the Octal key (numeric keypad))
	ScanKeypadOctal = C.SDL_SCANCODE_KP_OCTAL
	// ScanKeypadPercent "Keypad %" (the Percent key (numeric keypad))
	ScanKeypadPercent = C.SDL_SCANCODE_KP_PERCENT
	// ScanKeypadPeriod "Keypad ." (the . key (numeric keypad))
	ScanKeypadPeriod = C.SDL_SCANCODE_KP_PERIOD
	// ScanKeypadPlus "Keypad +" (the + key (numeric keypad))
	ScanKeypadPlus = C.SDL_SCANCODE_KP_PLUS
	// ScanKeypadPlusMinus "Keypad +/-" (the +/- key (numeric keypad))
	ScanKeypadPlusMinus = C.SDL_SCANCODE_KP_PLUSMINUS
	// ScanKeypadPower "Keypad ^" (the Power key (numeric keypad))
	ScanKeypadPower = C.SDL_SCANCODE_KP_POWER
	// ScanKeypadRightBrace "Keypad }" (the Right Brace key (numeric keypad))
	ScanKeypadRightBrace = C.SDL_SCANCODE_KP_RIGHTBRACE
	// ScanKeypadRightParen "Keypad )" (the Right Parenthesis key (numeric keypad))
	ScanKeypadRightParen = C.SDL_SCANCODE_KP_RIGHTPAREN
	// ScanKeypadSpace "Keypad Space" (the Space key (numeric keypad))
	ScanKeypadSpace = C.SDL_SCANCODE_KP_SPACE
	// ScanKeypadTab "Keypad Tab" (the Tab key (numeric keypad))
	ScanKeypadTab = C.SDL_SCANCODE_KP_TAB
	// ScanKeypadVerticalBar "Keypad |" (the | key (numeric keypad))
	ScanKeypadVerticalBar = C.SDL_SCANCODE_KP_VERTICALBAR
	// ScanKeypadXOR "Keypad XOR" (the XOR key (numeric keypad))
	ScanKeypadXOR = C.SDL_SCANCODE_KP_XOR

	ScanL = C.SDL_SCANCODE_L
	// ScanLeftAlt "Left Alt" (alt, option)
	ScanLeftAlt = C.SDL_SCANCODE_LALT

	ScanLeftCtrl = C.SDL_SCANCODE_LCTRL
	// ScanLeft "Left" (the Left arrow key (navigation keypad))
	ScanLeft = C.SDL_SCANCODE_LEFT

	ScanLeftBracket = C.SDL_SCANCODE_LEFTBRACKET
	// ScanLeftGUI "Left GUI" (windows, command (apple), meta)
	ScanLeftGUI = C.SDL_SCANCODE_LGUI

	ScanLeftShift = C.SDL_SCANCODE_LSHIFT
	ScanM         = C.SDL_SCANCODE_M
	// ScanMail "Mail" (the Mail/eMail key)
	ScanMail = C.SDL_SCANCODE_MAIL
	// ScanMediaSelect "MediaSelect" (the Media Select key)
	ScanMediaSelect = C.SDL_SCANCODE_MEDIASELECT

	ScanMenu  = C.SDL_SCANCODE_MENU
	ScanMinus = C.SDL_SCANCODE_MINUS
	// ScanModeSwitch "ModeSwitch" (I'm not sure if this is really not covered by any of the above, but since there's a special KMOD_MODE for it I'm adding it here)
	ScanModeSwitch = C.SDL_SCANCODE_MODE

	ScanMute = C.SDL_SCANCODE_MUTE
	ScanN    = C.SDL_SCANCODE_N
	// ScanNumlock "Numlock" (the Num Lock key (PC) / the Clear key (Mac))
	ScanNumlock = C.SDL_SCANCODE_NUMLOCKCLEAR

	ScanO        = C.SDL_SCANCODE_O
	ScanOper     = C.SDL_SCANCODE_OPER
	ScanOut      = C.SDL_SCANCODE_OUT
	ScanP        = C.SDL_SCANCODE_P
	ScanPageDown = C.SDL_SCANCODE_PAGEDOWN
	ScanPageUp   = C.SDL_SCANCODE_PAGEUP
	ScanPaste    = C.SDL_SCANCODE_PASTE
	// ScanPause "Pause" (the Pause / Break key)
	ScanPause = C.SDL_SCANCODE_PAUSE

	ScanPeriod = C.SDL_SCANCODE_PERIOD
	// ScanPower "Power" (The USB document says this is a status flag, not a physical key - but some Mac keyboards do have a power key.)
	ScanPower = C.SDL_SCANCODE_POWER

	ScanPrintScreen = C.SDL_SCANCODE_PRINTSCREEN
	ScanPrior       = C.SDL_SCANCODE_PRIOR
	ScanQ           = C.SDL_SCANCODE_Q
	ScanR           = C.SDL_SCANCODE_R
	// ScanRightAlt "Right Alt" (alt gr, option)
	ScanRightAlt = C.SDL_SCANCODE_RALT

	ScanRightCtrl = C.SDL_SCANCODE_RCTRL
	// ScanReturn "Return" (the Enter key (main keyboard))
	ScanReturn = C.SDL_SCANCODE_RETURN

	ScanReturn = C.SDL_SCANCODE_RETURN2
	// ScanRightGUI "Right GUI" (windows, command (apple), meta)
	ScanRightGUI = C.SDL_SCANCODE_RGUI
	// ScanRight "Right" (the Right arrow key (navigation keypad))
	ScanRight = C.SDL_SCANCODE_RIGHT

	ScanRightBracket = C.SDL_SCANCODE_RIGHTBRACKET
	ScanRightShift   = C.SDL_SCANCODE_RSHIFT
	ScanS            = C.SDL_SCANCODE_S
	ScanScrollLock   = C.SDL_SCANCODE_SCROLLLOCK
	ScanSelect       = C.SDL_SCANCODE_SELECT
	ScanSemicolon    = C.SDL_SCANCODE_SEMICOLON
	ScanSeparator    = C.SDL_SCANCODE_SEPARATOR
	Scan             = C.SDL_SCANCODE_SLASH
	// ScanSleep "Sleep" (the Sleep key)
	ScanSleep = C.SDL_SCANCODE_SLEEP
	// ScanSpace "Space" (the Space Bar key(s))
	ScanSpace = C.SDL_SCANCODE_SPACE

	ScanStop = C.SDL_SCANCODE_STOP
	// ScanSysReq "SysReq" (the SysReq key)
	ScanSysReq = C.SDL_SCANCODE_SYSREQ

	ScanT = C.SDL_SCANCODE_T
	// ScanTab "Tab" (the Tab key)
	ScanTab = C.SDL_SCANCODE_TAB
	// ScanThousandsSeparator "ThousandsSeparator" (the Thousands Separator key)
	ScanThousandsSeparator = C.SDL_SCANCODE_THOUSANDSSEPARATOR

	ScanU    = C.SDL_SCANCODE_U
	ScanUndo = C.SDL_SCANCODE_UNDO
	// ScanUnknown "" (no name, empty string)
	ScanUnknown = C.SDL_SCANCODE_UNKNOWN
	// ScanUp "Up" (the Up arrow key (navigation keypad))
	ScanUp = C.SDL_SCANCODE_UP

	ScanV          = C.SDL_SCANCODE_V
	ScanVolumeDown = C.SDL_SCANCODE_VOLUMEDOWN
	ScanVolumeUp   = C.SDL_SCANCODE_VOLUMEUP
	ScanW          = C.SDL_SCANCODE_W
	// ScanWWW "WWW" (the WWW/World Wide Web key)
	ScanWWW = C.SDL_SCANCODE_WWW

	ScanX = C.SDL_SCANCODE_X
	ScanY = C.SDL_SCANCODE_Y
	ScanZ = C.SDL_SCANCODE_Z
	// ScanInternational1 "" (no name, empty string; used on Asian keyboards, see footnotes in USB doc)
	ScanInternational1 = C.SDL_SCANCODE_INTERNATIONAL1
	// ScanInternational2 "" (no name, empty string)
	ScanInternational2 = C.SDL_SCANCODE_INTERNATIONAL2
	// ScanInternational3 "" (no name, empty string; Yen)
	ScanInternational3 = C.SDL_SCANCODE_INTERNATIONAL3
	// ScanInternational4 "" (no name, empty string)
	ScanInternational4 = C.SDL_SCANCODE_INTERNATIONAL4
	// ScanInternational5 "" (no name, empty string)
	ScanInternational5 = C.SDL_SCANCODE_INTERNATIONAL5
	// ScanInternational6 "" (no name, empty string)
	ScanInternational6 = C.SDL_SCANCODE_INTERNATIONAL6
	// ScanInternational7 "" (no name, empty string)
	ScanInternational7 = C.SDL_SCANCODE_INTERNATIONAL7
	// ScanInternational8 "" (no name, empty string)
	ScanInternational8 = C.SDL_SCANCODE_INTERNATIONAL8
	// ScanInternational9 "" (no name, empty string)
	ScanInternational9 = C.SDL_SCANCODE_INTERNATIONAL9
	// ScanLang1 "" (no name, empty string; Hangul/English toggle)
	ScanLang1 = C.SDL_SCANCODE_LANG1
	// ScanLang2 "" (no name, empty string; Hanja conversion)
	ScanLang2 = C.SDL_SCANCODE_LANG2
	// ScanLang3 "" (no name, empty string; Katakana)
	ScanLang3 = C.SDL_SCANCODE_LANG3
	// ScanLang4 "" (no name, empty string; Hiragana)
	ScanLang4 = C.SDL_SCANCODE_LANG4
	// ScanLang5 "" (no name, empty string; Zenkaku/Hankaku)
	ScanLang5 = C.SDL_SCANCODE_LANG5
	// ScanLang6 "" (no name, empty string; reserved)
	ScanLang6 = C.SDL_SCANCODE_LANG6
	// ScanLang7 "" (no name, empty string; reserved)
	ScanLang7 = C.SDL_SCANCODE_LANG7
	// ScanLang8 "" (no name, empty string; reserved)
	ScanLang8 = C.SDL_SCANCODE_LANG8
	// ScanLang9 "" (no name, empty string; reserved)
	ScanLang9 = C.SDL_SCANCODE_LANG9
	// ScanLockingCapsLock "" (no name, empty string)
	ScanLockingCapsLock = C.SDL_SCANCODE_LOCKINGCAPSLOCK
	// ScanLockingNumLock "" (no name, empty string)
	ScanLockingNumLock = C.SDL_SCANCODE_LOCKINGNUMLOCK
	// ScanLockingScrollLock "" (no name, empty string)
	ScanLockingScrollLock = C.SDL_SCANCODE_LOCKINGSCROLLLOCK
	// ScanNonUSBackslash "" (no name, empty string; This is the additional key that ISO keyboards have over ANSI ones, located between left shift and Y. Produces GRAVE ACCENT and TILDE in a US or UK Mac layout, REVERSE SOLIDUS (backslash) and VERTICAL LINE in a US or UK Windows layout, and LESS-THAN SIGN and GREATER-THAN SIGN in a Swiss German, German, or French layout.)
	ScanNonUSBackslash = C.SDL_SCANCODE_NONUSBACKSLASH
	// ScanNonUSHash "#" (ISO USB keyboards actually use this code instead of 49 for the same key, but all OSes I've seen treat the two codes identically. So, as an implementor, unless your keyboard generates both of those codes and your OS treats them differently, you should generate SDL_SCANCODE_BACKSLASH instead of this code. As a user, you should not rely on this code because SDL will never generate it with most (all?) keyboards.)
	ScanNonUSHash = C.SDL_SCANCODE_NONUSHASH
)
