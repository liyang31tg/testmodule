package testmodule

import "testing"

func TestHello(t *testing.T) {
	want := "V0 hello"
	if Hello() != want {
		t.Error("df")
	}

}
