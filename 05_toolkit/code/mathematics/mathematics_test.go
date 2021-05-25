package mathematics

import "testing"

func TestAdd(t *testing.T){
	total := Add(1,2,3,4,5)
	want := 15

	if want != total {
		t.Errorf("Expected %d, received %d", want, total)
	}
}