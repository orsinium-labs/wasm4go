package w4

import "unsafe"

//go:export blit
func blit(sprite *byte, x, y, width, height, flags u8)

//go:export blitSub
func blitSub(sprite *byte, x, y, width, height, srcX, srcY, stride, flags u8)

//go:export line
func line(x1, y1, x2, y2 u8)

//go:export hline
func hLine(x, y, len u8)

//go:export vline
func vLine(x, y, len u8)

//go:export oval
func oval(x, y, width, height u8)

//go:export rect
func rect(x, y, width, height u8)

//go:export textUtf8
func textUtf8(text string, x, y u8)

//go:export tone
func tone(frequency, duration, volume uint32, flags u8)

//go:export diskr
func diskR(ptr unsafe.Pointer, count u8) u8

//go:export diskw
func diskW(src unsafe.Pointer, count u8) u8

//go:export traceUtf8
func trace(str string)
