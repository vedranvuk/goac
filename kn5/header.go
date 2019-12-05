package kn5

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Header struct {
	Version int32
	Extra   int32
}

func (h *Header) String() string {
	return fmt.Sprintf(`Header:
	Version: %d
	Extra:   %d`,
		h.Version, h.Extra)
}

func (h *Header) Read(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if h.Version < 6 {
		return nil
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Extra); err != nil {
		return err
	}
	return nil
}

func newHeader() *Header {
	return &Header{}
}
