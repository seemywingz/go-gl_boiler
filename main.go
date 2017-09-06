package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 800
	height = 600
	title  = "go-gl Boiler"
)

var (
	camera       Camera
	shaders      []uint32
	window       *glfw.Window
	drawnObjects []DrawnObject

	triangle = []float32{
		0, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}
	square = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,

		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		0.5, -0.5, 0,
	}
)

func main() {
	runtime.LockOSThread()

	window = initGlfw(width, height, title)
	defer glfw.Terminate()

	initGL()
	loadShaders()

	camera = Camera{}.New(Position{0, 0, -10})

	for i := 0; i < 20; i++ {
		drawnObjects = append(drawnObjects, DrawnObjectData{}.New(Position{float32(i), 0, -10}, triangle, shaders[0]))
	}
	// drawnObjects = append(drawnObjects, DrawnObjectData{}.New(Position{1, 0, -50}, triangle, shaders[0]))

	for !window.ShouldClose() {
		camera.Update()
		draw()
	}
}

func loadShaders() {
	shaders = append(
		shaders,
		createGLprogram(
			readShaderFile("./shaders/vertex.glsl"),
			readShaderFile("./shaders/fragment.glsl"),
		))
}

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, obj := range drawnObjects {
		obj.Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
