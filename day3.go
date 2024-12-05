package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var sum = 0

func main() {
	var foundMulInstances []string

	fmt.Printf("Input: \n")

	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		foundMulInstances = append(foundMulInstances, re.FindAllString(line, -1)...)
	}

	fmt.Println("you entered: \n", foundMulInstances)

	for _, multiply := range foundMulInstances {

		cleaned := strings.TrimPrefix(strings.TrimSuffix(multiply, ")"), "mul(")

		numbers := strings.Split(cleaned, ",")

		x, err1 := strconv.Atoi(numbers[0])
		y, err2 := strconv.Atoi(numbers[1])

		if err1 != nil || err2 != nil {
			continue
		}

		mul(x, y)

	}

	fmt.Printf("sum: %v ", sum)
}

func mul(x int, y int) {
	sum += x * y
}
