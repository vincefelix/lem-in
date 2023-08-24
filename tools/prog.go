package lem_in

import (
	"fmt"
	"os"
)

func Lem_in_prog() {
	filename := os.Args[1]
	validFile, err := CheckValidityFile(filename)
	if !validFile {
		fmt.Printf("ERROR: invalid data format, %v \n", err)
		return
	}
	var anthill Ant

	antsize, chambre, _, _ := parseFile(filename)
	startRoom := StartRoom(chambre)
	endRoom := EndRoom(chambre)
	Allpaths := findPathsBFS(startRoom, endRoom)
	PathOptimized := OptimizedPaths(Allpaths)
	path_tab := ConvertToString(PathOptimized)
	path_tab = noRepeat(path_tab)

	anthill = anthill.Path(path_tab)
	anthill.Ant_per_path(numAnts, path_tab)
	anthill.Reorder(path_tab, antsize)
	anthill.PrintSeq(path_tab)
}
