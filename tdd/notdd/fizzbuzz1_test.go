package main

import "testing"

func TestFizzBuzz1(t *testing.T) {
   r := fizzbuzz(1)

   if r != "1" {
   	t.Error("expected 1 but got", r)
   }
}
