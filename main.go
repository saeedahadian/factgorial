package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <number>")
		return
	}

	inputScanner := bufio.NewScanner(os.Stdin)
	for inputScanner.Scan() {
		input := inputScanner.Text()

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(factorial(num))
	}

	if err := inputScanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func factorial(num int) (int, error) {
	if num < 0 {
		return -1, errors.New("it's mathematically impossible to take the factorial of negative numbers")
	}

	if num < 2 {
		return 1, nil
	}

	result := 1
	for i := 2; i <= num; i++ {
		result *= i
	}
	return result, nil
}
