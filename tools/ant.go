package lem_in

import (
	"fmt"
	"strconv"
)

// struct used to store all ants information
type Ants struct {
	Path          []string
	Ant_nbr       int
	Passing_order []int
}

type Ant []Ants

// Sequence method returns a map with all the passage sequences of the ant colony
func (anthill Ant) Sequence() map[int][]string {

	sequences := make(map[int][]string)
	for i := 0; i < len(anthill); i++ {
		ant_inf := anthill[i]
		count := 1
		//ranging the tab storing the passage order
		for y := 0; y < len(anthill[i].Passing_order); y++ {
			for i := 1; i < len(ant_inf.Path); i++ {
				j := i - 1
				if count > 1 {
					j = i + count - 2
				}

				//formatting the sequence output : L <ant_number> - <room_name>
				s := "L" + strconv.Itoa(ant_inf.Passing_order[y]) + "-" + ant_inf.Path[i]
				sequences[j] = append(sequences[j], s)
			}
			count++
		}

	}
	return sequences
}

// Ant_per_path determines the paths with few steps, the number of ants per path,
// gives them number as ID
func (anthill Ant) Ant_per_path(antsize int, tab [][]string) {
	if len(tab) > 1 { //multiple single path case
		//--adding an ant to the first path
		anthill[0].Ant_nbr = 1        //*increasing the ant number of the current path
		antsize -= anthill[0].Ant_nbr //** decrasing the ant number of the colony

		//***adding the ant to path and giving him the index as name
		anthill[0].Passing_order = append(anthill[0].Passing_order, 1)

		//--designating each ant to path
		for antsize > 0 {
			for i := 1; i < len(tab); {
				nbr := antsize + 1

				if nbr == 0 { //* stops the loop when the there's no more roomless ant
					break
				}
				//** comparing paths length
				next_col := len(anthill[i].Path) + anthill[i].Ant_nbr
				prev_col := len(anthill[i-1].Path) + anthill[i-1].Ant_nbr
				if next_col >= prev_col { //*** we keep designating the ants to previous path
					anthill[i-1].Ant_nbr += 1
					anthill[i-1].Passing_order = append(anthill[i-1].Passing_order, nbr)

				} else if next_col < prev_col { //**** ants are designated to the next path
					if nbr > 1 {
						anthill[i].Ant_nbr += 1
						anthill[i].Passing_order = append(anthill[i].Passing_order, nbr)
					}
					i++

				}

				antsize -= 1 //***** decrasing the ant number of the colony
			}

		}
	}
	//--single path case
	//giving it the colony size and name ants from index 1 to the ants number
	anthill[0].Ant_nbr += antsize - 1
	for i := 1; i <= antsize; i++ {
		anthill[0].Passing_order = append(anthill[0].Passing_order, i)
	}

}

// it reorders the colony's movement by changing their names to entrance order
func (anthill Ant) Reorder(path_tab [][]string, antsize int) {
	if len(path_tab) > 1 { //multiple path case
		for i := range anthill {
			index := i + 1
			var new_passing_order []int

			for v := range anthill[i].Passing_order {
				if v == 0 {
					new_passing_order = append(new_passing_order, index)
				} else {
					//by adding the paths number to each ant's passing order index
					// we get his entrance number
					index += len(path_tab)
					if index > antsize {
						index = antsize
					}

					new_passing_order = append(new_passing_order, index)
				}
			}

			anthill[i].Passing_order = new_passing_order
		}
	}

}

// PrintSeq prints the sequences after sorting the sequences after we reordered the moves
func (anthill Ant) PrintSeq(path_tab [][]string) {
	steps := anthill.Sequence()

	for i := 0; i < len(steps); i++ {
		slice := steps[i]

		Sort(slice)

		for j := 0; j < len(slice); j++ {
			if j == len(slice)-1 {
				fmt.Print(slice[j])
			} else {
				fmt.Print(slice[j], " ")
			}
		}
		fmt.Println()
	}

}

// Path stores and returns all paths in the Ant.Paths field of the Ant struct
func (anthill Ant) Path(Path_tab [][]string) Ant {
	for i := 0; i < len(Path_tab); i++ {
		pathant := Ants{Path: Path_tab[i], Ant_nbr: 0}
		anthill = append(anthill, pathant)
	}
	return anthill
}
