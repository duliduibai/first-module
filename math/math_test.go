package math

import "testing"

func TestAdd(t *testing.T) {
	total := Add([]float32{1,2,3,4.2})
	if total != 10.2 {
		t.Fatal("result err")
	}
}