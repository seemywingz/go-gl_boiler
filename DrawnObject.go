package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// DrawnObject : interface for opengl drawable object
type DrawnObject interface {
	Draw()
}

// DrawnObjectData : a struct to hold openGL object data
type DrawnObjectData struct {
	Vao     uint32
	Program uint32
	Points  []float32
	Position
	Translation int32
	Rotation    int32
	Model       mgl32.Mat4
}

// New : Create new DrawnObjectData
func (DrawnObjectData) New(position Position, points []float32, program uint32) *DrawnObjectData {

	ptrT, freeT := gl.Strs("translation")
	defer freeT()
	transloc := gl.GetUniformLocation(program, *ptrT)

	ptrR, freeR := gl.Strs("rotation")
	defer freeR()
	rotloc := gl.GetUniformLocation(program, *ptrR)

	return &DrawnObjectData{
		makeVao(points),
		program,
		points,
		position,
		transloc,
		rotloc,
		mgl32.Ident4(),
	}
}

var n float32

// Draw : draw the vertecies
func (d *DrawnObjectData) Draw() {
	gl.UseProgram(d.Program)
	gl.BindVertexArray(d.Vao)

	// translateMatrix := mgl32.Translate3D(d.X, d.Y, d.Z)
	// model := translateMatrix.Mul4(d.Model)
	// gl.UniformMatrix4fv(d.Translation, 1, false, &model[0])

	gl.Uniform4f(d.Translation, d.X, d.Y, d.Z, 1.0)

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(d.Points)/3))
}
