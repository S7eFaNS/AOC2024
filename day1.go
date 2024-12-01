package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var left_list, right_list []int
	total_distance := 0
	total_similarity := 0

	fmt.Printf("Input: \n")

	// read terminal input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Printf("Invalid format: %s\n", line)
		}

		left, err1 := strconv.Atoi(parts[0])
		right, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Printf("Error: %s\n", line)
		}

		left_list = append(left_list, left)
		right_list = append(right_list, right)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	// sort in ascending
	sort.Ints(left_list)
	sort.Ints(right_list)

	// calculate total distance between
	// 1st left element and 1st right element and the sum of all of it
	for i := 0; i < len(left_list); i++ {
		distance := left_list[i] - right_list[i]
		if distance < 0 {
			distance = distance * -1
		}
		total_distance += distance
	}

	// calculate how many times each left element is duplicated in the right list
	// and determene the overall similarity
	for i := 0; i < len(left_list); i++ {
		count := 0

		for j := 0; j < len(right_list); j++ {
			if left_list[i] == right_list[j] {
				count++
			}
		}

		total_similarity += left_list[i] * count

	}

	fmt.Printf("\n Total distance: %v\n", total_distance)

	fmt.Printf("Total similarity: %v\n", total_similarity)
}
