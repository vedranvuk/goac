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
	if _, err := io.ReadFull(r, buf); err != nil {
		return "", err
	}
	return string(buf), nil
}

func WriteString(w io.Writer, s string) error {
	if err := WriteInt32(w, int32(len(s))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(s)); err != nil {
		return err
	}
	return nil
}

func ReadInt32(r io.Reader) (int32, error) {
	val := int32(0)
	if err := binary.Read(r, binary.LittleEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func WriteInt32(w io.Writer, i int32) error {
	return binary.Write(w, binary.LittleEndian, i)
}

func ReadUint16(r io.Reader) (uint16, error) {
	val := uint16(0)
	if err := binary.Read(r, binary.LittleEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func WriteUint16(w io.Writer, i uint16) error {
	return binary.Write(w, binary.LittleEndian, i)
}

func ReadUint32(r io.Reader) (uint32, error) {
	val := uint32(0)
	if err := binary.Read(r, binary.LittleEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func WriteUint32(w io.Writer, i uint32) error {
	return binary.Write(w, binary.LittleEndian, i)
}

func ReadBoolean(r io.Reader) (bool, error) {
	buf := make([]byte, 1)
	if _, err := io.ReadFull(r, buf); err != nil {
		return false, err
	}
	return buf[0] > 1, nil
}

func WriteBoolean(w io.Writer, b bool) error {
	buf := []byte{0}
	if b {
		buf[0] = 1
	}
	if _, err := w.Write(buf); err != nil {
		return err
	}
	return nil
}

func ReadFloat(r io.Reader) (float32, error) {
	val := float32(0)
	if err := binary.Read(r, binary.LittleEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func WriteFloat(w io.Writer, f float32) error {
	return binary.Write(w, binary.LittleEndian, f)
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

func WriteFloats(w io.Writer, floats []float32) error {
	for i := 0; i < len(floats); i++ {
		if err := WriteFloat(w, floats[i]); err != nil {
			return err
		}
	}
	return nil
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

func WriteMatrix(w io.Writer, matrix []float32) error {
	for i := 0; i < 4; i++ {
		if err := WriteFloats(w, matrix[i*4:i*4+4]); err != nil {
			return err
		}
	}
	return nil
}
