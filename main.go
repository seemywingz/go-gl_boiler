package main

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	gt "github.com/seemywingz/gtils"
)

func main() {
	runtime.LockOSThread()
	gt.SetDirPath("github.com/seemywingz/go-gl_boiler")

	var windowWidth = 800
	var windowHeight = 600
	window = initGlfw(windowWidth, windowHeight, "go-gl Boiler")
	defer glfw.Terminate()

	initGL()
	loadShaders()
	loadTextures()
	loadLights()

	camera = Camera{}.New(Position{0, 0, 0}, false)

	randObject(200, -200, 200, cube, texture["box"], shader["basic"])
	randObject(200, -200, 200, cube, texture["box"], shader["texture"])
	randObject(200, -200, 200, cube, texture["box"], shader["phong"])
	drawnObjects = append(drawnObjects, Card{}.New(Position{0, 0, -5}))

	for !window.ShouldClose() {
		camera.Update()
		update()
	}
}

func randObject(numberOfObjects, min, max int, points []float32, textr, shadr uint32) {
	for i := 0; i < numberOfObjects; i++ {
		rand.Seed(time.Now().UnixNano())
		x := float32(rand.Intn(max-min) + min)
		y := float32(rand.Intn(max-min) + min)
		z := float32(rand.Intn(max-min) + min)
		d := DrawnObjectData{}.New(Position{x, y, z}, points, textr, shadr)
		d.DrawLogic = func(d *DrawnObjectData) {
			d.XRotation++
			d.YRotation++
		}
		drawnObjects = append(drawnObjects, d)
	}
}

func loadLights() {
}

func loadShaders() {
	shader = make(map[string]uint32)
	shader["basic"] = createGLprogram("shaders/basicVect.glsl", "shaders/basicFrag.glsl")
	shader["texture"] = createGLprogram("shaders/textureVect.glsl", "shaders/textureFrag.glsl")
	shader["phong"] = createGLprogram("shaders/blinnPhongVect.glsl", "shaders/blinnPhongFrag.glsl")
}

func loadTextures() {
	texture = make(map[string]uint32)
	texture["lifion"] = newTexture("textures/lifion.png")
	texture["box"] = newTexture("textures/square.jpg")
	texture["tk"] = newTexture("textures/tk.jpg")
	texture["back"] = newTexture("textures/back.jpg")
}

func update() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, obj := range drawnObjects {
		obj.Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
