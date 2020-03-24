package builder

import "testing"

func TestBuild(t *testing.T) {
	b := &ConcreteBuilder{}
	d := NewDirector(b)
	d.Construct()
	if b.Result() != "ABC" {
		t.Fatalf("Builder1 fail expect ABC acture %s", b.Result())
	}

}
