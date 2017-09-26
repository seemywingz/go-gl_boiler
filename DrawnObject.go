package gg

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// DrawnObject : a struct to hold openGL object data
type DrawnObject struct {
	Mesh           *Mesh
	MVPID          int32
	ModelMatrixID  int32
	NormalMatrixID int32
	ColorID        int32
	Color          Color
	Texture        uint32
	Scale          float32
	SceneData
}

// NewPointsObject :
func NewPointsObject(position Position, points []float32, texture uint32, program uint32) *DrawnObject {
	vao := MakeVAO(points, program)
	mg := make(map[string]*MaterialGroup)
	mg["dfault"] = &MaterialGroup{
		&defaultMaterial,
		[]*Face{},
		vao,
		int32(len(points)),
	}
	mesh := &Mesh{mg}
	return NewMeshObject(position, mesh, texture, program)
}

// NewMeshObject : Create new DrawnObject
func NewMeshObject(position Position, mesh *Mesh, texture uint32, program uint32) *DrawnObject {

	ModelMatrixID := gl.GetUniformLocation(program, gl.Str("MODEL\x00"))
	NormalMatrixID := gl.GetUniformLocation(program, gl.Str("NormalMatrix\x00"))
	MVPID := gl.GetUniformLocation(program, gl.Str("MVP\x00"))
	ColorID := gl.GetUniformLocation(program, gl.Str("COLOR\x00"))

	d := &DrawnObject{
		mesh,
		MVPID,
		ModelMatrixID,
		NormalMatrixID,
		ColorID,
		NewColor(1, 1, 1, 1),
		texture,
		1,
		SceneData{},
	}
	d.Position = position
	d.Program = program
	return d
}

func (d *DrawnObject) translateRotate() *mgl32.Mat4 {
	model := mgl32.Translate3D(d.X, d.Y, d.Z).
		Mul4(mgl32.Scale3D(d.Scale, d.Scale, d.Scale))
	xrotMatrix := mgl32.HomogRotate3DX(mgl32.DegToRad(d.XRotation))
	yrotMatrix := mgl32.HomogRotate3DY(mgl32.DegToRad(d.YRotation))
	zrotMatrix := mgl32.HomogRotate3DZ(mgl32.DegToRad(d.ZRotation))
	final := model.Mul4(xrotMatrix.Mul4(yrotMatrix.Mul4(zrotMatrix)))
	return &final
}

// Draw : draw the object
func (d *DrawnObject) Draw() {

	if d.SceneLogic != nil {
		d.SceneLogic(&d.SceneData)
	}

	modelMatrix := d.translateRotate()
	normalMatrix := modelMatrix.Inv().Transpose()

	gl.UseProgram(d.Program)
	gl.UniformMatrix4fv(d.MVPID, 1, false, &camera.MVP[0])
	gl.UniformMatrix4fv(d.ModelMatrixID, 1, false, &modelMatrix[0])
	gl.UniformMatrix4fv(d.NormalMatrixID, 1, false, &normalMatrix[0])
	// gl.Uniform4f(d.ColorID, d.Color.R, d.Color.G, d.Color.B, d.Color.A)

	for _, m := range d.Mesh.MaterialGroups {
		gl.Uniform4f(d.ColorID, m.Material.Diffuse[0], m.Material.Diffuse[1], m.Material.Diffuse[2], m.Material.Diffuse[3])
		gl.BindVertexArray(m.VAO)
		if d.Texture != NoTexture {
			gl.Enable(gl.TEXTURE_2D)
			gl.BindTexture(gl.TEXTURE_2D, d.Texture)
		}

		gl.DrawArrays(gl.TRIANGLES, 0, m.VertCount)
		gl.BindTexture(gl.TEXTURE_2D, 0)
	}

}
