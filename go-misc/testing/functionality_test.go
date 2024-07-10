package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1,2)
	t.Logf("Case A: 1, B: 2, expected: 3, result: %d",result)

	if result != 3 {
		t.Errorf("result is wrong")
	}
} 

func TestDiff(t *testing.T) {
	result := Diff(1,2)
	t.Logf("Case A: 1, B: 2, expected: 1, result: %d",result)

	if result != 1 {
		t.Errorf("result is wrong")
	}
} 

func BenchmarkAdd(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		Add(i, i+1)
	}
}