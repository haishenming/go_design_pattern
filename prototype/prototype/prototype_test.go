package prototype

import (
	"testing"
)

func TestPrototype_Clone(t *testing.T) {
	p := &Prototype{}
	p1 := p.Clone()

	// p1由p克隆出来，指针不应相等
	if p == p1 {
		t.Fatal("error")
	}
}
