package pdconv

import "errors"

const (
	lowNibble  = 0x0f
	highNibble = 0xf0
)

var errSyntax = errors.New("invalid syntax")

// Ptoi converts the packed decimal bytes to a int value.
func Ptoi(bs []byte) (int, error) {
    if len(bs) == 0  {
        return 0, errSyntax
    }

	n := 0
	for i, b := range bs {
		high := (int(b) & highNibble) >> 4
		if isDigit(high) != true {
			return 0, errSyntax
		}
		n = n*10 + high

		low := int(b) & lowNibble
		if i == len(bs)-1 {
			if isSign(low) {
				return n * applySign(low), nil
			}
			return 0, errSyntax
		}
		if isDigit(low) != true {
			return 0, errSyntax
		}
		n = n*10 + low
	}

	return 0, errSyntax
}

func isDigit(i int) bool {
	if i >= 0x00 && i <= 0x09 {
		return true
	}
	return false
}

func isSign(i int) bool {
	if i == 0x0C || i == 0x0D {
		return true
	}
	return false
}

func applySign(i int) int {
	if i == 0x0C {
		return 1
	}
	return -1
}
