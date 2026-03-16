package thefarm

import (
    "fmt";
    "errors"
)

type InvalidCowsError struct {
    cowsNum int
    errMsg string
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.cowsNum, e.errMsg)
}

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cowsNum int) (float64, error) {
    totalAmount, err := fc.FodderAmount(cowsNum)
    if err != nil {
        return 0, err
    }

    ff, err := fc.FatteningFactor()

    if err != nil {
        return 0, err
    }
    
    return (totalAmount * ff) / float64(cowsNum), err
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cowsNum int) (float64, error) {
    errMsg := errors.New("invalid number of cows")

    if cowsNum <= 0 {
        return 0, errMsg
    }

    return DivideFood(fc, cowsNum)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cowsNum int) error {
    if cowsNum < 0 {
        return &InvalidCowsError{
            cowsNum: cowsNum,
            errMsg: "there are no negative cows",
        }
    } else if cowsNum == 0 {
        return &InvalidCowsError{
            cowsNum: 0,
            errMsg: "no cows don't need food",
        }
    } 

    return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
