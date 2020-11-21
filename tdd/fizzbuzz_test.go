package tdd

import "testing"

func TestShouldReturn1WhenInput1(t *testing.T){
 r := fizzbuzz(1)

 if r != "1" {
 	t.Errorf("exected %s but got %s", "1", r)
 }
}
