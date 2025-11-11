package mathutil

import (
	"errors"

)

// Average returns the average of a slice of numbers.
// Returns an error if the slice is empty.
func Average(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("cannot average empty slice")
	}
	
	var sum float64
	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums)), nil
}