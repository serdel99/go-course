package main

import (
	"testing"
)

func TestSum (t *testing.T){
	total := Sum(5,5)
	if total != 10{
		t.Errorf("Sum was incorrect,  got %d expected %d", total, 10)
	}
	tables := [] struct {
		a int 
		b int 
		n int 
	}{
		{1,2,3},
		{2,2,4},
		{25,26, 51},
	}
	for _,item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n{
			t.Errorf("Sum was incorrect,  got %d expected %d", total, item.n)
		}
	}
	
}


func TestMax (t *testing.T){
	tables := []struct{
		a int 
		b int 
		c int 
	}{
		{4,2,4},
		{3,2,3},
	}

	for _, item  := range tables{
		max := GetMax(item.a, item.b)
		if max  != item.c {
			t.Errorf("GetMax was incorrect,  got %d, expected %d", max, item.c)
		}
	}
}

func TestFib(t  *testing.T){
	tables := []struct {
		a int
		n int 
	}{
		{1,1},
		{ 8, 21},
		{50,12586269025 },
	}
	for _, item := range tables{
		fib := Fibonnaci(item.a)
		if fib != item.n {
			t.Errorf("Fibonnacci was incorrect, got %d expected %d", fib, item.n)
		}
	}
}

