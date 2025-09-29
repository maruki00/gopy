package gopy_struct

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
)

func Unpack(val string) (uint32, error) {
	// pattern \xAB -> 4 symboles
	if len(val)%4 != 0 {
		return 0, errors.New("invalid data")
	}
	_bytesVal := []byte(val)

	var v uint32
	binary.Decode(_bytesVal, binary.LittleEndian, &v)
	fmt.Println(v, _bytesVal)
	return v, nil
}

func UnpackLong(val string) (uint64, error) {
	// pattern \xAB -> 4 symboles
	if len(val)%4 != 0 {
		return 0, errors.New("invalid data")
	}
	_bytesVal := []byte(val)

	var v uint64
	binary.Decode(_bytesVal, binary.LittleEndian, &v)
	fmt.Println(v, _bytesVal)
	return v, nil
}

func Unpack2Hex(val string) (string, error) {
	// pattern \xAB -> 4 symboles
	if len(val)%4 != 0 {
		return "", errors.New("invalid data")
	}
	_bytesVal := []byte(val)
	dd := make([]byte, 8)
	binary.Decode(_bytesVal, binary.BigEndian, dd)

	return "0x" + hex.EncodeToString(dd), nil
}
