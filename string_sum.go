package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf
// [-, 300, +, 5]
func StringSum(input string) (output string, err error) {
	const numbers = "0123456789"
	var nums []string
	var currentNumber string
	for i := 0; i < len(input); i++ {
		char := string(input[i])
		if char != " " && char != "-" && char != "+" && !strings.ContainsAny(numbers, char) {
			return "", fmt.Errorf("%s", "Invalid characters")
		}
		if (char == "-" || char == "+") && len(currentNumber) >= 1 {
			nums = append(nums, currentNumber)
			currentNumber = char
			continue
		}
		if char == "-" || char == "=" {
			currentNumber = char
			continue
		}
		if strings.ContainsAny(numbers, char) {
			currentNumber += char
		}
	}
	if len(nums) == 0 {
		return "", fmt.Errorf("%s", errorEmptyInput)
	}
	nums = append(nums, currentNumber)
	if len(nums) != 2 {
		return "", fmt.Errorf("%s", errorNotTwoOperands)
	}
	firstNumber, err := strconv.Atoi(nums[0])
	if err != nil {
		return "", fmt.Errorf("%s", "Invalid number")
	}
	secondNumber, err := strconv.Atoi(nums[1])
	if err != nil {
		return "", fmt.Errorf("%s", "Invalid number")
	}
	result := firstNumber + secondNumber
	output = strconv.Itoa(result)
	return
}
