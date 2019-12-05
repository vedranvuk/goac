package kn5

import "io"

import "fmt"

type TextureMapping struct {
	Name    string
	Texture string
	Slot    int32
}

func newTextureMapping() *TextureMapping {
	return &TextureMapping{}
}

func (tm *TextureMapping) String() string {
	return fmt.Sprintf(`TextureMapping:
Name:    %s
Texture: %s
Slot:    %d
`,
		tm.Name, tm.Texture, tm.Slot)
}

func (tm *TextureMapping) Read(r io.Reader) error {

	// Name.
	name, err := ReadString(r)
	if err != nil {
		return err
	}
	tm.Name = name

	// Slot.
	slot, err := ReadInt32(r)
	if err != nil {
		return err
	}
	tm.Slot = slot

	// Texture.
	texture, err := ReadString(r)
	if err != nil {
		return err
	}
	tm.Texture = texture

	return nil
}

type TextureMappings []*TextureMapping

func (tm *TextureMappings) Read(r io.Reader) error {

	count, err := ReadInt32(r)
	if err != nil {
		return err
	}

	if count == 0 {
		return nil
	}

	for i := int32(0); i < count; i++ {
		tmap := newTextureMapping()
		if err := tmap.Read(r); err != nil {
			return err
		}
		*tm = append(*tm, tmap)
	}

	return nil
}
