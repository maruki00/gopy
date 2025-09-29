package gopy

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// func UnpackLong(val string) (int64, error) {
// 	return 0, nil
// }

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
