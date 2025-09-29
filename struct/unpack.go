package gopy_struct

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
)

func Unpack(val string) (uint32, error) {
	if len(val)%4 != 0 {
		return 0, errors.New("invalid data")
	}
	_bytesVal := []byte(val)

	var v uint32
	binary.Decode(_bytesVal, binary.LittleEndian, &v)
	fmt.Println(v, _bytesVal)
	return v, nil
}

func Unpack2Hex(val string) (string, error) {
	if len(val)%4 != 0 {
		return "", errors.New("invalid data")
	}
	_bytesVal := []byte(val)
	return hex.EncodeToString(_bytesVal), nil
}

func UnpackLong(val string) (uint64, error) {
	if len(val)%4 != 0 {
		return 0, errors.New("invalid data")
	}
	_bytesVal := []byte(val)

	var v uint64
	binary.Decode(_bytesVal, binary.LittleEndian, &v)
	fmt.Println(v, _bytesVal)
	return v, nil
}
