package kn5

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type File struct {
	Header    *Header
	Textures  *Textures
	Materials *Materials
	RootNode  *Node
}

func newFile() *File {
	f := &File{
		Header:    newHeader(),
		Textures:  newTextures(),
		Materials: newMaterials(),
		RootNode:  newNode(),
	}
	return f
}

func (f *File) String() string {
	return fmt.Sprintf(`File:
Header:    
%v

Textures:  
%v

Materials: 
%v

Nodes:
%v
`,
		f.Header, f.Textures, f.Materials, f.RootNode)
}

func (f *File) Read(r io.Reader) error {
	// Magic.
	magic := make([]byte, 6)
	if _, err := r.Read(magic); err != nil {
		return err
	}
	if string(magic) != "sc6969" {
		return errors.New("Not a valid KN5 file.")
	}
	// Header.
	if err := f.Header.Read(r); err != nil {
		return err
	}
	// Textures.
	if err := f.Textures.Read(r); err != nil {
		return err
	}
	// Materials.
	if err := f.Materials.Read(r); err != nil {
		return err
	}
	// Nodes.
	if err := f.RootNode.Read(r); err != nil {
		return err
	}
	return nil
}

func (f *File) Write(w io.Writer) error {
	// Magic.
	if _, err := w.Write([]byte("sc6969")); err != nil {
		return err
	}
	// Header.
	if err := f.Header.Write(w); err != nil {
		return err
	}
	// Textures.
	if err := f.Textures.Write(w); err != nil {
		return err
	}
	// Materials.
	if err := f.Materials.Write(w); err != nil {
		return err
	}
	// Nodes.
	if err := f.RootNode.Write(w); err != nil {
		return err
	}
	return nil
}

func Load(filename string) (*File, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := bufio.NewReader(file)

	f := newFile()
	if err := f.Read(r); err != nil {
		return nil, fmt.Errorf("error reading kn5 file: %w", err)
	}

	return f, nil
}

func Save(filename string, f *File) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return f.Write(file)
}
