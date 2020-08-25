package stay_minimal

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// max finds maximum value from int array
func max(values []int) (int, error) {
	if len(values) == 0 {
		return 0, errors.New("cannot find max value in empty array")
	}
	var maxFound int
	for i, e := range values {
		if i == 0 || e > maxFound {
			maxFound = e
		}
	}
	return maxFound, nil
}

func TestFindingMaximumValueWhenInTheMiddleOfInputArray(t *testing.T) {
	// given
	elements := []int{1, 3, 2}

	// when
	maxFound, err := max(elements)

	// then
	assert.Nil(t, err)
	assert.Equal(t, 3, maxFound)
}
