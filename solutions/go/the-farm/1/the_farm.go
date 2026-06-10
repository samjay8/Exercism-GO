package thefarm

import (
	"fmt"
)


// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, numofcows int) (float64, error) {
	if numofcows <= 0 {
		return 0, fmt.Errorf("something went wrong")
	}
	fodder, err := f.FodderAmount(numofcows)
	if err != nil {
		return 0, err
	}
	factor, err := f.FatteningFactor()
	if err != nil {
		return 0, err
	}
	result := fodder / float64(numofcows) * factor
	return result, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(f FodderCalculator, numofcows int) (float64, error) {

	if numofcows > 0 {
		return DivideFood(f, numofcows)

	} else if numofcows <= 0 {
		return 0, fmt.Errorf("invalid number of cows")
	}
	return 0, nil
}

type InvalidCowsError struct {
	cows    int
	message string
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numofcows int) error {
	if numofcows < 0 {
		return fmt.Errorf("%d cows are invalid: there are no negative cows", numofcows)
	}
	if numofcows == 0 {
		return fmt.Errorf("%d cows are invalid: no cows don't need food", numofcows)
	}
	return nil
}
