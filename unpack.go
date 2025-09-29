package gopy

import (
	"encoding/binary"
	"fmt"
)

func UnpackLong(_type string, val int64) (int64, error) {
	d := make([]byte, 8)
	binary.Encode(d, binary.LittleEndian, val)
	fmt.Println(d)
	return 0, nil
}

func Unpack(val int32) (int64, error) {
	d := make([]byte, 4)
	binary.Encode(d, binary.LittleEndian, val)
	fmt.Println(d)
	return 0, nil
}
