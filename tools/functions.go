package lem_in

import "strconv"

// func Sort(table []string, pathsize int)
//
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

//func GetNumber(n string) int
//
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
