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
	safeCount := 0

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

		ascendingCheck := isSortedAscendingAndOneElementOut(currentRow)
		fmt.Printf("ascending order: %v\n", ascendingCheck)
		descendingCheck := isSortedDescendingAndOneElementOut(currentRow)
		fmt.Printf("descending order: %v\n", descendingCheck)

		if ascendingCheck != nil {
			for j := 0; j < len(ascendingCheck)-1; j++ {
				diff := ascendingCheck[j+1] - ascendingCheck[j] // Ascending difference
				if diff > 3 || diff < 1 {
					isSafe = false
					break
				}
			}
		} else if descendingCheck != nil {
			for j := 0; j < len(descendingCheck)-1; j++ {
				diff := descendingCheck[j] - descendingCheck[j+1] // Descending difference
				if diff > 3 || diff < 1 {
					isSafe = false
					break
				}
			}
		} else {
			isSafe = false
			fmt.Printf("Row %d is unsafe: %v\n", i, currentRow)
		}

		if isSafe {
			safeCount++
		}
	}

	// star 1
	fmt.Println(safeCount)
}

func isSortedAscending(rowList []int) []int {
	for i := 0; i < len(rowList)-1; i++ {
		if rowList[i] >= rowList[i+1] {
			return nil
		}
	}
	return rowList
}

func isSortedDescending(rowList []int) []int {
	for i := 0; i < len(rowList)-1; i++ {
		if rowList[i] <= rowList[i+1] {
			return nil
		}
	}
	return rowList
}

func isSortedAscendingAndOneElementOut(rowList []int) []int {
	if isSortedAscending(rowList) != nil {
		return rowList
	}

	for i := 0; i < len(rowList); i++ {
		modifiedRow := append([]int{}, rowList[:i]...)
		modifiedRow = append(modifiedRow, rowList[i+1:]...)

		if isSortedAscending(modifiedRow) != nil {
			return modifiedRow
		}
	}
	return nil
}

func isSortedDescendingAndOneElementOut(rowList []int) []int {
	if isSortedDescending(rowList) != nil {
		return rowList
	}

	for i := 0; i < len(rowList); i++ {
		modifiedRow := append([]int{}, rowList[:i]...)
		modifiedRow = append(modifiedRow, rowList[i+1:]...)

		if isSortedDescending(modifiedRow) != nil {
			return modifiedRow
		}
	}
	return nil
}
