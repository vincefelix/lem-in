package main

import (
	"fmt"
	lmin "lem_in/tools"
)

func main() {
	var anthill lmin.Ant

	antsize := 50
	path_tab := [][]string{
		{"0", "2", "1"},
		{"0", "3", "1"},
	}

	anthill = anthill.Path(path_tab)
	anthill.Ant_per_path(antsize, path_tab)
	anthill.Reorder(path_tab)
	fmt.Println(anthill)
	anthill.PrintSeq(path_tab)

}
