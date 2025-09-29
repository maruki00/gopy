package gopy

import (
	"errors"
	"fmt"
)

// func UnpackLong(val string) (int64, error) {
// 	return 0, nil
// }

func Unpack(val string) (int64, error) {
	if len(val)%4 != 0 {
		return 0, errors.New("invalid data")
	}
	for i, j := range []byte(val) {
		fmt.Println(i, j)
	}

	return 0, nil
}
