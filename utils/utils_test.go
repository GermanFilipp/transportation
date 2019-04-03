package utils

import "testing"

func TestTotalSum(t *testing.T) {
	expected := 10
	actual := TotalSum([]int{5, 5})
	if actual != expected {
		t.Errorf("TestTotalSum test expected [%+v],\n actual [%+v]", expected, actual)
	}
}

func TestCopyArray(t *testing.T) {
	input := []int{2, 3, 4, 5}
	newArray := CopyArray(input)
	if &input == &newArray || &input[0] == &newArray[0] {
		t.Errorf("TestCopyArray test expected new array, actual - point for old")
	}
}
