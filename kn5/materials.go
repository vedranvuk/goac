package kn5

import (
	"fmt"
	"io"
)

type Material struct {
	Name        string
	ShaderName  string
	BlendMode   MaterialBlendMode
	AlphaTested bool
	DepthMode   MaterialDepthMode
	ShaderProperties
	TextureMappings
}

func newMaterial() *Material { return &Material{} }

func (m *Material) String() string {
	return fmt.Sprintf(`Material:
Name:             %s
ShaderName:       %s
BlendMode:        %v
AlphaTested:      %t
DepthMode:        %v
ShaderProperties: %v
TextureMappings:  %v
`,
		m.Name, m.ShaderName, m.BlendMode, m.AlphaTested, m.DepthMode, m.ShaderProperties, m.TextureMappings)
}

func (m *Material) Read(r io.Reader) error {
	// Name.
	name, err := ReadString(r)
	if err != nil {
		return err
	}
	m.Name = name
	// ShaderName.
	shadername, err := ReadString(r)
	if err != nil {
		return err
	}
	m.ShaderName = shadername
	// BlendMode.
	if err := m.BlendMode.Read(r); err != nil {
		return err
	}
	// AlphaTested.
	alphatested, err := ReadBoolean(r)
	if err != nil {
		return err
	}
	m.AlphaTested = alphatested
	// DepthMode.
	if err := m.DepthMode.Read(r); err != nil {
		return err
	}
	// ShaderProperties.
	if err := m.ShaderProperties.Read(r); err != nil {
		return err
	}
	// TextureMappings.
	if err := m.TextureMappings.Read(r); err != nil {
		return err
	}
	return nil
}

func (m *Material) Write(w io.Writer) error {
	// Name.
	if err := WriteString(w, m.Name); err != nil {
		return err
	}
	// ShaderName.
	if err := WriteString(w, m.ShaderName); err != nil {
		return err
	}
	// BlendMode.
	if err := m.BlendMode.Write(w); err != nil {
		return nil
	}
	// ALphaTested.
	if err := WriteBoolean(w, m.AlphaTested); err != nil {
		return err
	}
	// DepthMode.
	if err := m.DepthMode.Write(w); err != nil {
		return err
	}
	// ShaderProperties.
	if err := m.ShaderProperties.Write(w); err != nil {
		return err
	}
	// TextureMappings.
	if err := m.TextureMappings.Write(w); err != nil {
		return err
	}
	return nil
}

type Materials map[string]*Material

func newMaterials() *Materials {
	m := Materials(make(map[string]*Material))
	return &m
}

func (m *Materials) Read(r io.Reader) error {
	// Count.
	count, err := ReadInt32(r)
	if err != nil {
		return err
	}
	// No data.
	if count == 0 {
		return nil
	}
	// Materials.
	for i := int32(0); i < count; i++ {
		mat := newMaterial()
		if err := mat.Read(r); err != nil {
			return err
		}
		map[string]*Material(*m)[mat.Name] = mat
	}
	return nil
}

func (m *Materials) Write(w io.Writer) error {
	// Count.
	if err := WriteInt32(w, int32(len(*m))); err != nil {
		return err
	}
	// Materials.
	for _, mat := range *m {
		if err := mat.Write(w); err != nil {
			return err
		}
	}
	return nil
}
