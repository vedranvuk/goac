package kn5

import (
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
	// Version.
	version, err := ReadInt32(r)
	if err != nil {
		return err
	}
	h.Version = version
	// No extra.
	if h.Version < 6 {
		return nil
	}
	// Extra.
	extra, err := ReadInt32(r)
	if err != nil {
		return err
	}
	h.Extra = extra
	return nil
}

func (h *Header) Write(w io.Writer) error {
	// Version.
	if err := WriteInt32(w, h.Version); err != nil {
		return err
	}
	// No extra.
	if h.Version < 6 {
		return nil
	}
	// Extra.
	if err := WriteInt32(w, h.Extra); err != nil {
		return err
	}
	return nil
}

func newHeader() *Header { return &Header{} }
