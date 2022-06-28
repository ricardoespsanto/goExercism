package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
	message string
}

func (e SillyNephewError) Error() string {
	return e.message
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()

	switch {
	case errors.Is(err, ErrScaleMalfunction):
		if amount < 0 {
			return 0, errors.New("negative fodder")
		}

		amount *= 2
	case err != nil:
		return 0, err
	case cows == 0:
		return 0, errors.New("division by zero")
	case cows < 0:
		return 0, SillyNephewError{message: fmt.Sprintf("silly nephew, there cannot be %d cows", cows)}
	case amount < 0 && err == nil:
		return 0, errors.New("negative fodder")
	}

	return amount / float64(cows), nil
}
