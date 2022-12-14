package unit_test_lab

import "testing"

func TestSum(t *testing.T) {
	a, b := 1, 2
	result := Sum(a, b)
	if result == 3 {
		t.Logf("expected:%d,actual=%d", 3, result)
	} else {
		t.Fatalf("expected:%d,actual=%d", 3, result)
	}

	t.Log("test done")
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}
