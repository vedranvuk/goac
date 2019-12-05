package kn5

import (
	"encoding/binary"
	"io"
)

func ReadString(r io.Reader) (string, error) {
	length := int32(0)
	if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
		return "", err
	}
	buf := make([]byte, length)
	if _, err := r.Read(buf); err != nil {
		return "", err
	}
	return string(buf), nil
}

func ReadInt32(r io.Reader) (int32, error) {
	val := int32(0)
	if err := binary.Read(r, binary.LittleEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func ReadUint16(r io.Reader) (uint16, error) {
	val := uint16(0)
	if err := binary.Read(r, binary.LittleEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func ReadUint32(r io.Reader) (uint32, error) {
	val := uint32(0)
	if err := binary.Read(r, binary.LittleEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func ReadBoolean(r io.Reader) (bool, error) {
	buf := make([]byte, 1)
	if _, err := r.Read(buf); err != nil {
		return false, err
	}
	return buf[0] > 1, nil
}

func ReadFloat(r io.Reader) (float32, error) {
	val := float32(0)
	if err := binary.Read(r, binary.LittleEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func ReadFloats(r io.Reader, count int) ([]float32, error) {
	floats := make([]float32, count)
	for i := 0; i < count; i++ {
		float, err := ReadFloat(r)
		if err != nil {
			return nil, err
		}
		floats[i] = float
	}
	return floats, nil
}

func ReadMatrix(r io.Reader) ([]float32, error) {
	var floats []float32
	for i := 0; i < 4; i++ {
		vals, err := ReadFloats(r, 4)
		if err != nil {
			return nil, err
		}
		floats = append(floats, vals...)
	}
	return floats, nil
}
