package gopy

import (
	"encoding/binary"
	"fmt"
	"strings"
)

func UnpackLong(val int64) (int64, error) {
	var out strings.Builder
	dBytes := make([]byte, 8)
	binary.Encode(dBytes, binary.LittleEndian, val)
	for _, byte := range dBytes {
		out.WriteString(fmt.Sprintf("\\x%.2x", byte))
	}
	fmt.Println(out.String())
}

func Unpack(val int32) (int64, error) {
	var out strings.Builder
	dBytes := make([]byte, 4)
	binary.Encode(dBytes, binary.LittleEndian, val)
	for _, byte := range dBytes {
		out.WriteString(fmt.Sprintf("\\x%.2x", byte))
	}
	fmt.Println(out.String())
	return 0, nil
}
