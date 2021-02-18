package code

import (
	"fmt"
	"strconv"
	"strings"
)

// The parameter of "func utama" could be inputted on "start.go"
func Utama(x string) int {
	inputNum := strings.Split(x, " ") // The string input value will be convert to Array
	arrayInput := make([]int, len(inputNum))

	for i := range arrayInput {
		arrayInput[i], _ = strconv.Atoi(inputNum[i]) // The conversion from String to Integer will be done here
	}
	count := 0

	for i := arrayInput[0]; i <= arrayInput[1]; i++ {
		startingPoint := i
		result := reverse(i) // This line refers to "func reverse" on bellow

		if startingPoint == result {
			count += 1 // Everytime startingPoint value matched with the Result, it will increase the Count value by 1
		}
	}
	return count // The result of Counted Palindrome
}

// The reverse version of Parameter value will be done in this function
func reverse(n int) int { 
	reverseInt := strconv.Itoa(n) // The conversion from Integer to String will be done here
	endPoint := ""

	for x := len(reverseInt); x > 0; x-- {
		endPoint += string(reverseInt[x-1])
	}

	newInt, err := strconv.Atoi(endPoint) // The conversion from String to Integer will be done here
	if err != nil {
		fmt.Println(err)
	}
	return newInt
}
