package kn5

import "io"

import "fmt"

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

func newShaderProperty() *ShaderProperty {
	return &ShaderProperty{}
}

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

type ShaderProperties []*ShaderProperty

func (sp *ShaderProperties) Read(r io.Reader) error {

	count, err := ReadInt32(r)
	if err != nil {
		return err
	}

	if count == 0 {
		return nil
	}

	for i := int32(0); i < count; i++ {
		prop := newShaderProperty()
		if err := prop.Read(r); err != nil {
			return err
		}
		*sp = append(*sp, prop)
	}

	return nil
}
