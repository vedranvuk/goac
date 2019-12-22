package kn5

import (
	"fmt"
	"io"
)

type ShaderProperty struct {
	Name   string
	ValueA float32
	ValueB []float32
	ValueC []float32
	ValueD []float32
}

func (sp *ShaderProperty) String() string {
	return fmt.Sprintf(`ShaderProperty:
Name:   %s
ValueA: %f
ValueB: %v
ValueC: %v
ValueD: %v
`,
		sp.Name, sp.ValueA, sp.ValueB, sp.ValueC, sp.ValueD)
}

func newShaderProperty() *ShaderProperty { return &ShaderProperty{} }

func (sp *ShaderProperty) Read(r io.Reader) error {
	// Name.
	name, err := ReadString(r)
	if err != nil {
		return err
	}
	sp.Name = name
	// ValueA
	valA, err := ReadFloat(r)
	if err != nil {
		return err
	}
	sp.ValueA = valA
	// ValueB
	valB, err := ReadFloats(r, 2)
	if err != nil {
		return err
	}
	sp.ValueB = valB
	// ValueC
	valC, err := ReadFloats(r, 3)
	if err != nil {
		return err
	}
	sp.ValueC = valC
	// ValueD
	valD, err := ReadFloats(r, 4)
	if err != nil {
		return err
	}
	sp.ValueD = valD
	return nil
}

func (sp *ShaderProperty) Write(w io.Writer) error {
	// Name.
	if err := WriteString(w, sp.Name); err != nil {
		return err
	}
	// ValueA.
	if err := WriteFloat(w, sp.ValueA); err != nil {
		return err
	}
	// ValueB.
	if err := WriteFloats(w, sp.ValueB); err != nil {
		return err
	}
	// ValueC.
	if err := WriteFloats(w, sp.ValueC); err != nil {
		return err
	}
	// ValueD.
	if err := WriteFloats(w, sp.ValueD); err != nil {
		return err
	}
	return nil
}

type ShaderProperties []*ShaderProperty

func (sp *ShaderProperties) Read(r io.Reader) error {
	// Count.
	count, err := ReadInt32(r)
	if err != nil {
		return err
	}
	// No data.
	if count == 0 {
		return nil
	}
	// ShadersProperties.
	for i := int32(0); i < count; i++ {
		prop := newShaderProperty()
		if err := prop.Read(r); err != nil {
			return err
		}
		*sp = append(*sp, prop)
	}
	return nil
}

func (sp *ShaderProperties) Write(w io.Writer) error {
	// Count.
	if err := WriteInt32(w, int32(len(*sp))); err != nil {
		return err
	}
	// ShaderProperties.
	for _, prop := range *sp {
		if err := prop.Write(w); err != nil {
			return err
		}
	}
	return nil
}
