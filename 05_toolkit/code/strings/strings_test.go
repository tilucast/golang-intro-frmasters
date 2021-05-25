package strings

import "testing"

func TestMakeExcited(t *testing.T){
	expected := "I AM EXCITED"
	actual := MakeExcited("i am excited")

	if actual != expected {
		t.Errorf("Expecting: '%s', but found '%s'.", expected, actual)
	}
}