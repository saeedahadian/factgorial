package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/saeedahadian/flags"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number>")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Error parsing your number %s.", os.Args[1])
		os.Exit(1)
	}
	fs, err := flags.ParseFlags(os.Args[2:])
	if err != nil {
		fmt.Printf("Error parsing flags: %v", os.Args[2:])
	}
	if hasAllFlag(fs) {
		for i := 1; i <= num; i++ {
			result, err := factorial(i)
			if err != nil {
				fmt.Printf("Error calculating factorial(%d): %s\n", i, err.Error())
				os.Exit(1)
			}
			fmt.Printf("factorial(%d) = %d\n", i, result)
		}
		return
	}

	result, err := factorial(num)
	if err != nil {
		fmt.Printf("Error calculating factorial(%d): %s\n", num, err.Error())
		os.Exit(1)
	}
	fmt.Printf("factorial(%d) = %d\n", num, result)
}

func hasAllFlag(fs []*flags.Flag) bool {
	for _, flag := range fs {
		if flag.Key == "all" && flag.Value.String() == "true" {
			return true
		}
	}
	return false
}

func factorial(num int) (uint64, error) {
	if num < 0 {
		return 0, errors.New("it's mathematically impossible to take the factorial of negative numbers")
	}

	if num < 2 {
		return 1, nil
	}

	var result uint64 = 1
	for i := 2; i <= num; i++ {
		result *= uint64(i)
	}
	return result, nil
}
