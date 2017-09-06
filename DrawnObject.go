package main

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// DrawnObject : interface for opengl drawable object
type DrawnObject interface {
	Draw()
}

// Color : struct to store RGB colors as float32
type Color struct {
	R float32
	G float32
	B float32
}

// Position : struct to store 3D coords
type Position struct {
	X float32
	Y float32
	Z float32
}

// DrawnObjectData : a struct to hold openGL object data
type DrawnObjectData struct {
	Vao     uint32
	Program uint32
	MV      int32
	Points  []float32
	Position
	Color
}

// New : Create a new object
func (d DrawnObjectData) New(position Position, color Color, points []float32) DrawnObjectData {
	vertexShaderSource := readShaderFile("./shaders/vertex.glsl")
	fragmentShaderSource := readShaderFile("./shaders/fragment.glsl")

	program := createGLprogram(vertexShaderSource, fragmentShaderSource)

	gl.UseProgram(program)
	gl.BindVertexArray(d.Vao)

	// mvPointer, free := gl.Strs("MVP")
	// defer free()
	mvID := int32(0) //:= gl.GetUniformLocation(program, *mvPointer)

	return DrawnObjectData{
		makeVao(points),
		program,
		mvID,
		points,
		position,
		color,
	}

}

// Draw : draw the triangle
func (d DrawnObjectData) Draw() {
	gl.UseProgram(d.Program)
	gl.BindVertexArray(d.Vao)

	// mvPointer, free := gl.Strs("MVP")
	// defer free()
	// mvID := gl.GetUniformLocation(d.Program, *mvPointer)
	// d.MV = mvID
	fmt.Println("ModelView IDs: ", d.MV)
	gl.UniformMatrix4fv(d.MV, 1, false, &camera.MVP[0])

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(d.Points)/3))
}