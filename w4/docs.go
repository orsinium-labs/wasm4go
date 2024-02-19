// # Callbacks
//
// The are 2 callbacks that you can register by assigning a function to them:
//
//   - [Start] is called once at the very beginning.
//   - [Update] is called on each frame update.
//
// # Drawing functions
//
// The library provides the following functions
// directly calling the host functions in WASM-4:
//
//   - [Blit]: copy raw sprite bytes from memory onto the frame buffer.
//   - [BlitSub]: like [Blit] but copies only the given region of a sprite.
//   - [DrawLine]: draw a line between two given points.
//   - [DrawHorLine]: draw horizontal line.
//   - [DrawVertLine]: draw vertical line.
//   - [DrawEllipse]: draw an ellipse (aka oval) or a circle.
//   - [DrawRect]: draw a rectangle (or square)
//   - [DrawText]: draw text using the built-in 8x8 font.
//   - [PlayTone]: play a sound.
//   - [Trace]: write a log message into the console.
//   - [Save]: write bytes into a persistent storage (save the game).
//   - [Load]: read bytes from the persistent storage (load the game).
//
// # Memory access
//
// Unlike the default WASM-4 bindings, the library doesn't require direct
// memory manipulation. Instead, there is a set of wrapper singletones
// for accessing specific regions of memory:
//
//   - [Palette] defines the 4 colors used to render the next frame.
//   - [DrawColors] defines which colors from the palette the draw functions should use.
//   - [Gamepads] lets you check which buttons on the gamepads are pressed.
//   - [Mouse] provides information about the mouse cursor position and pressed mouse buttons.
//   - [SystemFlags] allows modifying the WASM-4 behavior.
//   - [NetPlay] provides information about multiplayer.
//   - [FrameBuffer] provides direct access to the frame buffer in memory.
//
// Methods available for each of these singletones aren't shown in the web documetation
// because their types are private. Use your IDE to see available methods and their docs.
//
// # Types
//
// The library defines the following helper types for working with memory:
//
//   - [Color] is used to set colors for the [Palette].
//   - [DrawColor] is used by [DrawColors] to map draw functions colors to the palette.
package w4
