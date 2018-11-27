package moda

import "testing"

func Test_Echo(t *testing.T) {
	if Echo("test")!="test" {
		t.Fail()
	}
}
