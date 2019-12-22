package kn5

import (
	"fmt"
	"io"
)

type NodeClass int32

func (nc NodeClass) String() string {
	switch nc {
	case Base:
		return "Dummy"
	case Mesh:
		return "Mesh"
	case SkinnedMesh:
		return "Skinned mesh"
	}
	return "INVALID"
}

func (nc *NodeClass) Read(r io.Reader) error {
	val, err := ReadInt32(r)
	if err != nil {
		return err
	}
	*nc = NodeClass(val)
	return nil
}

func (nc *NodeClass) Write(w io.Writer) error {
	return WriteInt32(w, int32(*nc))
}

const (
	Base NodeClass = iota + 1
	Mesh
	SkinnedMesh
)

type Vertex struct {
	Position []float32
	Normal   []float32
	TexC     []float32
	TangentU []float32
}

func newVertex() *Vertex { return &Vertex{} }

func (v *Vertex) String() string {
	return fmt.Sprintf(`Vertex:
Position: %v
Normal:   %v
TexC:     %v
TangentU: %v
`, v.Position, v.Normal, v.TexC, v.TangentU)
}

func (v *Vertex) Read(r io.Reader) error {
	// Position.
	val, err := ReadFloats(r, 3)
	if err != nil {
		return err
	}
	v.Position = val
	// Normal.
	val, err = ReadFloats(r, 3)
	if err != nil {
		return err
	}
	v.Normal = val
	// TexC.
	val, err = ReadFloats(r, 2)
	if err != nil {
		return err
	}
	v.TexC = val
	// TangentU.
	val, err = ReadFloats(r, 3)
	if err != nil {
		return err
	}
	v.TangentU = val
	return nil
}

func (v *Vertex) Write(w io.Writer) error {
	// Position.
	if err := WriteFloats(w, v.Position); err != nil {
		return err
	}
	// Normal.
	if err := WriteFloats(w, v.Normal); err != nil {
		return err
	}
	// TexC.
	if err := WriteFloats(w, v.TexC); err != nil {
		return err
	}
	// TangentU.
	if err := WriteFloats(w, v.TangentU); err != nil {
		return err
	}
	return nil
}

type VerticeWeight struct {
	Weights []float32
	Indices []float32
}

func newVerticeWeight() *VerticeWeight {
	return &VerticeWeight{}
}

func (vw *VerticeWeight) String() string {
	return fmt.Sprintf(`VerticeWeight:
Weights: %v
Indices: %v
`, vw.Weights, vw.Indices)
}

func (vw *VerticeWeight) Read(r io.Reader) error {
	// Weights.
	val, err := ReadFloats(r, 4)
	if err != nil {
		return err
	}
	vw.Weights = val
	// Indices.
	val, err = ReadFloats(r, 4)
	if err != nil {
		return err
	}
	vw.Indices = val
	return nil
}

func (vw *VerticeWeight) Write(w io.Writer) error {
	// Weights.
	if err := WriteFloats(w, vw.Weights); err != nil {
		return err
	}
	// Indices.
	if err := WriteFloats(w, vw.Indices); err != nil {
		return err
	}
	return nil
}

type Bone struct {
	Name      string
	Transform []float32
}

func newBone() *Bone { return &Bone{} }

func (b *Bone) String() string {
	return fmt.Sprintf(`Bone:
Name:      %s
Transform: %v
`, b.Name, b.Transform)
}

func (b *Bone) Read(r io.Reader) error {
	// Name.
	name, err := ReadString(r)
	if err != nil {
		return err
	}
	b.Name = name
	// Transform.
	tform, err := ReadMatrix(r)
	if err != nil {
		return err
	}
	b.Transform = tform
	return nil
}

func (b *Bone) Write(w io.Writer) error {
	// Name.
	if err := WriteString(w, b.Name); err != nil {
		return err
	}
	// Transform.
	if err := WriteMatrix(w, b.Transform); err != nil {
		return err
	}
	return nil
}

type Node struct {
	NodeClass
	Name   string
	Active bool

	// Base
	Transform []float32
	Children  []*Node

	// Mesh
	CastShadows   bool
	IsVisible     bool
	IsTransparent bool
	IsRenderable  bool

	Bones          []*Bone
	Vertices       []*Vertex
	VerticeWeights []*VerticeWeight
	Indices        []uint16

	MaterialID uint32
	Layer      uint32

	LodIn  float32
	LodOut float32

	BoundingSphereCenter []float32
	BoundingSphereRadius float32

	// Skinned mesh
	MisteryBytes [8]byte
}

func newNode() *Node { return &Node{} }

func (n *Node) String() string {
	return fmt.Sprintf(`Node:
Class: %v
Name: %s
Active: %t
Transform: %v
Children: %v
CastShadows: %t
IsVisible: %t
IsTransparent: %t
IsRenderable: %t
Bones: %v
Vertices: %v
VerticeWeights: %v
Indices: %v
MaterialID: %d
Layer: %d
LodIn: %f
LodOut: %f
BoundingSphereCenter: %v
BoundingSphereRadius: %f
MisteryBytes: %v`,
		n.NodeClass, n.Name, n.Active, n.Transform, n.Children, n.CastShadows, n.IsVisible, n.IsTransparent,
		n.IsRenderable, n.Bones, n.Vertices, n.VerticeWeights, n.Indices, n.MaterialID, n.Layer, n.LodIn, n.LodOut,
		n.BoundingSphereCenter, n.BoundingSphereRadius, n.MisteryBytes)
}

func (n *Node) Read(r io.Reader) error {
	// NodeClass.
	if err := n.NodeClass.Read(r); err != nil {
		return err
	}
	// Name.
	name, err := ReadString(r)
	if err != nil {
		return err
	}
	n.Name = name
	// Child count.
	childcount, err := ReadInt32(r)
	if err != nil {
		return err
	}
	n.Children = make([]*Node, childcount)
	// Active.
	active, err := ReadBoolean(r)
	if err != nil {
		return err
	}
	n.Active = active
	// NodeClass.
	switch n.NodeClass {
	case Base:
		matrix, err := ReadMatrix(r)
		if err != nil {
			return err
		}
		n.Transform = matrix
	case Mesh:
		// CastShadows.
		val, err := ReadBoolean(r)
		if err != nil {
			return err
		}
		n.CastShadows = val
		// Visible.
		val, err = ReadBoolean(r)
		if err != nil {
			return err
		}
		n.IsVisible = val
		// Transparent.
		val, err = ReadBoolean(r)
		if err != nil {
			return err
		}
		n.IsTransparent = val
		// Vertexes.
		vertexcount, err := ReadUint32(r)
		if err != nil {
			return err
		}
		n.Vertices = make([]*Vertex, vertexcount)
		for i := 0; uint32(i) < vertexcount; i++ {
			vtx := newVertex()
			if err := vtx.Read(r); err != nil {
				return err
			}
			n.Vertices[i] = vtx
		}
		// Indices.
		indicecount, err := ReadUint32(r)
		if err != nil {
			return err
		}
		n.Indices = make([]uint16, indicecount)
		for i := uint32(0); i < indicecount; i++ {
			indice, err := ReadUint16(r)
			if err != nil {
				return err
			}
			n.Indices[int(i)] = indice
		}
		// MaterialID.
		materialid, err := ReadUint32(r)
		if err != nil {
			return err
		}
		n.MaterialID = materialid
		// Layer.
		layer, err := ReadUint32(r)
		if err != nil {
			return err
		}
		n.Layer = layer
		// LodIn.
		lodin, err := ReadFloat(r)
		if err != nil {
			return err
		}
		n.LodIn = lodin
		// LodOut.
		lodout, err := ReadFloat(r)
		if err != nil {
			return err
		}
		n.LodOut = lodout
		// BoundingSphereCenter
		bsc, err := ReadFloats(r, 3)
		if err != nil {
			return err
		}
		n.BoundingSphereCenter = bsc
		// BoundingSphereRadius
		bsr, err := ReadFloat(r)
		if err != nil {
			return err
		}
		n.BoundingSphereRadius = bsr
		// IsRenderable.
		rendable, err := ReadBoolean(r)
		if err != nil {
			return err
		}
		n.IsRenderable = rendable
	case SkinnedMesh:
		// CastShadows.
		val, err := ReadBoolean(r)
		if err != nil {
			return err
		}
		n.CastShadows = val
		// Visible.
		val, err = ReadBoolean(r)
		if err != nil {
			return err
		}
		n.IsVisible = val
		// Transparent.
		val, err = ReadBoolean(r)
		if err != nil {
			return err
		}
		n.IsTransparent = val
		// Bones.
		bonecount, err := ReadUint32(r)
		if err != nil {
			return err
		}
		n.Bones = make([]*Bone, bonecount)
		for i := uint32(0); i < bonecount; i++ {
			bone := newBone()
			if err := bone.Read(r); err != nil {
				return err
			}
			n.Bones[i] = bone
		}
		// Vertices.
		vertexcount, err := ReadUint32(r)
		if err != nil {
			return err
		}
		n.Vertices = make([]*Vertex, vertexcount)
		n.VerticeWeights = make([]*VerticeWeight, vertexcount)
		for i := 0; uint32(i) < vertexcount; i++ {
			vtx := newVertex()
			if err := vtx.Read(r); err != nil {
				return err
			}
			n.Vertices[i] = vtx

			vwv := newVerticeWeight()
			if err := vwv.Read(r); err != nil {
				return err
			}
			n.VerticeWeights[i] = vwv
		}
		// Indices.
		indicecount, err := ReadUint32(r)
		if err != nil {
			return err
		}
		n.Indices = make([]uint16, indicecount)
		for i := uint32(0); i < indicecount; i++ {
			indice, err := ReadUint16(r)
			if err != nil {
				return err
			}
			n.Indices[int(i)] = indice
		}
		// MaterialID.
		materialid, err := ReadUint32(r)
		if err != nil {
			return err
		}
		n.MaterialID = materialid
		// Layer.
		layer, err := ReadUint32(r)
		if err != nil {
			return err
		}
		n.Layer = layer
		// MisteryBytes.
		if _, err := r.Read(n.MisteryBytes[:]); err != nil {
			return err
		}
		// IsRenderable.
		n.IsRenderable = true
	}
	// Read children.
	for i := 0; i < len(n.Children); i++ {
		node := newNode()
		if err := node.Read(r); err != nil {
			return err
		}
		n.Children[i] = node
	}
	return nil
}

func (n *Node) Write(w io.Writer) error {
	// NodeClass.
	if err := n.NodeClass.Write(w); err != nil {
		return err
	}
	// Name.
	if err := WriteString(w, n.Name); err != nil {
		return err
	}
	// Child count.
	if err := WriteInt32(w, int32(len(n.Children))); err != nil {
		return err
	}
	// Active.
	if err := WriteBoolean(w, n.Active); err != nil {
		return err
	}
	// NodeClass.
	switch n.NodeClass {
	case Base:
		if err := WriteMatrix(w, n.Transform); err != nil {
			return err
		}
	case Mesh:
		// CastShadows.
		if err := WriteBoolean(w, n.CastShadows); err != nil {
			return err
		}
		// Visible.
		if err := WriteBoolean(w, n.IsVisible); err != nil {
			return err
		}
		// Transparent.
		if err := WriteBoolean(w, n.IsTransparent); err != nil {
			return err
		}
		// Vertexes.
		if err := WriteUint32(w, uint32(len(n.Vertices))); err != nil {
			return err
		}
		for _, vtx := range n.Vertices {
			if err := vtx.Write(w); err != nil {
				return err
			}
		}
		// Indices.
		if err := WriteUint32(w, uint32(len(n.Indices))); err != nil {
			return err
		}
		for _, indice := range n.Indices {
			if err := WriteUint16(w, indice); err != nil {
				return err
			}
		}
		// MaterialID.
		if err := WriteUint32(w, n.MaterialID); err != nil {
			return err
		}
		// Layer.
		if err := WriteUint32(w, n.Layer); err != nil {
			return err
		}
		// LodIn.
		if err := WriteFloat(w, n.LodIn); err != nil {
			return err
		}
		// LodOut.
		if err := WriteFloat(w, n.LodOut); err != nil {
			return err
		}
		// BoundingSphereCenter
		if err := WriteFloats(w, n.BoundingSphereCenter); err != nil {
			return err
		}
		// BoundingSphereRadius
		if err := WriteFloat(w, n.BoundingSphereRadius); err != nil {
			return err
		}
		// IsRenderable.
		if err := WriteBoolean(w, n.IsRenderable); err != nil {
			return err
		}
	case SkinnedMesh:
		// CastShadows.
		if err := WriteBoolean(w, n.CastShadows); err != nil {
			return err
		}
		// Visible.
		if err := WriteBoolean(w, n.IsVisible); err != nil {
			return err
		}
		// Transparent.
		if err := WriteBoolean(w, n.IsTransparent); err != nil {
			return err
		}
		// Bones.
		if err := WriteUint32(w, uint32(len(n.Bones))); err != nil {
			return err
		}
		for _, bone := range n.Bones {
			if err := bone.Write(w); err != nil {
				return err
			}
		}
		// Vertices.
		if err := WriteUint32(w, uint32(len(n.Vertices))); err != nil {
			return err
		}
		for _, vtx := range n.Vertices {
			if err := vtx.Write(w); err != nil {
				return err
			}
		}
		// Indices.
		if err := WriteUint32(w, uint32(len(n.Indices))); err != nil {
			return err
		}
		for _, indice := range n.Indices {
			if err := WriteUint16(w, indice); err != nil {
				return err
			}
		}
		// MaterialID.
		if err := WriteUint32(w, n.MaterialID); err != nil {
			return err
		}
		// Layer.
		if err := WriteUint32(w, n.Layer); err != nil {
			return err
		}
		// MisteryBytes.
		if _, err := w.Write(n.MisteryBytes[:]); err != nil {
			return err
		}
		// IsRenderable.
		n.IsRenderable = true
	}
	// Write children.
	for _, child := range n.Children {
		if err := child.Write(w); err != nil {
			return err
		}
	}
	return nil
}
