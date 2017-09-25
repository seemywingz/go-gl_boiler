package main

import (
	"github.com/seemywingz/gg"
)

func main() {

	var objects []*gg.DrawnObject

	gg.Init(800, 600, "Wavefront Loader")
	gg.SetCameraPosition(gg.NewPosition(0, 0.5, 2))

	l := gg.NewLight()
	l.Position = gg.NewPosition(-10, 10, 10)
	l.Radius = 30
	// l.Draw = true

	gg.SetDirPath("github.com/seemywingz/gg/examples/assets")
	mesh := gg.LoadObject("models/buddha.obj")
	obj := gg.NewMeshObject(gg.Position{}, mesh, gg.NewTexture("textures/buddha.jpg"), gg.Shader["phong"])
	obj.SceneLogic = func(s *gg.SceneData) {
		s.YRotation++
	}
	objects = append(objects, obj)

	box := gg.NewPointsObject(gg.Position{}, gg.Cube, gg.NewTexture("textures/box.jpg"), gg.Shader["phong"])
	box.Scale = 0.1
	objects = append(objects, box)

	for !gg.ShouldClose() {
		gg.Update()
		for _, o := range objects {
			o.Draw()
		}
		gg.SwapBuffers()
	}
}