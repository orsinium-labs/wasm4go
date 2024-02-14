package w4

type u8 = uint8

// A point on the plane.
type Point struct {
	X u8
	Y u8
}

// Size of a 2D shape.
type Size struct {
	Width  u8
	Height u8
}

type BlitFlags u8

const (
	TwoBPP BlitFlags = 1
	FlipX  BlitFlags = 2
	FlipY  BlitFlags = 4
	Rotate BlitFlags = 8
)

type Tone struct {
	// Start wave frequency in hertz.
	StartFreq uint

	// End wave frequency in hertz, used to describe a pitch slide effect.
	EndFreq uint

	// Volume of the sustain duration, between 0 and 100.
	SustainVol u8

	// Volume of the attack duration, between 0 and 100.
	AttackVol u8

	Channel   Channel
	DutyCycle DutyCycle
	Pan       Pan

	// Duration of the tone in frames (1/60th of a second), up to 255 frames.
	// Sustain time of ADSR envelope.
	Sustain u8

	// Attack time of ADSR envelope.
	Attack u8

	// Decay time of ADSR envelope.
	Decay u8

	// Release time of ADSR envelope.
	Release u8
}

type Channel u8

const (
	Pulse1   Channel = 0
	Pulse2   Channel = 1
	Triangle Channel = 2
	Noise    Channel = 3
)

type DutyCycle u8

const (
	DutyCycle1p8 DutyCycle = 0
	DutyCycle1p4 DutyCycle = 4
	DutyCycle1p2 DutyCycle = 8
	DutyCycle3p4 DutyCycle = 12
)

type Pan u8

const (
	Center Pan = 0
	Left   Pan = 16
	Right  Pan = 32
)

type ToneFlags u8

const (
	TonePulse1   = 0
	TonePulse2   = 1
	ToneTriangle = 2
	ToneNoise    = 3
	ToneMode1    = 0
	ToneMode2    = 4
	ToneMode3    = 8
	ToneMode4    = 12
)

// Copies pixels to the framebuffer.
//
// https://wasm4.org/docs/reference/functions#blit-spriteptr-x-y-width-height-flags
func Blit(sprite []byte, p Point, s Size, f BlitFlags) {
	blit(&sprite[0], p.X, p.Y, s.Width, s.Height, u8(f))
}

// Copies a subregion within a larger sprite atlas to the framebuffer.
//
// https://wasm4.org/docs/reference/functions#blitsub-spriteptr-x-y-width-height-srcx-srcy-stride-flags
func BlitSub(sprite []byte, dst Point, s Size, src Point, stride u8, f BlitFlags) {
	blitSub(&sprite[0], dst.X, dst.Y, s.Width, s.Height, src.X, src.Y, stride, u8(f))
}

// Draws a line between two points.
//
// https://wasm4.org/docs/reference/functions#line-x1-y1-x2-y2
func DrawLine(p1, p2 Point) {
	line(p1.X, p1.Y, p2.X, p2.Y)
}

// Draws a horizontal line between (x, y) and (x + len - 1, y).
//
// https://wasm4.org/docs/reference/functions#hlinex-y-len
func DrawHorLine(p Point, len u8) {
	hLine(p.X, p.Y, len)
}

// Draws a vertical line between (x, y) and (x, y + len - 1).
//
// https://wasm4.org/docs/reference/functions#vlinex-y-len
func DrawVertLine(p Point, len u8) {
	vLine(p.X, p.Y, len)
}

// Draws an oval (or circle).
//
// The point is coordinates of the left-upper corner of the bounding box.
// The size is the size of the bounding box.
//
// https://wasm4.org/docs/reference/functions#oval-x-y-width-height
func DrawEllipse(p Point, s Size) {
	oval(p.X, p.Y, s.Width, s.Height)
}

// Draws a rectangle.
//
// https://wasm4.org/docs/reference/functions#rect-x-y-width-height
func DrawRect(p Point, s Size) {
	rect(p.X, p.Y, s.Width, s.Height)
}

// Draws text using the built-in system font.
//
// The string may contain new-line (\n) characters.
//
// The font is 8x8 pixels per character.
//
// https://wasm4.org/docs/reference/functions#text-str-x-y
func DrawText(text string, p Point) {
	textUtf8(text, p.X, p.Y)
}

// Play a sound tone.
//
// https://wasm4.org/docs/reference/functions#tone-frequency-duration-volume-flags
func PlayTone(t Tone) {
	flags := u8(t.Channel) | u8(t.DutyCycle) | u8(t.Pan)
	freq := (uint32(t.StartFreq) << 8) | uint32(t.EndFreq)
	vol := (uint32(t.AttackVol) << 8) | uint32(t.SustainVol)
	dur := uint32(t.Sustain)
	dur |= uint32(t.Release) << 8
	dur |= uint32(t.Decay) << 16
	dur |= uint32(t.Attack) << 24
	tone(freq, dur, vol, flags)
}

// Prints a message to the debug console.
func Trace(text string) {
	trace(text)
}

// Reads bytes from persistent storage into the buffer.
//
// Make sure the buffer has cap enough to fit the data.
func Load(buf []byte) uint {
	return diskR(&buf[0], len(buf))
}

// Writes bytes from the buffer into persistent storage.
func Save(buf []byte) uint {
	return diskW(&buf[0], len(buf))
}
