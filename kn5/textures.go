package kn5

import (
	"fmt"
	"io"
)

type Texture struct {
	Name   string
	Active bool
	Length int32
	Data   []byte
}

func newTexture() *Texture { return &Texture{} }

func (t *Texture) String() string {
	return fmt.Sprintf(`Texture:
Name:   %s
Active: %t
Length: %d
Data:   %d

`, t.Name, t.Active, t.Length, len(t.Data))
}

func (t *Texture) Read(r io.Reader) error {
	// Active.
	active, err := ReadInt32(r)
	if err != nil {
		return err
	}
	t.Active = active > 0
	// Name.
	name, err := ReadString(r)
	if err != nil {
		return err
	}
	t.Name = name
	// Length.
	length, err := ReadInt32(r)
	if err != nil {
		return err
	}
	t.Length = length
	// No data.
	if length == 0 {
		return nil
	}
	// Data.
	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	t.Data = buf
	return nil
}

func (t *Texture) Write(w io.Writer) error {
	// Active.
	b := int32(0)
	if t.Active {
		b = 1
	}
	if err := WriteInt32(w, b); err != nil {
		return err
	}
	// Name.
	if err := WriteString(w, t.Name); err != nil {
		return err
	}
	// Length.
	if err := WriteInt32(w, t.Length); err != nil {
		return err
	}
	// Data.
	if _, err := w.Write(t.Data); err != nil {
		return err
	}
	return nil
}

type Textures map[string]*Texture

func (t *Textures) Read(r io.Reader) error {
	// Texture count.
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
		tex := newTexture()
		if err := tex.Read(r); err != nil {
			return err
		}
		map[string]*Texture(*t)[tex.Name] = tex
	}
	return nil
}

func (t *Textures) Write(w io.Writer) error {
	// Texture count.
	if err := WriteInt32(w, int32(len(*t))); err != nil {
		return err
	}
	// Textures.
	for _, tex := range *t {
		if err := tex.Write(w); err != nil {
			return err
		}
	}
	return nil
}

func newTextures() *Textures {
	t := Textures(make(map[string]*Texture))
	return &t
}
