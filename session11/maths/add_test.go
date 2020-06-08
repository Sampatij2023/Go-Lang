package maths

import "testing"

//T of Test capital letterer
func TestAdd(t *testing.T) {
	actual := add(20, 10)
	// expected:=
	var expected int64 = 30

	if actual != expected {
		t.Error("Actual is :", actual, "but expected is :", expected)
	}
}

func TestAddBulk(t *testing.T){
	var list =[]struct{
		first int 
		second int 
		expected int64

	}{
		{
			20,10,30
		},
		{
			-20,10,-10
		},
		{
			25,-15,10
		},
		{
			-120,-10,-130
		},
	}
	for _,str:= range list{
		if actual:= add(str.first,str.second); actual !=str.expected{
			t.Error("Actual is :", actual, "but expected is :", expected)
		}
	}
}