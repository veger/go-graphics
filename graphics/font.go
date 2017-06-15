package graphics

import (
	"fmt"
	"image"
	"image/draw"
	"io"
	"io/ioutil"
	"log"

	mgl "github.com/go-gl/mathgl/mgl32"
	"github.com/golang/freetype/truetype"
	"github.com/veger/go-graphics/utils"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// Font provides means to draw text on the screen
//
// It is (currently) intended for UI font (fixed scale, front facing, etc.)
type Font struct {
	engine             *Engine
	face               font.Face
	program            *Program
	recorder           *Recorder
	vbo                *Buffer
	ebo                *ElementBuffer
	projection         mgl.Mat4
	projectionLocation UniformLocation

	size int
}

// NewFont creates a Font instance
func (e *Engine) NewFont(ttf io.Reader, size int) (*Font, error) {
	fData, err := ioutil.ReadAll(ttf)
	if err != nil {
		return nil, fmt.Errorf("could not read font data:: %s", err)
	}

	ttFont, err := truetype.Parse(fData)
	if err != nil {
		return nil, fmt.Errorf("could not parse truetype font: %s", err)
	}

	face := truetype.NewFace(ttFont, &truetype.Options{
		Size:    float64(size),
		DPI:     180,
		Hinting: font.HintingNone,
	})

	program, err := e.NewProgram(vertexShaderFont, fragmentShaderFont)
	if err != nil {
		return nil, err
	}

	recorder, err := e.NewRecorder()
	if err != nil {
		program.Release()
		return nil, err
	}
	recorder.Activate()
	defer recorder.Deactivate()

	// TODO We should use GL_DYNAMIC_DRAW for buffer usage
	vbo, err := e.NewEmptyBuffer(BufferTypeData, 4*4*FloatSize)
	if err != nil {
		recorder.Release()
		program.Release()
		return nil, err
	}

	indicesData := []TriangleIndices{
		TriangleIndices{0, 1, 2},
		TriangleIndices{0, 2, 3},
	}
	ebo, err := e.NewElementBuffer(indicesData)
	if err != nil {
		vbo.Release()
		recorder.Release()
		program.Release()
		return nil, err
	}

	vbo.BindVariable(vertexLayout, 4, 4, 0)
	projectionLocation := program.GetUniformLocation("projection")

	recorder.Deactivate()

	font := &Font{
		engine:             e,
		recorder:           recorder,
		face:               face,
		program:            program,
		vbo:                vbo,
		ebo:                ebo,
		projectionLocation: projectionLocation,

		size: size,
	}

	// TODO Get rid of fixed width/height
	font.UpdateViewport(600, 800)
	e.RegisterViewportUpdater(font)

	return font, nil
}

// Release releases all resources, the font cannot be used afterwards
func (f *Font) Release() {
	f.engine.DeregisterViewportUpdater(f)
	f.ebo.Release()
	f.ebo = nil
	f.vbo.Release()
	f.vbo = nil
	f.recorder.Release()
	f.recorder = nil
	f.program.Release()
	f.program = nil
	f.face = nil
	f.engine = nil
}

// UpdateViewport recalculates the projection matrix based on the given viewport dimensions
func (f *Font) UpdateViewport(width, height int) {
	// Calculate projection
	f.projection = mgl.Ortho(float32(-width), float32(width), float32(-height), float32(height), 0, 1)

	// Udpate projection with view
	// pos := mgl.Vec3{width, height, 0}
	// dir := mgl.Vec3{width, height, -1}
	// up := mgl.Vec3{0, 1, 0}
	// view := mgl.LookAtV(pos, dir, up)
	view := mgl.Ident4()
	view[12] = float32(-width)
	view[13] = float32(-height)
	f.projection = f.projection.Mul4(view)
}

// Draw renders text at the given coordinates
func (f *Font) Draw(x, y float32, str string) {
	f.face.Metrics()
	// Render text to image
	d := &font.Drawer{
		Src:  image.White,
		Face: f.face,
		Dot: fixed.Point26_6{
			X: fixed.I(0),
			Y: fixed.I(f.size),
		},
	}
	b, _ := d.BoundString(str)

	w := b.Max.X.Ceil() - b.Min.X.Floor() + 1
	h := b.Max.Y.Ceil() - b.Min.Y.Floor() + 1
	im := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(im, im.Bounds(), image.Black, image.ZP, draw.Src)
	d.Dst = im
	d.Dot.Y = b.Max.Y - b.Min.Y
	d.DrawString(str)

	// Activate OpenGL stuff
	f.program.Activate()
	f.recorder.Activate()
	defer f.recorder.Deactivate()

	// Create texture and upload
	tex := f.engine.NewTexture(im)
	defer tex.Release()
	tex.Bind()
	defer tex.Unbind()

	// Update vertices to create surface to render the texture
	vertices := [4 * 4]float32{ // 4 * vec4 < vec2 position, vec2 tex>
		x, y + float32(h), 0.0, 0.0,
		x, y, 0.0, 1.0,
		x + float32(w), y, 1.0, 1.0,
		x + float32(w), y + float32(h), 1.0, 0.0,
	}

	data, err := utils.CastToBytes(vertices)
	if err != nil {
		log.Fatalln("CastToBytes failed:", err)
	}
	f.vbo.SetData(0, data)

	f.program.SetUnitformMatrix4(f.projectionLocation, f.projection[:])
	f.ebo.Render()
}

const vertexLayout = 0

const vertexShaderFont = `#version 300 es

layout (location = 0) in vec4 vertex; // <vec2 position, vec2 tex>
out vec2 TexCoord;

uniform mat4 projection;

void main() {
  gl_Position = projection * vec4(vertex.xy, 0.0, 1.0);
  TexCoord = vertex.zw;
}`

const fragmentShaderFont = `#version 300 es
precision mediump float;

in vec2 TexCoord;
out vec4 color;

uniform sampler2D text;

void main()
{
    color = vec4(1.0, 1.0, 1.0, 1.0) * texture(text, TexCoord).r;
}`
