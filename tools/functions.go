package lem_in

import "strconv"

// it sorts the array in ascending name order
func Sort(table []string) {
	var t string
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table)-1; j++ {
			prev := GetNumber(table[j])
			next := GetNumber(table[j+1])

			if prev > next {
				t = table[j+1]
				table[j+1] = table[j]
				table[j] = t
			}

		}
	}
}

func Maxpath(pathtab [][]string) int {
	size := len(pathtab[0])
	for i := range pathtab {
		if len(pathtab[i]) > size {
			size = len(pathtab[i])
		}
	}
	return size
}

func sumOfColumn(room [][]int) []int {
	var columnSizes []int

	if len(room) == 0 {
		return columnSizes
	}

	maxRowLength := 0
	for _, row := range room {
		if len(row) > maxRowLength {
			maxRowLength = len(row)
		}
	}

	for j := 0; j < maxRowLength; j++ {
		columnSize := 0
		for i := 0; i < len(room); i++ {
			if j < len(room[i]) {
				columnSize++
			}
		}
		columnSizes = append(columnSizes, columnSize)
	}

	return columnSizes
}

// it trims and return the between the <L> and <->
func GetNumber(n string) int {
	number := ""
	for i, v := range n {
		if v == '-' {
			number = n[1:i]
			break
		}
	}

	if number == "" {
		number = n[1:]
	}
	newnumber, err := strconv.Atoi(number)
	if err != nil {
		return 0
	}
	return newnumber

}

func isNumber(value string) bool {
	num, err := strconv.Atoi(value)

	if err != nil && num == 0 {
		return false
	}
	return true
}
func Atoi(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
