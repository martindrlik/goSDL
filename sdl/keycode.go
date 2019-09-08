package sdl

// #include <SDL2/SDL_keycode.h>
import "C"

// Key contains key information used in key events.
type Key struct {
	Physical C.SDL_Scancode
	Virtual  C.SDL_Keycode
	Mod      uint16
	_        uint32
}

const (
	Key0                  = C.SDLK_0                  // "0" SDL_SCANCODE_0
	Key1                  = C.SDLK_1                  // "1" SDL_SCANCODE_1
	Key2                  = C.SDLK_2                  // "2" SDL_SCANCODE_2
	Key3                  = C.SDLK_3                  // "3" SDL_SCANCODE_3
	Key4                  = C.SDLK_4                  // "4" SDL_SCANCODE_4
	Key5                  = C.SDLK_5                  // "5" SDL_SCANCODE_5
	Key6                  = C.SDLK_6                  // "6" SDL_SCANCODE_6
	Key7                  = C.SDLK_7                  // "7" SDL_SCANCODE_7
	Key8                  = C.SDLK_8                  // "8" SDL_SCANCODE_8
	Key9                  = C.SDLK_9                  // "9" SDL_SCANCODE_9
	Keya                  = C.SDLK_a                  // "A" SDL_SCANCODE_A
	KeyACBack             = C.SDLK_AC_BACK            // "AC Back" (the Back key (application control keypad)) SDL_SCANCODE_AC_BACK
	KeyACBookmarks        = C.SDLK_AC_BOOKMARKS       // "AC Bookmarks" (the Bookmarks key (application control keypad)) SDL_SCANCODE_AC_BOOKMARKS
	KeyACForward          = C.SDLK_AC_FORWARD         // "AC Forward" (the Forward key (application control keypad)) SDL_SCANCODE_AC_FORWARD
	KeyACHome             = C.SDLK_AC_HOME            // "AC Home" (the Home key (application control keypad)) SDL_SCANCODE_AC_HOME
	KeyACRefresh          = C.SDLK_AC_REFRESH         // "AC Refresh" (the Refresh key (application control keypad)) SDL_SCANCODE_AC_REFRESH
	KeyACSearch           = C.SDLK_AC_SEARCH          // "AC Search" (the Search key (application control keypad)) SDL_SCANCODE_AC_SEARCH
	KeyACStop             = C.SDLK_AC_STOP            // "AC Stop" (the Stop key (application control keypad)) SDL_SCANCODE_AC_STOP
	KeyAgain              = C.SDLK_AGAIN              // "Again" (the Again key (Redo)) SDL_SCANCODE_AGAIN
	KeyAltErase           = C.SDLK_ALTERASE           // "AltErase" (Erase-Eaze) SDL_SCANCODE_ALTERASE
	KeyQuote              = C.SDLK_QUOTE              // "'" SDL_SCANCODE_APOSTROPHE
	KeyApplication        = C.SDLK_APPLICATION        // "Application" (the Application / Compose / Context Menu (Windows) key) SDL_SCANCODE_APPLICATION
	KeyAudioMute          = C.SDLK_AUDIOMUTE          // "AudioMute" (the Mute volume key) SDL_SCANCODE_AUDIOMUTE
	KeyAudioNext          = C.SDLK_AUDIONEXT          // "AudioNext" (the Next Track media key) SDL_SCANCODE_AUDIONEXT
	KeyAudioPlay          = C.SDLK_AUDIOPLAY          // "AudioPlay" (the Play media key) SDL_SCANCODE_AUDIOPLAY
	KeyAudioPrev          = C.SDLK_AUDIOPREV          // "AudioPrev" (the Previous Track media key) SDL_SCANCODE_AUDIOPREV
	KeyAudioStop          = C.SDLK_AUDIOSTOP          // "AudioStop" (the Stop media key) SDL_SCANCODE_AUDIOSTOP
	Keyb                  = C.SDLK_b                  // "B" SDL_SCANCODE_B
	KeyBackslash          = C.SDLK_BACKSLASH          // "\" (Located at the lower left of the return key on ISO keyboards and at the right end of the QWERTY row on ANSI keyboards. Produces REVERSE SOLIDUS (backslash) and VERTICAL LINE in a US layout, REVERSE SOLIDUS and VERTICAL LINE in a UK Mac layout, NUMBER SIGN and TILDE in a UK Windows layout, DOLLAR SIGN and POUND SIGN in a Swiss German layout, NUMBER SIGN and APOSTROPHE in a German layout, GRAVE ACCENT and POUND SIGN in a French Mac layout, and ASTERISK and MICRO SIGN in a French Windows layout.) SDL_SCANCODE_BACKSLASH
	KeyBackspace          = C.SDLK_BACKSPACE          // "Backspace" SDL_SCANCODE_BACKSPACE
	KeyBrightnessDown     = C.SDLK_BRIGHTNESSDOWN     // "BrightnessDown" (the Brightness Down key) SDL_SCANCODE_BRIGHTNESSDOWN
	KeyBrightnessUp       = C.SDLK_BRIGHTNESSUP       // "BrightnessUp" (the Brightness Up key) SDL_SCANCODE_BRIGHTNESSUP
	Keyc                  = C.SDLK_c                  // "C" SDL_SCANCODE_C
	KeyCalculator         = C.SDLK_CALCULATOR         // "Calculator" (the Calculator key) SDL_SCANCODE_CALCULATOR
	KeyCancel             = C.SDLK_CANCEL             // "Cancel" SDL_SCANCODE_CANCEL
	KeyCapsLock           = C.SDLK_CAPSLOCK           // "CapsLock" SDL_SCANCODE_CAPSLOCK
	KeyClear              = C.SDLK_CLEAR              // "Clear" SDL_SCANCODE_CLEAR
	KeyClearAgain         = C.SDLK_CLEARAGAIN         // "Clear / Again" SDL_SCANCODE_CLEARAGAIN
	KeyComma              = C.SDLK_COMMA              // "," SDL_SCANCODE_COMMA
	KeyComputer           = C.SDLK_COMPUTER           // "Computer" (the My Computer key) SDL_SCANCODE_COMPUTER
	KeyCopy               = C.SDLK_COPY               // "Copy" SDL_SCANCODE_COPY
	KeyCrSel              = C.SDLK_CRSEL              // "CrSel" SDL_SCANCODE_CRSEL
	KeyCurrencySubUnit    = C.SDLK_CURRENCYSUBUNIT    // "CurrencySubUnit" (the Currency Subunit key) SDL_SCANCODE_CURRENCYSUBUNIT
	KeyCurrencyUnit       = C.SDLK_CURRENCYUNIT       // "CurrencyUnit" (the Currency Unit key) SDL_SCANCODE_CURRENCYUNIT
	KeyCut                = C.SDLK_CUT                // "Cut" SDL_SCANCODE_CUT
	Keyd                  = C.SDLK_d                  // "D" SDL_SCANCODE_D
	KeyDecimalSeparator   = C.SDLK_DECIMALSEPARATOR   // "DecimalSeparator" (the Decimal Separator key) SDL_SCANCODE_DECIMALSEPARATOR
	KeyDelete             = C.SDLK_DELETE             // "Delete" SDL_SCANCODE_DELETE
	KeyDisplaySwitch      = C.SDLK_DISPLAYSWITCH      // "DisplaySwitch" (display mirroring/dual display switch, video mode switch) SDL_SCANCODE_DISPLAYSWITCH
	KeyDown               = C.SDLK_DOWN               // "Down" (the Down arrow key (navigation keypad)) SDL_SCANCODE_DOWN
	Keye                  = C.SDLK_e                  // "E" SDL_SCANCODE_E
	KeyEject              = C.SDLK_EJECT              // "Eject" (the Eject key) SDL_SCANCODE_EJECT
	KeyEnd                = C.SDLK_END                // "End" SDL_SCANCODE_END
	KeyEquals             = C.SDLK_EQUALS             // "=" SDL_SCANCODE_EQUALS
	KeyEscape             = C.SDLK_ESCAPE             // "Escape" (the Esc key) SDL_SCANCODE_ESCAPE
	KeyExecute            = C.SDLK_EXECUTE            // "Execute" SDL_SCANCODE_EXECUTE
	KeyExSel              = C.SDLK_EXSEL              // "ExSel" SDL_SCANCODE_EXSEL
	Keyf                  = C.SDLK_f                  // "F" SDL_SCANCODE_F
	KeyF1                 = C.SDLK_F1                 // "F1" SDL_SCANCODE_F1
	KeyF10                = C.SDLK_F10                // "F10" SDL_SCANCODE_F10
	KeyF11                = C.SDLK_F11                // "F11" SDL_SCANCODE_F11
	KeyF12                = C.SDLK_F12                // "F12" SDL_SCANCODE_F12
	KeyF13                = C.SDLK_F13                // "F13" SDL_SCANCODE_F13
	KeyF14                = C.SDLK_F14                // "F14" SDL_SCANCODE_F14
	KeyF15                = C.SDLK_F15                // "F15" SDL_SCANCODE_F15
	KeyF16                = C.SDLK_F16                // "F16" SDL_SCANCODE_F16
	KeyF17                = C.SDLK_F17                // "F17" SDL_SCANCODE_F17
	KeyF18                = C.SDLK_F18                // "F18" SDL_SCANCODE_F18
	KeyF19                = C.SDLK_F19                // "F19" SDL_SCANCODE_F19
	KeyF2                 = C.SDLK_F2                 // "F2" SDL_SCANCODE_F2
	KeyF20                = C.SDLK_F20                // "F20" SDL_SCANCODE_F20
	KeyF21                = C.SDLK_F21                // "F21" SDL_SCANCODE_F21
	KeyF22                = C.SDLK_F22                // "F22" SDL_SCANCODE_F22
	KeyF23                = C.SDLK_F23                // "F23" SDL_SCANCODE_F23
	KeyF24                = C.SDLK_F24                // "F24" SDL_SCANCODE_F24
	KeyF3                 = C.SDLK_F3                 // "F3" SDL_SCANCODE_F3
	KeyF4                 = C.SDLK_F4                 // "F4" SDL_SCANCODE_F4
	KeyF5                 = C.SDLK_F5                 // "F5" SDL_SCANCODE_F5
	KeyF6                 = C.SDLK_F6                 // "F6" SDL_SCANCODE_F6
	KeyF7                 = C.SDLK_F7                 // "F7" SDL_SCANCODE_F7
	KeyF8                 = C.SDLK_F8                 // "F8" SDL_SCANCODE_F8
	KeyF9                 = C.SDLK_F9                 // "F9" SDL_SCANCODE_F9
	KeyFind               = C.SDLK_FIND               // "Find" SDL_SCANCODE_FIND
	Keyg                  = C.SDLK_g                  // "G" SDL_SCANCODE_G
	KeyBackQuote          = C.SDLK_BACKQUOTE          // "`" (Located in the top left corner (on both ANSI and ISO keyboards). Produces GRAVE ACCENT and TILDE in a US Windows layout and in US and UK Mac layouts on ANSI keyboards, GRAVE ACCENT and NOT SIGN in a UK Windows layout, SECTION SIGN and PLUS-MINUS SIGN in US and UK Mac layouts on ISO keyboards, SECTION SIGN and DEGREE SIGN in a Swiss German layout (Mac: only on ISO keyboards), CIRCUMFLEX ACCENT and DEGREE SIGN in a German layout (Mac: only on ISO keyboards), SUPERSCRIPT TWO and TILDE in a French Windows layout, COMMERCIAL AT and NUMBER SIGN in a French Mac layout on ISO keyboards, and LESS-THAN SIGN and GREATER-THAN SIGN in a Swiss German, German, or French Mac layout on ANSI keyboards.) SDL_SCANCODE_GRAVE
	Keyh                  = C.SDLK_h                  // "H" SDL_SCANCODE_H
	KeyHelp               = C.SDLK_HELP               // "Help" SDL_SCANCODE_HELP
	KeyHome               = C.SDLK_HOME               // "Home" SDL_SCANCODE_HOME
	Keyi                  = C.SDLK_i                  // "I" SDL_SCANCODE_I
	KeyInsert             = C.SDLK_INSERT             // "Insert" (insert on PC, help on some Mac keyboards (but does send code 73, not 117)) SDL_SCANCODE_INSERT
	Keyj                  = C.SDLK_j                  // "J" SDL_SCANCODE_J
	Keyk                  = C.SDLK_k                  // "K" SDL_SCANCODE_K
	KeyKBDIllumDown       = C.SDLK_KBDILLUMDOWN       // "KBDIllumDown" (the Keyboard Illumination Down key) SDL_SCANCODE_KBDILLUMDOWN
	KeyKBDIllumToggle     = C.SDLK_KBDILLUMTOGGLE     // "KBDIllumToggle" (the Keyboard Illumination Toggle key) SDL_SCANCODE_KBDILLUMTOGGLE
	KeyKBDIllumUp         = C.SDLK_KBDILLUMUP         // "KBDIllumUp" (the Keyboard Illumination Up key) SDL_SCANCODE_KBDILLUMUP
	KeyKp0                = C.SDLK_KP_0               // "Keypad 0" (the 0 key (numeric keypad)) SDL_SCANCODE_KP_0
	KeyKp00               = C.SDLK_KP_00              // "Keypad 00" (the 00 key (numeric keypad)) SDL_SCANCODE_KP_00
	KeyKp000              = C.SDLK_KP_000             // "Keypad 000" (the 000 key (numeric keypad)) SDL_SCANCODE_KP_000
	KeyKp1                = C.SDLK_KP_1               // "Keypad 1" (the 1 key (numeric keypad)) SDL_SCANCODE_KP_1
	KeyKp2                = C.SDLK_KP_2               // "Keypad 2" (the 2 key (numeric keypad)) SDL_SCANCODE_KP_2
	KeyKp3                = C.SDLK_KP_3               // "Keypad 3" (the 3 key (numeric keypad)) SDL_SCANCODE_KP_3
	KeyKp4                = C.SDLK_KP_4               // "Keypad 4" (the 4 key (numeric keypad)) SDL_SCANCODE_KP_4
	KeyKp5                = C.SDLK_KP_5               // "Keypad 5" (the 5 key (numeric keypad)) SDL_SCANCODE_KP_5
	KeyKp6                = C.SDLK_KP_6               // "Keypad 6" (the 6 key (numeric keypad)) SDL_SCANCODE_KP_6
	KeyKp7                = C.SDLK_KP_7               // "Keypad 7" (the 7 key (numeric keypad)) SDL_SCANCODE_KP_7
	KeyKp8                = C.SDLK_KP_8               // "Keypad 8" (the 8 key (numeric keypad)) SDL_SCANCODE_KP_8
	KeyKp9                = C.SDLK_KP_9               // "Keypad 9" (the 9 key (numeric keypad)) SDL_SCANCODE_KP_9
	KeyKpA                = C.SDLK_KP_A               // "Keypad A" (the A key (numeric keypad)) SDL_SCANCODE_KP_A
	KeyKpAmpersand        = C.SDLK_KP_AMPERSAND       // "Keypad &" (the & key (numeric keypad)) SDL_SCANCODE_KP_AMPERSAND
	KeyKpAt               = C.SDLK_KP_AT              // "Keypad @" (the @ key (numeric keypad)) SDL_SCANCODE_KP_AT
	KeyKpB                = C.SDLK_KP_B               // "Keypad B" (the B key (numeric keypad)) SDL_SCANCODE_KP_B
	KeyKpBackspace        = C.SDLK_KP_BACKSPACE       // "Keypad Backspace" (the Backspace key (numeric keypad)) SDL_SCANCODE_KP_BACKSPACE
	KeyKpBinary           = C.SDLK_KP_BINARY          // "Keypad Binary" (the Binary key (numeric keypad)) SDL_SCANCODE_KP_BINARY
	KeyKpC                = C.SDLK_KP_C               // "Keypad C" (the C key (numeric keypad)) SDL_SCANCODE_KP_C
	KeyKpClear            = C.SDLK_KP_CLEAR           // "Keypad Clear" (the Clear key (numeric keypad)) SDL_SCANCODE_KP_CLEAR
	KeyKpClearEntry       = C.SDLK_KP_CLEARENTRY      // "Keypad ClearEntry" (the Clear Entry key (numeric keypad)) SDL_SCANCODE_KP_CLEARENTRY
	KeyKpColon            = C.SDLK_KP_COLON           // "Keypad :" (the : key (numeric keypad)) SDL_SCANCODE_KP_COLON
	KeyKpComma            = C.SDLK_KP_COMMA           // "Keypad ," (the Comma key (numeric keypad)) SDL_SCANCODE_KP_COMMA
	KeyKpD                = C.SDLK_KP_D               // "Keypad D" (the D key (numeric keypad)) SDL_SCANCODE_KP_D
	KeyKpDblAmpersand     = C.SDLK_KP_DBLAMPERSAND    // "Keypad &&" (the && key (numeric keypad)) SDL_SCANCODE_KP_DBLAMPERSAND
	KeyKpDblVerticalBar   = C.SDLK_KP_DBLVERTICALBAR  // "Keypad ||" (the || key (numeric keypad)) SDL_SCANCODE_KP_DBLVERTICALBAR
	KeyKpDecimal          = C.SDLK_KP_DECIMAL         // "Keypad Decimal" (the Decimal key (numeric keypad)) SDL_SCANCODE_KP_DECIMAL
	KeyKpDivide           = C.SDLK_KP_DIVIDE          // "Keypad /" (the / key (numeric keypad)) SDL_SCANCODE_KP_DIVIDE
	KeyKpE                = C.SDLK_KP_E               // "Keypad E" (the E key (numeric keypad)) SDL_SCANCODE_KP_E
	KeyKpEnter            = C.SDLK_KP_ENTER           // "Keypad Enter" (the Enter key (numeric keypad)) SDL_SCANCODE_KP_ENTER
	KeyKpEquals           = C.SDLK_KP_EQUALS          // "Keypad =" (the = key (numeric keypad)) SDL_SCANCODE_KP_EQUALS
	KeyKpEqualsAS400      = C.SDLK_KP_EQUALSAS400     // "Keypad = (AS400)" (the Equals AS400 key (numeric keypad)) SDL_SCANCODE_KP_EQUALSAS400
	KeyKpExclam           = C.SDLK_KP_EXCLAM          // "Keypad !" (the ! key (numeric keypad)) SDL_SCANCODE_KP_EXCLAM
	KeyKpF                = C.SDLK_KP_F               // "Keypad F" (the F key (numeric keypad)) SDL_SCANCODE_KP_F
	KeyKpGreater          = C.SDLK_KP_GREATER         // "Keypad >" (the Greater key (numeric keypad)) SDL_SCANCODE_KP_GREATER
	KeyKpHash             = C.SDLK_KP_HASH            // "Keypad #" (the # key (numeric keypad)) SDL_SCANCODE_KP_HASH
	KeyKpHexadecimal      = C.SDLK_KP_HEXADECIMAL     // "Keypad Hexadecimal" (the Hexadecimal key (numeric keypad)) SDL_SCANCODE_KP_HEXADECIMAL
	KeyKpLeftBrace        = C.SDLK_KP_LEFTBRACE       // "Keypad {" (the Left Brace key (numeric keypad)) SDL_SCANCODE_KP_LEFTBRACE
	KeyKpLeftParen        = C.SDLK_KP_LEFTPAREN       // "Keypad (" (the Left Parenthesis key (numeric keypad)) SDL_SCANCODE_KP_LEFTPAREN
	KeyKpLess             = C.SDLK_KP_LESS            // "Keypad <" (the Less key (numeric keypad)) SDL_SCANCODE_KP_LESS
	KeyKpMemAdd           = C.SDLK_KP_MEMADD          // "Keypad MemAdd" (the Mem Add key (numeric keypad)) SDL_SCANCODE_KP_MEMADD
	KeyKpMemClear         = C.SDLK_KP_MEMCLEAR        // "Keypad MemClear" (the Mem Clear key (numeric keypad)) SDL_SCANCODE_KP_MEMCLEAR
	KeyKpMemDivide        = C.SDLK_KP_MEMDIVIDE       // "Keypad MemDivide" (the Mem Divide key (numeric keypad)) SDL_SCANCODE_KP_MEMDIVIDE
	KeyKpMemMultiply      = C.SDLK_KP_MEMMULTIPLY     // "Keypad MemMultiply" (the Mem Multiply key (numeric keypad)) SDL_SCANCODE_KP_MEMMULTIPLY
	KeyKpMemRecall        = C.SDLK_KP_MEMRECALL       // "Keypad MemRecall" (the Mem Recall key (numeric keypad)) SDL_SCANCODE_KP_MEMRECALL
	KeyKpMemStore         = C.SDLK_KP_MEMSTORE        // "Keypad MemStore" (the Mem Store key (numeric keypad)) SDL_SCANCODE_KP_MEMSTORE
	KeyKpMemSubtract      = C.SDLK_KP_MEMSUBTRACT     // "Keypad MemSubtract" (the Mem Subtract key (numeric keypad)) SDL_SCANCODE_KP_MEMSUBTRACT
	KeyKpMinus            = C.SDLK_KP_MINUS           // "Keypad -" (the - key (numeric keypad)) SDL_SCANCODE_KP_MINUS
	KeyKpMultiply         = C.SDLK_KP_MULTIPLY        // "Keypad *" (the * key (numeric keypad)) SDL_SCANCODE_KP_MULTIPLY
	KeyKpOctal            = C.SDLK_KP_OCTAL           // "Keypad Octal" (the Octal key (numeric keypad)) SDL_SCANCODE_KP_OCTAL
	KeyKpPercent          = C.SDLK_KP_PERCENT         // "Keypad %" (the Percent key (numeric keypad)) SDL_SCANCODE_KP_PERCENT
	KeyKpPeriod           = C.SDLK_KP_PERIOD          // "Keypad ." (the . key (numeric keypad)) SDL_SCANCODE_KP_PERIOD
	KeyKpPlus             = C.SDLK_KP_PLUS            // "Keypad +" (the + key (numeric keypad)) SDL_SCANCODE_KP_PLUS
	KeyKpPlusMinus        = C.SDLK_KP_PLUSMINUS       // "Keypad +/-" (the +/- key (numeric keypad)) SDL_SCANCODE_KP_PLUSMINUS
	KeyKpPower            = C.SDLK_KP_POWER           // "Keypad ^" (the Power key (numeric keypad)) SDL_SCANCODE_KP_POWER
	KeyKpRightBrace       = C.SDLK_KP_RIGHTBRACE      // "Keypad }" (the Right Brace key (numeric keypad)) SDL_SCANCODE_KP_RIGHTBRACE
	KeyKpRightParen       = C.SDLK_KP_RIGHTPAREN      // "Keypad )" (the Right Parenthesis key (numeric keypad)) SDL_SCANCODE_KP_RIGHTPAREN
	KeyKpSpace            = C.SDLK_KP_SPACE           // "Keypad Space" (the Space key (numeric keypad)) SDL_SCANCODE_KP_SPACE
	KeyKpTab              = C.SDLK_KP_TAB             // "Keypad Tab" (the Tab key (numeric keypad)) SDL_SCANCODE_KP_TAB
	KeyKpVerticalBar      = C.SDLK_KP_VERTICALBAR     // "Keypad |" (the | key (numeric keypad)) SDL_SCANCODE_KP_VERTICALBAR
	KeyKpXOR              = C.SDLK_KP_XOR             // "Keypad XOR" (the XOR key (numeric keypad)) SDL_SCANCODE_KP_XOR
	Keyl                  = C.SDLK_l                  // "L" SDL_SCANCODE_L
	KeyLAlt               = C.SDLK_LALT               // "Left Alt" (alt, option) SDL_SCANCODE_LALT
	KeyLCtrl              = C.SDLK_LCTRL              // "Left Ctrl" SDL_SCANCODE_LCTRL
	KeyLeft               = C.SDLK_LEFT               // "Left" (the Left arrow key (navigation keypad)) SDL_SCANCODE_LEFT
	KeyLeftBracket        = C.SDLK_LEFTBRACKET        // "[" SDL_SCANCODE_LEFTBRACKET
	KeyLGUI               = C.SDLK_LGUI               // "Left GUI" (windows, command (apple), meta) SDL_SCANCODE_LGUI
	KeyLShift             = C.SDLK_LSHIFT             // "Left Shift" SDL_SCANCODE_LSHIFT
	Keym                  = C.SDLK_m                  // "M" SDL_SCANCODE_M
	KeyMail               = C.SDLK_MAIL               // "Mail" (the Mail/eMail key) SDL_SCANCODE_MAIL
	KeyMediaSelect        = C.SDLK_MEDIASELECT        // "MediaSelect" (the Media Select key) SDL_SCANCODE_MEDIASELECT
	KeyMenu               = C.SDLK_MENU               // "Menu" SDL_SCANCODE_MENU
	KeyMinus              = C.SDLK_MINUS              // "-" SDL_SCANCODE_MINUS
	KeyMode               = C.SDLK_MODE               // "ModeSwitch" (I'm not sure if this is really not covered by any of the above, but since there's a special KMOD_MODE for it I'm adding it here) SDL_SCANCODE_MODE
	KeyMute               = C.SDLK_MUTE               // "Mute" SDL_SCANCODE_MUTE
	Keyn                  = C.SDLK_n                  // "N" SDL_SCANCODE_N
	KeyNumLockClear       = C.SDLK_NUMLOCKCLEAR       // "Numlock" (the Num Lock key (PC) / the Clear key (Mac)) SDL_SCANCODE_NUMLOCKCLEAR
	Keyo                  = C.SDLK_o                  // "O" SDL_SCANCODE_O
	KeyOper               = C.SDLK_OPER               // "Oper" SDL_SCANCODE_OPER
	KeyOut                = C.SDLK_OUT                // "Out" SDL_SCANCODE_OUT
	Keyp                  = C.SDLK_p                  // "P" SDL_SCANCODE_P
	KeyPageDown           = C.SDLK_PAGEDOWN           // "PageDown" SDL_SCANCODE_PAGEDOWN
	KeyPageUp             = C.SDLK_PAGEUP             // "PageUp" SDL_SCANCODE_PAGEUP
	KeyPaste              = C.SDLK_PASTE              // "Paste" SDL_SCANCODE_PASTE
	KeyPause              = C.SDLK_PAUSE              // "Pause" (the Pause / Break key) SDL_SCANCODE_PAUSE
	KeyPeriod             = C.SDLK_PERIOD             // "." SDL_SCANCODE_PERIOD
	KeyPower              = C.SDLK_POWER              // "Power" (The USB document says this is a status flag, not a physical key - but some Mac keyboards do have a power key.) SDL_SCANCODE_POWER
	KeyPrintScreen        = C.SDLK_PRINTSCREEN        // "PrintScreen" SDL_SCANCODE_PRINTSCREEN
	KeyPrior              = C.SDLK_PRIOR              // "Prior" SDL_SCANCODE_PRIOR
	Keyq                  = C.SDLK_q                  // "Q" SDL_SCANCODE_Q
	Keyr                  = C.SDLK_r                  // "R" SDL_SCANCODE_R
	KeyRAlt               = C.SDLK_RALT               // "Right Alt" (alt gr, option) SDL_SCANCODE_RALT
	KeyRCtrl              = C.SDLK_RCTRL              // "Right Ctrl" SDL_SCANCODE_RCTRL
	KeyReturn             = C.SDLK_RETURN             // "Return" (the Enter key (main keyboard)) SDL_SCANCODE_RETURN
	KeyReturn2            = C.SDLK_RETURN2            // "Return" SDL_SCANCODE_RETURN2
	KeyRGUI               = C.SDLK_RGUI               // "Right GUI" (windows, command (apple), meta) SDL_SCANCODE_RGUI
	KeyRight              = C.SDLK_RIGHT              // "Right" (the Right arrow key (navigation keypad)) SDL_SCANCODE_RIGHT
	KeyRightBracket       = C.SDLK_RIGHTBRACKET       // "]" SDL_SCANCODE_RIGHTBRACKET
	KeyRShift             = C.SDLK_RSHIFT             // "Right Shift" SDL_SCANCODE_RSHIFT
	Keys                  = C.SDLK_s                  // "S" SDL_SCANCODE_S
	KeyScrollLock         = C.SDLK_SCROLLLOCK         // "ScrollLock" SDL_SCANCODE_SCROLLLOCK
	KeySelect             = C.SDLK_SELECT             // "Select" SDL_SCANCODE_SELECT
	KeySemicolon          = C.SDLK_SEMICOLON          // ";" SDL_SCANCODE_SEMICOLON
	KeySeparator          = C.SDLK_SEPARATOR          // "Separator" SDL_SCANCODE_SEPARATOR
	KeySlash              = C.SDLK_SLASH              // "/" SDL_SCANCODE_SLASH
	KeySleep              = C.SDLK_SLEEP              // "Sleep" (the Sleep key) SDL_SCANCODE_SLEEP
	KeySpace              = C.SDLK_SPACE              // "Space" (the Space Bar key(s)) SDL_SCANCODE_SPACE
	KeyStop               = C.SDLK_STOP               // "Stop" SDL_SCANCODE_STOP
	KeySysReq             = C.SDLK_SYSREQ             // "SysReq" (the SysReq key) SDL_SCANCODE_SYSREQ
	Keyt                  = C.SDLK_t                  // "T" SDL_SCANCODE_T
	KeyTab                = C.SDLK_TAB                // "Tab" (the Tab key) SDL_SCANCODE_TAB
	KeyThousandsSeparator = C.SDLK_THOUSANDSSEPARATOR // "ThousandsSeparator" (the Thousands Separator key) SDL_SCANCODE_THOUSANDSSEPARATOR
	Keyu                  = C.SDLK_u                  // "U" SDL_SCANCODE_U
	KeyUndo               = C.SDLK_UNDO               // "Undo" SDL_SCANCODE_UNDO
	KeyUnknown            = C.SDLK_UNKNOWN            // "" (no name, empty string) SDL_SCANCODE_UNKNOWN
	KeyUp                 = C.SDLK_UP                 // "Up" (the Up arrow key (navigation keypad)) SDL_SCANCODE_UP
	Keyv                  = C.SDLK_v                  // "V" SDL_SCANCODE_V
	KeyVolumeDown         = C.SDLK_VOLUMEDOWN         // "VolumeDown" SDL_SCANCODE_VOLUMEDOWN
	KeyVolumeUp           = C.SDLK_VOLUMEUP           // "VolumeUp" SDL_SCANCODE_VOLUMEUP
	Keyw                  = C.SDLK_w                  // "W" SDL_SCANCODE_W
	KeyWWW                = C.SDLK_WWW                // "WWW" (the WWW/World Wide Web key) SDL_SCANCODE_WWW
	Keyx                  = C.SDLK_x                  // "X" SDL_SCANCODE_X
	Keyy                  = C.SDLK_y                  // "Y" SDL_SCANCODE_Y
	Keyz                  = C.SDLK_z                  // "Z" SDL_SCANCODE_Z

	// "" (no name, empty string; used on Asian keyboards, see footnotes in USB doc) SDL_SCANCODE_INTERNATIONAL1 (none)
	// "" (no name, empty string) SDL_SCANCODE_INTERNATIONAL2 (none)
	// "" (no name, empty string; Yen) SDL_SCANCODE_INTERNATIONAL3 (none)
	// "" (no name, empty string) SDL_SCANCODE_INTERNATIONAL4 (none)
	// "" (no name, empty string) SDL_SCANCODE_INTERNATIONAL5 (none)
	// "" (no name, empty string) SDL_SCANCODE_INTERNATIONAL6 (none)
	// "" (no name, empty string) SDL_SCANCODE_INTERNATIONAL7 (none)
	// "" (no name, empty string) SDL_SCANCODE_INTERNATIONAL8 (none)
	// "" (no name, empty string) SDL_SCANCODE_INTERNATIONAL9 (none)
	// "" (no name, empty string; Hangul/English toggle) SDL_SCANCODE_LANG1 (none)
	// "" (no name, empty string; Hanja conversion) SDL_SCANCODE_LANG2 (none)
	// "" (no name, empty string; Katakana) SDL_SCANCODE_LANG3 (none)
	// "" (no name, empty string; Hiragana) SDL_SCANCODE_LANG4 (none)
	// "" (no name, empty string; Zenkaku/Hankaku) SDL_SCANCODE_LANG5 (none)
	// "" (no name, empty string; reserved) SDL_SCANCODE_LANG6 (none)
	// "" (no name, empty string; reserved) SDL_SCANCODE_LANG7 (none)
	// "" (no name, empty string; reserved) SDL_SCANCODE_LANG8 (none)
	// "" (no name, empty string; reserved) SDL_SCANCODE_LANG9 (none)
	// "" (no name, empty string) SDL_SCANCODE_LOCKINGCAPSLOCK (none)
	// "" (no name, empty string) SDL_SCANCODE_LOCKINGNUMLOCK (none)
	// "" (no name, empty string) SDL_SCANCODE_LOCKINGSCROLLLOCK (none)
	// "" (no name, empty string; This is the additional key that ISO keyboards have over ANSI ones, located between left shift and Y. Produces GRAVE ACCENT and TILDE in a US or UK Mac layout, REVERSE SOLIDUS (backslash) and VERTICAL LINE in a US or UK Windows layout, and LESS-THAN SIGN and GREATER-THAN SIGN in a Swiss German, German, or French layout.) SDL_SCANCODE_NONUSBACKSLASH (none)
	// "#" (ISO USB keyboards actually use this code instead of 49 for the same key, but all OSes I've seen treat the two codes identically. So, as an implementor, unless your keyboard generates both of those codes and your OS treats them differently, you should generate SDL_SCANCODE_BACKSLASH instead of this code. As a user, you should not rely on this code because SDL will never generate it with most (all?) keyboards.) SDL_SCANCODE_NONUSHASH (none)

	KeyAmpersand  = C.SDLK_AMPERSAND  // "&" (none)
	KeyAsterisk   = C.SDLK_ASTERISK   // "*" (none)
	KeyAt         = C.SDLK_AT         // "@" (none)
	KeyCaret      = C.SDLK_CARET      // "^" (none)
	KeyColon      = C.SDLK_COLON      // ":" (none)
	KeyDollar     = C.SDLK_DOLLAR     // "$" (none)
	KeyExclaim    = C.SDLK_EXCLAIM    // "!" (none)
	KeyGreater    = C.SDLK_GREATER    // ">" (none)
	KeyHash       = C.SDLK_HASH       // "#" (none)
	KeyLeftParen  = C.SDLK_LEFTPAREN  // "(" (none)
	KeyLess       = C.SDLK_LESS       // "<" (none)
	KeyPercent    = C.SDLK_PERCENT    // "%" (none)
	KeyPlus       = C.SDLK_PLUS       // "+" (none)
	KeyQuestion   = C.SDLK_QUESTION   // "?" (none)
	KeyQuoteDbl   = C.SDLK_QUOTEDBL   // """ (none)
	KeyRightParen = C.SDLK_RIGHTPAREN // ")" (none)
	KeyUnderscore = C.SDLK_UNDERSCORE // "_" (none)
)
