package utils

import (
	"strconv"
	"unicode"
)

// TentativePartition gives an approximate number of net chunks we can expect
func TentativePartition(n int) int {
	n = n / 50
	res := 0
	a := (n / 10) * 10
	b := a + 10

	if n-a > b-n { //nearest 10s round off
		res = b
	} else {
		res = a
	}
	return res
}

// CountAndPartition takes in a string and returns the array required and the total number of partitions
func CountAndPartition(str string) ([]string, int) {
	lastSpacePos := -1
	var tempArray []string
	length := 0
	max := 50
	lastwordBreakPos := 0
	numberOfPartition := 0
	tentativePartition := TentativePartition(len(str))
	toSaveSpace := tentativePartition + 3

	for i, r := range str {
		if unicode.IsSpace(r) && (i+1) != len(str) {
			lastSpacePos = i
		}
		extraSpaceforTotApproxAnnotations := (len(strconv.Itoa(numberOfPartition)))
		// println("extraSpaceforAnnotations", extraSpaceforTotApproxAnnotations)
		length++
		if length >= (max-toSaveSpace)-extraSpaceforTotApproxAnnotations {
			// println("current position of i: ", i)
			// if we found a space earlier and the current rune is not a space
			if lastSpacePos != -1 && lastwordBreakPos != lastSpacePos {
				// println("max length reached ::: Partitioning")
				numberOfPartition += 1
				temp := str[lastwordBreakPos:lastSpacePos]
				tempArray = append(tempArray, temp)
				i = lastSpacePos
				// println(temp)
				// println("now i is at ", i)
				lastwordBreakPos = i
				temp = ""
				// println("last word break at: ", lastwordBreakPos)
				length = 0
			}
			// reached the end of string but not the chunk limit
		} else if len(str)-1 == i {

			// println("outer loop i value", i)
			temp := str[lastwordBreakPos : i+1]
			tempArray = append(tempArray, temp)
			// println("Extra", temp)
		}
	}
	// no spaces encountered
	if lastSpacePos == -1 {
		println(str)
	}
	return tempArray, numberOfPartition
}
