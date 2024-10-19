package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSupermanRescure(t *testing.T) {
	t.Run("Roof's length is 5 and the chicken's amout is 5", func(t *testing.T) {
		maximum := MaximumRescue(5, []int{2, 5, 10, 12, 15})
		assert.Equal(t, 2, maximum)
	})
	t.Run("Roof's length is 10 and the chicken's amout is 6", func(t *testing.T) {
		maximum := MaximumRescue(10, []int{1, 11, 30, 34, 35, 37})
		assert.Equal(t, 4, maximum)
	})
	t.Run("Roof's length is 5 and the chicken's amout is 10 with maximum in left side", func(t *testing.T) {
		maximum := MaximumRescue(5, []int{1, 2, 3, 4, 5, 10, 15, 18, 20, 22})
		assert.Equal(t, 5, maximum)
	})
	t.Run("Roof's length is 5 and the chicken's amout is 10 with maximum in right side", func(t *testing.T) {
		maximum := MaximumRescue(5, []int{1, 3, 5, 7, 9, 10, 21, 22, 23, 24, 25})
		assert.Equal(t, 5, maximum)
	})
	t.Run("Roof's length is 5 and the chicken's amout is 10 with maximum in the middle", func(t *testing.T) {
		maximum := MaximumRescue(10, []int{2, 10, 21, 27, 30, 50, 51, 52, 53, 54, 55, 56, 57, 70, 74, 81, 88})
		assert.Equal(t, 8, maximum)
	})
}
