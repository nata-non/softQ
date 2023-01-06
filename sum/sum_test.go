package sum_test

import "testing"

func TestSum(t *testing.T) {
	t.Run("Is it 4?", func(t *testing.T) {
		if sum(2, 2) == 4 {
			t.Log("Pass")
		} else {
			t.Errorf("Not Pass")
		}
	})
	t.Run("Is it 5?", func(t *testing.T) {
		if sum(1, 3) == 5 {
			t.Log("Pass")
		} else {
			t.Errorf("Not Pass")
		}
	})
	t.Run("Is it 6?", func(t *testing.T) {
		if sum(1, 5) == 6 {
			t.Log("Pass")
		} else {
			t.Errorf("Not Pass")
		}
	})
}
func sum(a, b int) int {
	return a + b
}
