package gopy

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
)

func Pack(val int32) (string, error) {
	var out strings.Builder
	dBytes := make([]byte, 4)
	_, err := binary.Encode(dBytes, binary.LittleEndian, val)
	if err != nil {
		return "", errors.New("could not encode the value : " + err.Error())
	}
	for _, byte := range dBytes {
		out.WriteString(fmt.Sprintf("\\x%.2x", byte))
	}
	return out.String(), nil
}

func PackLong(val int64) (string, error) {
	var out strings.Builder
	dBytes := make([]byte, 8)
	_, err := binary.Encode(dBytes, binary.LittleEndian, val)
	if err != nil {
		return "", errors.New("could not encode the value : " + err.Error())
	}
	for _, byte := range dBytes {
		out.WriteString(fmt.Sprintf("\\x%.2x", byte))
	}
	return out.String(), nil
}
