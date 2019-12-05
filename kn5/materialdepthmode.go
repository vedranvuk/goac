package kn5

import "io"

type MaterialDepthMode int32

const (
	DepthNormal MaterialDepthMode = iota
	DepthNoWrite
	DepthOff
)

func (mdm *MaterialDepthMode) Read(r io.Reader) error {
	val, err := ReadInt32(r)
	if err != nil {
		return err
	}
	*mdm = MaterialDepthMode(val)

	return nil
}

func (dm MaterialDepthMode) String() string {
	switch dm {
	case DepthNormal:
		return "Normal"
	case DepthNoWrite:
		return "Read-only"
	case DepthOff:
		return "Off"
	}
	return "INVALID"
}
