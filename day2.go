package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var arrayRows [][]int
	var safeCount int

	fmt.Printf("Input: \n")

	scanner := bufio.NewScanner(os.Stdin)

	//
	for scanner.Scan() {
		rows := scanner.Text()

		parts := strings.Fields(rows)

		var tempRow []int

		for i := 0; i < len(parts); i++ {

			partInt, err := strconv.Atoi(parts[i])
			if err != nil {
				fmt.Printf("Error: %s\n", rows)
				continue
			}

			tempRow = append(tempRow, partInt)

		}

		arrayRows = append(arrayRows, tempRow)

	}

	// logic checking if array is sorted in either ascending or descending and if there's a difference less than 3
	for i := 0; i < len(arrayRows); i++ {
		isSafe := true
		currentRow := arrayRows[i]

		fmt.Println("Processing row:", currentRow)

		ascendingCheck := isSortedAscendingAndOneElementOut(currentRow)
		descendingCheck := isSortedDescendingAndOneElementOut(currentRow)
		fmt.Println("Is sorted ascending with one element out?", ascendingCheck)
		fmt.Println("Is sorted descending with one element out?", descendingCheck)

		if isSortedAscendingAndOneElementOut(currentRow) || isSortedDescendingAndOneElementOut(currentRow) {
			for j := 0; j < len(currentRow)-1; j++ {

				diff := currentRow[j] - currentRow[j+1]

				fmt.Printf("Difference between %d and %d: %d\n", currentRow[j], currentRow[j+1], diff)

				if diff < -3 || diff > 3 || diff == 0 {
					isSafe = false
					break
				}
			}
		} else {
			isSafe = false
		}

		if isSafe {
			fmt.Println("Row is safe")
			safeCount++
		} else {
			fmt.Println("Row is unsafe")
		}
	}

	// star 1
	fmt.Println(safeCount)
}

func isSortedAscending(rowList []int) bool {
	for i := 0; i < len(rowList)-1; i++ {
		if rowList[i] >= rowList[i+1] {
			return false
		}
	}
	return true
}

func isSortedDescending(rowList []int) bool {
	for i := 0; i < len(rowList)-1; i++ {
		if rowList[i] <= rowList[i+1] {
			return false
		}
	}
	return true
}

func isSortedAscendingAndOneElementOut(rowList []int) bool {
	if isSortedAscending(rowList) {
		return true
	}

	fmt.Println("Row is not sorted, checking for removals...")

	for i := 0; i < len(rowList); i++ {
		modifiedRow := append([]int{}, rowList[:i]...)
		modifiedRow = append(modifiedRow, rowList[i+1:]...)
		fmt.Println("Modified row after removing index", i, ":", modifiedRow)

		if isSortedAscending(modifiedRow) {
			fmt.Println("Modified row is sorted ascending.")
			return true
		}
	}
	return false
}

func isSortedDescendingAndOneElementOut(rowList []int) bool {
	if isSortedDescending(rowList) {
		return true
	}

	fmt.Println("Row is not sorted, checking for removals...")

	for i := 0; i < len(rowList); i++ {
		modifiedRow := append([]int{}, rowList[:i]...)
		modifiedRow = append(modifiedRow, rowList[i+1:]...)
		fmt.Println("Modified row after removing index", i, ":", modifiedRow)

		if isSortedDescending(modifiedRow) {
			fmt.Println("Modified row is sorted descending.")
			return true
		}
	}
	return false
}
