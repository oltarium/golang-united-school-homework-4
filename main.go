package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res, err := StringSum("5  + 1")
	println(res)
	if err != nil {
		println(err.Error())
	}

}

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// +3 + 5
func StringSum(input string) (output string, err error) {
	if len(strings.Trim(input, " ")) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	var nums []string
	var currentNumber string
	var sign string

	for i := 0; i < len(input); i++ {
		char := string(input[i])
		if char == " " {
			continue
		}

		if char == "-" || char == "+" {
			if len(currentNumber) > 0 {
				nums = append(nums, sign+currentNumber)
			}
			sign = char
			currentNumber = ""
			continue
		}
		currentNumber += char
	}
	if len(currentNumber) > 0 {
		nums = append(nums, sign+currentNumber)
	}
	if len(nums) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	firstNumber, err := strconv.Atoi(nums[0])
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	secondNumber, err := strconv.Atoi(nums[1])
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	result := firstNumber + secondNumber
	return strconv.Itoa(result), nil
}
