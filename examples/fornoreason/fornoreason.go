package main

import (
	"math/rand"
	"time"

	"github.com/seemywingz/gg"
)

var (
	texture      map[string]uint32
	sceneObjects []*gg.DrawnObject
)

func randObjects(numberOfObjects, min, max int, points []float32, textr, shadr uint32) {
	for i := 0; i < numberOfObjects; i++ {
		var color []float32
		rand.Seed(time.Now().UnixNano())
		x := float32(rand.Intn(max-min) + min)
		y := float32(rand.Intn(max-min) + min)
		z := float32(rand.Intn(max-min) + min)
		if textr != gg.NoTexture {
			color = []float32{1, 1, 1}
		} else {
			color = []float32{
				rand.Float32(),
				rand.Float32(),
				rand.Float32(),
			}
		}
		d := gg.NewPointsObject(gg.NewPosition(x, y, z), points, textr, color, shadr)
		sceneObjects = append(sceneObjects, d)
	}
}

func loadTextures() {
	texture = make(map[string]uint32)
	texture["none"] = gg.NoTexture
	gg.SetDirPath("github.com/seemywingz/gg/examples/assets/textures")
	texture["box"] = gg.NewTexture("box.jpg")
}

func main() {

	gg.Init(800, 600, "Good Game")
	gg.SetCameraPosition(gg.NewPosition(0, 5, 100))
	gg.Enable(gg.PointerLock, true)
	gg.Enable(gg.FlyMode, true)
	light := gg.NewLight()
	light.Draw = true

	loadTextures()
	min, max := -30, 30
	randObjects(200, min, max, gg.Cube, texture["none"], gg.Shader["phong"])
	randObjects(700, min, max, gg.Cube, texture["box"], gg.Shader["phong"])

	for !gg.ShouldClose() {
		gg.Update()

		for _, obj := range sceneObjects {
			obj.Draw()
		}

		gg.SwapBuffers()
	}
}
