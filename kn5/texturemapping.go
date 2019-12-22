package kn5

import (
	"fmt"
	"io"
)

type TextureMapping struct {
	Name    string
	Texture string
	Slot    int32
}

func newTextureMapping() *TextureMapping { return &TextureMapping{} }

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

func (tm *TextureMapping) Write(w io.Writer) error {
	// Name.
	if err := WriteString(w, tm.Name); err != nil {
		return err
	}
	// Slot.
	if err := WriteInt32(w, tm.Slot); err != nil {
		return err
	}
	// Texture.
	if err := WriteString(w, tm.Texture); err != nil {
		return err
	}
	return nil
}

type TextureMappings []*TextureMapping

func (tm *TextureMappings) Read(r io.Reader) error {
	// Count.
	count, err := ReadInt32(r)
	if err != nil {
		return err
	}
	// No data.
	if count == 0 {
		return nil
	}
	// Textures.
	for i := int32(0); i < count; i++ {
		tmap := newTextureMapping()
		if err := tmap.Read(r); err != nil {
			return err
		}
		*tm = append(*tm, tmap)
	}
	return nil
}

func (tm *TextureMappings) Write(w io.Writer) error {
	// Count.
	if err := WriteInt32(w, int32(len(*tm))); err != nil {
		return err
	}
	// Textures.
	for _, tmap := range *tm {
		if err := tmap.Write(w); err != nil {
			return err
		}
	}
	return nil
}
