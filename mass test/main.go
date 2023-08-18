package main

import (
	"fmt"
	"sort"
)

type Paths [][]string

// -----sort's side functions---------//
func (p *Paths) Len() int {
	return len(*p)
}

func (p Paths) Less(i, j int) bool {
	return len(p[i]) < len(p[j])
}

func (p Paths) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Paths) Sortpaths() {
	sort.Sort(p)
}

// -------------------------------//
func ant_per_path(pathtab Paths, antsize int) []int {
	passing_ant := make([]int, pathtab.Len())
	roomsize := make([]int, pathtab.Len())
	for i := range roomsize {
		roomsize[i] = len(pathtab[i])
	}

	index := 0
	for i := 0; i < antsize; i++ {
		roomspace := roomsize[index] + 1 + passing_ant[index]
		if roomsize[index] != roomsize[len(roomsize)-1] {
			if roomspace <= roomsize[index+1] {
				passing_ant[index]++
			} else {
				passing_ant[index+1]++
				index++
			}
		} else {
			passing_ant[0]++

		}
	}
	//removing the nil slices
	for i, v := range passing_ant {
		if v == 0 {
			if i < len(passing_ant)-1 {
				passing_ant = append(passing_ant[:i], passing_ant[i+1:]...)
			} else if i == len(passing_ant)-1 {
				passing_ant = passing_ant[:i]
			}
		}
	}

	return passing_ant
}

func Sequences(antsize int, tab Paths) {
	mirror := 0
	passers := ant_per_path(tab, antsize)
	fmt.Println("ants per path")
	fmt.Println(passers)
	initial := antsize
	for i := 1; initial >= 0; i++ {
		for j := 0; j < tab.Len(); j++ {
			if i < len(tab[j]) {
				if j < len(tab[j]) {
					if i == 1 {
						step := fmt.Sprintf("L%v-L%s ", j+1, tab[j][i])
						fmt.Print(step)
					} else {
						var add_on []string
						for k := 0; k < mirror; k++ {
							add_on = append(add_on, "*")
						}
						for _, v := range tab[mirror] {
							add_on = append(add_on, v)
						}
						tab = append(tab, tab[mirror])
						mirror++
						step := fmt.Sprintf("L%v-L%s ", j+1, tab[j][i])
						fmt.Print(step)

					}
				}
			}

		}

		initial--
		println()
	}
}

func main() {
	tab := Paths{
		{"1", "2", "5", "6", "0"},
		{"1", "3", "4", "0"},
	}

	fmt.Println("-----before sorting---------")
	fmt.Println(tab)
	fmt.Println("-----after sorting---------")
	tab.Sortpaths()
	fmt.Println(tab)
	fmt.Println("-----sequences---------")
	Sequences(3, tab)
}
