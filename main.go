package main

import (
	lmin "lem_in/tools"
)

func main() {
	var anthill lmin.Ant

	antsize := 2
	path_tab := [][]string{
		{""},
	}

	anthill = anthill.Path(path_tab)
	anthill.Ant_per_path(antsize, path_tab)
	anthill.Reorder(path_tab, antsize)
	anthill.PrintSeq(path_tab)

}
