# Veger Graphics Engine

This (golang) graphics engine is meant to provide an abstraction layer between an application/game engine and the (OpenGL) graphics backend/drivers.

This graphics abstraction engine has been written, because I am not aware of any abstraction layers for (OpenGL) graphics. The ones out there are either for a specific (device) type (for example [mobile/gl](https://godoc.org/golang.org/x/mobile/gl) or [go-gl/gl](https://github.com/go-gl/gl)) or provide a complete game engine that is (barely) not usage without using their complete engine (like [go-sld2](https://github.com/veandco/go-sdl2) and  [go-gl/fwgl](https://github.com/go-gl/glfw)). And of course it is ~~fun~~ educational to write the low-level OpenGL yourself!

## Supported graphics backends

OpenGL:

* [golang.org/x/mobile/gl](https://godoc.org/golang.org/x/mobile/gl) (requires 'vao' branch of [my fork](https://github.com/veger/mobile))

# Usage

The central part is the [_Engine_ struct](graphics/engine.go), it must be created first in order to make the rest of the functionalities available. This is done to prevent the need of passing around the graphical context around to all functionalities.

Using the engine, it is possible to create ([Element](graphics/element_buffer.go))[Buffers](graphics/buffer.go), [Programs](graphics/program.go) (OpenGL shader programs), [Recorders](graphics/recorder.go) (record graphics state for ease and efficiency) or [Textures](graphics/texture.go).

## Usage notice

It is a (very) early version of the engine and its interfaces still might change (a lot) as I seem fit for my applications.
On the other hand, (code) additions (e.g. [go-gl](https://github.com/go-gl/gl) support ;) ) are very welcome and together we can see how can make sure that they can be merged into the engine.

Long story short: Have fun with this, but use at own risk!
