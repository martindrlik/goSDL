package sdl

// #include <SDL2/SDL_render.h>
import "C"

type PixelFormat uint

const (
	PixelFormatUnknown     PixelFormat = C.SDL_PIXELFORMAT_UNKNOWN
	PixelFormatIndex1LSB   PixelFormat = C.SDL_PIXELFORMAT_INDEX1LSB
	PixelFormatIndex1MSB   PixelFormat = C.SDL_PIXELFORMAT_INDEX1MSB
	PixelFormatIndex4LSB   PixelFormat = C.SDL_PIXELFORMAT_INDEX4LSB
	PixelFormatIndex4MSB   PixelFormat = C.SDL_PIXELFORMAT_INDEX4MSB
	PixelFormatIndex8      PixelFormat = C.SDL_PIXELFORMAT_INDEX8
	PixelFormatRGB332      PixelFormat = C.SDL_PIXELFORMAT_RGB332
	PixelFormatRGB444      PixelFormat = C.SDL_PIXELFORMAT_RGB444
	PixelFormatRGB555      PixelFormat = C.SDL_PIXELFORMAT_RGB555
	PixelFormatBGR555      PixelFormat = C.SDL_PIXELFORMAT_BGR555
	PixelFormatARGB4444    PixelFormat = C.SDL_PIXELFORMAT_ARGB4444
	PixelFormatRGBA4444    PixelFormat = C.SDL_PIXELFORMAT_RGBA4444
	PixelFormatABGR4444    PixelFormat = C.SDL_PIXELFORMAT_ABGR4444
	PixelFormatBGRA4444    PixelFormat = C.SDL_PIXELFORMAT_BGRA4444
	PixelFormatARGB1555    PixelFormat = C.SDL_PIXELFORMAT_ARGB1555
	PixelFormatRGBA5551    PixelFormat = C.SDL_PIXELFORMAT_RGBA5551
	PixelFormatABGR1555    PixelFormat = C.SDL_PIXELFORMAT_ABGR1555
	PixelFormatBGRA5551    PixelFormat = C.SDL_PIXELFORMAT_BGRA5551
	PixelFormatRGB565      PixelFormat = C.SDL_PIXELFORMAT_RGB565
	PixelFormatBGR565      PixelFormat = C.SDL_PIXELFORMAT_BGR565
	PixelFormatRGB24       PixelFormat = C.SDL_PIXELFORMAT_RGB24
	PixelFormatBGR24       PixelFormat = C.SDL_PIXELFORMAT_BGR24
	PixelFormatRGB888      PixelFormat = C.SDL_PIXELFORMAT_RGB888
	PixelFormatRGBX8888    PixelFormat = C.SDL_PIXELFORMAT_RGBX8888
	PixelFormatBGR888      PixelFormat = C.SDL_PIXELFORMAT_BGR888
	PixelFormatBGRX8888    PixelFormat = C.SDL_PIXELFORMAT_BGRX8888
	PixelFormatARGB8888    PixelFormat = C.SDL_PIXELFORMAT_ARGB8888
	PixelFormatRGBA8888    PixelFormat = C.SDL_PIXELFORMAT_RGBA8888
	PixelFormatABGR8888    PixelFormat = C.SDL_PIXELFORMAT_ABGR8888
	PixelFormatBGRA8888    PixelFormat = C.SDL_PIXELFORMAT_BGRA8888
	PixelFormatARGB2101010 PixelFormat = C.SDL_PIXELFORMAT_ARGB2101010
	PixelFormatRGBA32      PixelFormat = C.SDL_PIXELFORMAT_RGBA32 // alias for RGBA byte array of color data, for the current platform (>= SDL 2.0.5)
	PixelFormatARGB32      PixelFormat = C.SDL_PIXELFORMAT_ARGB32 // alias for ARGB byte array of color data, for the current platform (>= SDL 2.0.5)
	PixelFormatBGRA32      PixelFormat = C.SDL_PIXELFORMAT_BGRA32 // alias for BGRA byte array of color data, for the current platform (>= SDL 2.0.5)
	PixelFormatABGR32      PixelFormat = C.SDL_PIXELFORMAT_ABGR32 // alias for ABGR byte array of color data, for the current platform (>= SDL 2.0.5)
	PixelFormatYV12        PixelFormat = C.SDL_PIXELFORMAT_YV12   // planar mode: Y + V + U (3 planes)
	PixelFormatIYUV        PixelFormat = C.SDL_PIXELFORMAT_IYUV   // planar mode: Y + U + V (3 planes)
	PixelFormatYUY2        PixelFormat = C.SDL_PIXELFORMAT_YUY2   // packed mode: Y0+U0+Y1+V0 (1 plane)
	PixelFormatUYVY        PixelFormat = C.SDL_PIXELFORMAT_UYVY   // packed mode: U0+Y0+V0+Y1 (1 plane)
	PixelFormatYVYU        PixelFormat = C.SDL_PIXELFORMAT_YVYU   // packed mode: Y0+V0+Y1+U0 (1 plane)
	PixelFormatNV12        PixelFormat = C.SDL_PIXELFORMAT_NV12   // planar mode: Y + U/V interleaved (2 planes) (>= SDL 2.0.4)
	PixelFormatNV21        PixelFormat = C.SDL_PIXELFORMAT_NV21   // planar mode: Y + V/U interleaved (2 planes) (>= SDL 2.0.4)
)
