package graphics

import (
	"image"
	"image/draw"

	"golang.org/x/mobile/gl"
)

// Texture is a abstraction used to communicate between engine and graphical backend
type Texture struct {
	engine *Engine
	glTex  gl.Texture
}

// NewTexture creates an Texture from the provided image source
func (e *Engine) NewTexture(imgSrc image.Image) *Texture {
	b := imgSrc.Bounds()

	RGBA := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(RGBA, b, imgSrc, b.Min, draw.Src)

	glTex := e.glctx.CreateTexture()
	// TODO Make function
	e.glctx.BindTexture(gl.TEXTURE_2D, glTex)
	e.glctx.TexImage2D(gl.TEXTURE_2D, 0, b.Dx(), b.Dy(), gl.RGBA, gl.UNSIGNED_BYTE, nil)

	// Copy texture to graphic memory
	e.glctx.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, b.Dx(), b.Dy(), gl.RGBA, gl.UNSIGNED_BYTE, RGBA.Pix)

	// Set texture parameters
	e.glctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	e.glctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	e.glctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	e.glctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	e.glctx.BindTexture(gl.TEXTURE_2D, gl.Texture{Value: 0})

	return &Texture{engine: e, glTex: glTex}
}

// Release releases the texture
// Using the texture after calling this function results in a panic
func (t *Texture) Release() {
	t.engine.glctx.DeleteTexture(t.glTex)
	t.engine = nil
}

// Bind binds the texture to the graphical backend, so it can be used
// TODO Only one texture can be bound at a specitic time
func (t *Texture) Bind() {
	// TODO Check if already bound (prevent overhead?)
	t.engine.glctx.BindTexture(gl.TEXTURE_2D, t.glTex)
}

// Unbind unbinds the texture from the graphical backend, making sure it is not used (by accident) anymore
func (t *Texture) Unbind() {
	t.engine.glctx.BindTexture(gl.TEXTURE_2D, gl.Texture{Value: 0})
}
