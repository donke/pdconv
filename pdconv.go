package pdconv

import (
	"errors"
	"fmt"
)

const (
	highMask = 0xf0
	lowMask  = 0x0f
)

var errSyntax = errors.New("invalid syntax")

// Ptoi interprets a bs the packed decimal bytes in the base 10 and returns the value.
func Ptoi(bs []byte) (int, error) {
	if len(bs) == 0 {
		return 0, errSyntax
	}

	var hi, lo, n int

	for i := 0; i < len(bs); i++ {
		hi = (int(bs[i]) & highMask) >> 4
		if !isDigit(hi) {
			return 0, errSyntax
		}
		n = n*10 + hi

		lo = int(bs[i]) & lowMask
		if !isDigit(lo) && !isSign(lo) {
			return 0, errSyntax
		}
		if isDigit(lo) {
			n = n*10 + lo
		}
	}
	if !isSign(lo) {
		return 0, errSyntax
	}
	return n * applySign(lo), nil
}

// Itop returns the packed decimal bytes of x in the given base 10.
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
	return i >= 0x00 && i <= 0x09
}

func isSign(i int) bool {
	return i == 0x0C || i == 0x0D
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
