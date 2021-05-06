package helper

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

// Readn will read exactly n bytes, It will return err when Read return err or there is not enough n bytes.
func Readn(r io.Reader, n int) ([]byte, error) {
	data := make([]byte, n)
	v, err := r.Read(data)
	if err != nil {
		return nil, err
	}

	if v != n {
		return nil, errors.Errorf("Readn failed, expect %v bytes, actual %v bytes", n, v)
	}
	return data, nil
}

// ReadnWithPanic will return exactly n bytes, or panic
func ReadnWithPanic(r io.Reader, n int) []byte {
	data, err := Readn(r, n)
	if err != nil {
		panic(err)
	}

	return data
}

// Read1 return 1 bytes
func Read1(r io.Reader) (byte, error) {
	data, err := Readn(r, 1)
	if err != nil {
		return 0, err
	}

	return data[0], nil
}

// ReadObject will read data and unmarshal into v object, it use LittleEndian ByteOrder
func ReadObject(r io.Reader, v interface{}) error {
	return binary.Read(r, binary.LittleEndian, v)
}
