package kn5

import "io"

type MaterialBlendMode byte

func (mbm *MaterialBlendMode) Read(r io.Reader) error {
	buf := make([]byte, 1)
	if _, err := r.Read(buf); err != nil {
		return err
	}
	*mbm = MaterialBlendMode(buf[0])
	return nil
}

func (bm MaterialBlendMode) String() string {
	switch bm {
	case Opaque:
		return "Opaque"
	case AlphaBlend:
		return "Alpha blend"
	case AlphaToCoverage:
		return "Alpha to coverage"
	}
	return "INVALID"
}

const (
	Opaque MaterialBlendMode = iota
	AlphaBlend
	AlphaToCoverage
)
