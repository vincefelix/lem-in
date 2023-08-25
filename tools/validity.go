package lem_in

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var lastline bool

func FileToTable(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return nil
	} else {
		example := bufio.NewScanner(file)

		for example.Scan() {
			lines = append(lines, example.Text())
		}
	}
	return lines
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
	// fmt.Println(rooms)
	if rooms[len(rooms)-1] == "##end" {
		lastline = true
	}
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
	// fmt.Println(rooms)
}
func Classification(lines []string) (rooms, links []string) {
	for i := 1; i < len(lines); i++ {
		if strings.Contains(lines[i], "-") {
			links = append(links, lines[i])
		} else {
			rooms = append(rooms, lines[i])
		}
	}
	ReformateRooms(rooms)
	//delete start and end commmand
	rooms = append(rooms[:0], rooms[1:]...)

	if rooms[len(rooms)-1] == "##end" {
		rooms = rooms[:len(rooms)-1]
	} else {
		rooms = append(rooms[:len(rooms)-2], rooms[len(rooms)-1:]...)
	}
	//delete comment and other command
	for i := 0; i < len(rooms); i++ {
		if strings.HasPrefix(rooms[i], "#") {
			rooms = append(rooms[:i], rooms[i+1:]...)
		}
	}

	return rooms, links
}
func RoomAndLinksFormat(rooms, links []string) (bool, string) {
	//name x y
	names := make(map[string]bool)
	for _, room := range rooms {
		roomTable := strings.Split(room, " ")
		if len(roomTable) != 3 {
			return false, "invalid room"
		} else {
			if names[roomTable[0]] {
				return false, "invalid room"
			} else {
				names[roomTable[0]] = true
			}
			_, errx := strconv.Atoi(roomTable[1])
			_, erry := strconv.Atoi(roomTable[2])
			if (strings.HasPrefix(roomTable[0], "L") || strings.HasPrefix(roomTable[0], "#")) || errx != nil || erry != nil {
				return false, "invalid room"
			}

		}

	}
	//check links
	for i := 0; i < len(links); i++ {
		link := strings.Split(links[i], "-")
		// fmt.Println(link,len(link))
		if len(link) != 2 {
			return false, "invalid link format"
		} else {
			if link[0] == link[1] {
				return false, ""
			} else if !names[link[0]] || !names[link[1]] {
				return false, "invalid link format"
			}
		}
	}

	return true, ""
}

func CheckValidityFile(filename string) (valide bool, err string) {
	lines := FileToTable(filename)
	if lines == nil {
		return false, "wrong file name"
	}
	number_of_ants, _ := strconv.Atoi(lines[0])
	rooms, links := Classification(lines)
	// validity := false
	// fmt.Println(rooms)

	if !lastline {
		if (number_of_ants > 0) && StartAndEnd(lines) {
			validity, message := RoomAndLinksFormat(rooms, links)
			valide, err = validity, message
		}

	}
	return valide, err
}
