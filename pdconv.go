package pdconv

import (
	"errors"
	"fmt"
)

const (
	lowNibble  = 0x0f
	highNibble = 0xf0
)

var errSyntax = errors.New("invalid syntax")

// Ptoi converts the packed decimal bytes to a int value.
func Ptoi(bs []byte) (int, error) {
	if len(bs) == 0 {
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

// Itop converts the int to packed decimal bytes.
func Itop(x int, nb int) ([]byte, error) {
	if len(fmt.Sprintf("%d", abs(x))) > nb*2-1 {
		return nil, errSyntax
	}
	pdFmt := "%0" + fmt.Sprintf("%dd0", nb*2-1)

	bs := []byte(fmt.Sprintf(pdFmt, abs(x)))
	pd := make([]byte, nb)

	for i, j := 0, 0; i < nb; i++ {
		pd[i] = (bs[j]-0x30)<<4 | (bs[j+1] - 0x30)
		j += 2
	}
	if x > 0 {
		pd[nb-1] |= 0xC
	} else {
		pd[nb-1] |= 0xD
	}

	return pd, nil
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}
