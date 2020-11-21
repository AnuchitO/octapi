package tdd

import "testing"

func TestShouldReturn1WhenInput1(t *testing.T){
 r := fizzbuzz(1)

 if r != "1" {
 	t.Errorf("exected %s but got %s", "1", r)
 }
}

func TestShouldReturn2WhenInput2(t *testing.T){
	r := fizzbuzz(2)

	if r != "2" {
		t.Errorf("exected %s but got %s", "2", r)
	}
}

func TestShouldReturnFizzWhenInput3(t *testing.T){
	r := fizzbuzz(3)

	if r != "Fizz" {
		t.Errorf("exected %s but got %s", "Fizz", r)
	}
}

func TestShouldReturn4WhenInput4(t *testing.T){
	r := fizzbuzz(4)

	if r != "4" {
		t.Errorf("exected %s but got %s", "4", r)
	}
}
