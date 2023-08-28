package lem_in

// package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type roomformat struct {
	name    string
	coord_x int
	coord_y int
}

var tabrooms []roomformat

func FileToTable(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur: nous ne parvenons pas a lire le fichier" + filename)

	} else {
		example := bufio.NewScanner(file)

		for example.Scan() {
			lines = append(lines, example.Text())
		}
	}
	return lines, err
}

func linesWithoutExtraSpaces(lines []string) (lineswithoutpaces []string) {
	for i := 0; i < len(lines); i++ {
		if lines[i] != "" {
			lineswithoutpaces = append(lineswithoutpaces, lines[i])
		}

	}
	return lineswithoutpaces
}

func trimspacesinlines(lines []string) (linestrimspaces []string) {
	for i := 0; i < len(lines); i++ {
		linestrimspaces = append(linestrimspaces, strings.TrimSpace(lines[i]))
	}
	return linestrimspaces
}

func deleteComments(lines []string) (lineswithoutcomments []string) {
	for i := 0; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "#") && lines[i][1] != '#' {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}
	lineswithoutcomments = lines
	return lineswithoutcomments
}

func dispatchingLinesAndRooms(lines []string) (rooms, links []string) {
	for i := 1; i < len(lines); i++ {
		if strings.Contains(lines[i], "-") {
			links = append(links, lines[i])
		} else {
			rooms = append(rooms, lines[i])
		}
	}
	return rooms, links
}

func StartAndEnd(lines []string) bool {
	var start, end = 0, 0
	for i := 0; i < len(lines); i++ {
		if lines[i] == "##start" {
			start++
		} else if lines[i] == "##end" {
			end++
		}
	}
	return start == 1 && end == 1
}

func ReformateRooms(rooms []string) {

	for j := len(rooms) - 1; j >= 1; j-- {
		if rooms[j] == "##start" {
			rooms[j-1], rooms[j], rooms[j+1] = rooms[j], rooms[j+1], rooms[j-1]

		}
	}
	for i := 2; i < len(rooms)-2; i++ {
		if rooms[i] == "##end" {
			rooms[i], rooms[i+1], rooms[i+2] = rooms[i+2], rooms[i], rooms[i+1]

		}
	}
}

func samecoordxy(x, y int, tabrooms []roomformat) bool {
	repeat := false
	for i := 0; i < len(tabrooms); i++ {
		room := tabrooms[i]

		if room.coord_x == x && room.coord_y == y {
			fmt.Println("x", x, "y", y)
			repeat = true
		}
	}
	return repeat
}

func RoomAndLinksFormat(rooms, links []string) (bool, string) {
	if len(rooms) == 0 || len(links) == 0 {
		err := "no rooms/links found"
		return false, err
	}
	for i := 0; i < len(rooms); i++ {
		if i == len(rooms)-1 {
			if rooms[i] == "##start" || rooms[i] == "##end" {
				return false, "no room for ##start/##end"
			}
		} else {
			if (rooms[i] == "##start" && rooms[i+1] == "##end") || (rooms[i] == "##end" && rooms[i+1] == "##start") {
				fmt.Println("here", rooms, i, i+1)
				return false, "##start & ##end must not follow each other"
			}
		}
	}
	//reformate Rooms
	ReformateRooms(rooms)

	//delete command ##start & ##end
	for i := 0; i < len(rooms); i++ {
		if strings.HasPrefix(rooms[i], "#") {
			rooms = append(rooms[:i], rooms[i+1:]...)
		}
	}

	names := make(map[string]bool)
	// fmt.Println(rooms, "rooms")
	for _, room := range rooms {

		roomTable := strings.Split(room, " ")

		if len(roomTable) != 3 {
			err := "Error invalid room format:" + string(room[0])
			return false, err
		} else {
			if names[roomTable[0]] {
				err := "The room " + roomTable[0] + " already exists"
				return false, err
			} else {
				names[roomTable[0]] = true
			}
			x, errx := strconv.Atoi(roomTable[1])
			y, erry := strconv.Atoi(roomTable[2])
			if strings.HasPrefix(roomTable[0], "L") || strings.HasPrefix(roomTable[0], "#") {
				err := "A room must begin with neither L nor #"
				return false, err

			} else if errx != nil || erry != nil {
				err := "x & y coordinates must be numeric"
				return false, err
			} else {
				if samecoordxy(x, y, tabrooms) {
					return false, "the rooms must not have the same x y coordinates" + roomTable[0] + " " + strconv.Itoa(x) + " " + strconv.Itoa(y)
				}
				chambre := roomformat{name: roomTable[0], coord_x: x, coord_y: y}
				tabrooms = append(tabrooms, chambre)

			}

		}

	}

	//check links
	for i := 0; i < len(links); i++ {
		link := strings.Split(links[i], "-")
		if len(link) != 2 {
			err := "invalid link format "
			return false, err
		} else {
			if link[0] == link[1] {
				err := "room connected to itself "
				return false, err
			} else if !names[link[0]] || !names[link[1]] {
				err := "non-existent room name "
				return false, err
			}
		}
	}

	return true, ""
}
func CreateAndWriteInAFile(output string, finalContent []string) {
	//création du fichier de sortie
	newFile, _ := os.Create(output)
	//écriture dans le nouveau fichier
	// writeInFile, _ := newFile.WriteString(finalContent)
	for i := 0; i < len(finalContent); i++ {
		if i != len(finalContent)-1 {
			_, err := newFile.WriteString(finalContent[i] + "\n")
			if err != nil {
				fmt.Println("error detected")
			}
		} else {
			_, err := newFile.WriteString(finalContent[i])
			if err != nil {
				fmt.Println("error detected")
			}
		}
	}

	// fmt.Printf("last: %v\n", writeInFile)

}
func reconstruction(rooms []string, links []string, number string) (tablines []string) {
	tablines = append(tablines, number)
	for i := 0; i < len(rooms); i++ {
		if i == 0 {
			tablines = append(tablines, "##start"+"\n"+rooms[i])
		} else if i == len(rooms)-1 {
			tablines = append(tablines, "##end"+"\n"+rooms[i])

		} else {
			tablines = append(tablines, rooms[i])

		}
	}
	for i := 0; i < len(links); i++ {
		tablines = append(tablines, links[i])

	}

	return tablines
}
func CheckValidityFile(lines []string) (validity bool, answer string) {
	// fmt.Println(lines,"first")
	//delete empty lines
	lines = linesWithoutExtraSpaces(lines)
	// fmt.Println(lines, "delete /n")
	//trim extra spaces in lines
	lines = trimspacesinlines(lines)
	// fmt.Println(lines, "trim space")
	//delete all the comments
	lines = deleteComments(lines)

	// fmt.Println(lines, "delete comments")
	//assign the first line of the table as the number of ants
	number_of_ants, _ := strconv.Atoi(lines[0])
	// fmt.Println("number of ants", number_of_ants)
	if number_of_ants > 0 {
		if StartAndEnd(lines) {
			rooms, links := dispatchingLinesAndRooms(lines)
			validity, answer = RoomAndLinksFormat(rooms, links)
			// fmt.Println("rooms", rooms)
			// fmt.Println("links", links)
			tab := reconstruction(rooms, links, lines[0])
			CreateAndWriteInAFile("newfilename.txt", tab)

		} else {
			validity, answer = false, "the ##start & ##end commands must each be 1."
		}

	} else {
		validity, answer = false, "the number of ants must be in the first position, be numerical and exceed 0"

	}
	return validity, answer
}

// func main() {
// 	if len(os.Args) == 2 {
// 		lines, error := FileToTable(os.Args[1])
// 		lines = linesWithoutExtraSpaces(lines)
// 		lines = deleteComments(lines)
// 		if len(lines) == 0 && error == nil {
// 			fmt.Println("empty file")
// 		} else if error == nil && len(lines) > 0 {
// 			valid, answer := CheckValidityFile(lines)
// 			if valid {
// 				fmt.Println(valid)
// 			} else {
// 				fmt.Println(answer)
// 			}
// 		}
// 	} else {
// 		fmt.Println("Trop ou peu d'arguments")
// 	}

// }
