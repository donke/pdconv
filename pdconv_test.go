package pdconv

import (
	"reflect"
	"testing"
)

func TestPtoi(t *testing.T) {
	bs := [][]byte{{0x01, 0x2C}, {0x12, 0x3D}, {0x00, 0x00, 0x12, 0x3C}}
	es := []int{12, -123, 123}
	for i, b := range bs {
		actual, _ := Ptoi(b)
		if es[i] != actual {
			t.Errorf("expected: %v actual: %v", es[i], actual)
		}
	}
}

func TestItoP(t *testing.T) {
	is := [][]int{{12, 2}, {-123, 2}, {123, 4}}
	es := [][]byte{{0x01, 0x2C}, {0x12, 0x3D}, {0x00, 0x00, 0x12, 0x3C}}
	for i, v := range is {
		actual, _ := Itop(v[0], v[1])
		if !reflect.DeepEqual(es[i], actual) {
			t.Errorf("expected: %v actual: %v", es[i], actual)
		}
	}
}

func TestPtoiWithInvalidValue(t *testing.T) {
	bs := [][]byte{{}, {0xA1}, {0x01, 0x2A}, {0x01}}
	for _, b := range bs {
		actual, _ := Ptoi(b)
		if 0 != actual {
			t.Errorf("expected: 0 actual: %v", actual)
		}
	}
}
