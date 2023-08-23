package lem_in

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseFile(filename string) (int, []Room, []Link, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "##start") {
			scanner.Scan()
			roomLine := scanner.Text()
			roomParts := strings.Fields(roomLine)
			if len(roomParts) != 3 {
				return 0, nil, nil, fmt.Errorf("invalid room format: %v", roomLine)
			}
			Rooms = append(Rooms, Room{Name: roomParts[0], X: Atoi(roomParts[1]), Y: Atoi(roomParts[2]), Start: true})
		} else if strings.HasPrefix(line, "##end") {
			scanner.Scan()
			roomLine := scanner.Text()
			roomParts := strings.Fields(roomLine)
			if len(roomParts) != 3 {
				return 0, nil, nil, fmt.Errorf("invalid room format: %v", roomLine)
			}
			Rooms = append(Rooms, Room{Name: roomParts[0], X: Atoi(roomParts[1]), Y: Atoi(roomParts[2]), End: true})
		} else if !strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "L") && !strings.Contains(line, "-") && !isNumber(line) {
			parts := strings.Fields(line)
			if len(parts) == 3 {
				Rooms = append(Rooms, Room{Name: parts[0], X: Atoi(parts[1]), Y: Atoi(parts[2])})
			} else if len(parts) == 2 {
				links = append(links, Link{Room1: parts[0], Room2: parts[1]})
			}
		} else if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			room1 := findRoom(Rooms, parts[0])
			room2 := findRoom(Rooms, parts[1])
			if room1 != nil && room2 != nil {
				room1.Links = append(room1.Links, room2)
				room2.Links = append(room2.Links, room1)
			}
		} else if isNumber(line) {
			num := Atoi(line)
			numAnts = num
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("there is an erron while parsing file: %v\n", err)
		return 0, nil, nil, err
	}

	return numAnts, Rooms, links, nil
}
func findPathsBFS(startRoom *Room, endRoom *Room) [][]*Room {
	var paths [][]*Room
	queue := [][]*Room{{startRoom}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		currentRoom := path[len(path)-1]

		if currentRoom.Name == endRoom.Name {
			copiedPath := make([]*Room, len(path))
			copy(copiedPath, path)
			paths = append(paths, copiedPath)
		}

		for _, neighbor := range currentRoom.Links {
			// Check if the neighbor is not already in the current path
			alreadyInPath := false
			for _, room := range path {
				if room == neighbor {
					alreadyInPath = true
					break
				}
			}
			if !alreadyInPath {
				newPath := append([]*Room{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return paths
}
func StartRoom(rooms []Room) *Room {
	for _, room := range rooms {
		if room.Start {
			return &room
		}
	}
	return nil
}

func EndRoom(rooms []Room) *Room {
	for _, room := range rooms {
		if room.End {
			return &room
		}
	}
	return nil
}

func findRoom(rooms []Room, name string) *Room {
	for i := range rooms {
		if rooms[i].Name == name {
			return &rooms[i]

		}
	}
	return nil
}
func FindPathsBFS(startRoom *Room, endRoom *Room) [][]*Room {
	var paths [][]*Room
	queue := [][]*Room{{startRoom}}
	fmt.Println(startRoom.Name)
	fmt.Println(endRoom.Name)

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		currentRoom := path[len(path)-1]

		if currentRoom.Name == endRoom.Name {
			copiedPath := make([]*Room, len(path))
			copy(copiedPath, path)
			paths = append(paths, copiedPath)
		}

		for _, neighbor := range currentRoom.Links {
			// Check if the neighbor is not already in the current path
			alreadyInPath := false
			for _, room := range path {
				if room == neighbor {
					alreadyInPath = true
					break
				}
			}
			if !alreadyInPath {
				newPath := append([]*Room{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return paths
}
func ConvertToString(rooms [][]*Room) [][]string {
	var stringRooms [][]string
	var stringRoom []string

	for _, way := range rooms {
		for _, room := range way {
			stringRoom = append(stringRoom, room.Name)
		}
		stringRooms = append(stringRooms, stringRoom)
	}

	return stringRooms
}

func OptimizedPaths(rooms [][]*Room) [][]*Room {
	var noCollision [][]*Room
	for _, room := range rooms {
		if !HasCollision(room, noCollision) {
			noCollision = append(noCollision, room)
		}
	}

	return noCollision
}

func noRepeat(t [][]string) [][]string {
	fmt.Println("initial tab : ", t)
	size := 0
	for i := range t {
		fmt.Println("before: ", t[i])
		if i > 0 {
			size += len(t[i-1])
			t[i] = t[i][size:]
		}
		fmt.Println("After: ", t[i])
	}
	return t
}

// func deleteRepeatition(t [][]string) [][]string {

// }

func HasCollision(room []*Room, Ways [][]*Room) bool {
	for _, Way := range Ways {
		if IsCollision(room, Way) {
			return true
		}
	}
	return false
}

func IsCollision(room1, room2 []*Room) bool {
	for i := 0; i < len(room1)-1; i++ {
		for j := 0; j < len(room2)-1; j++ {
			if room1[i] == room2[j] && room1[i+1] == room2[j+1] {
				return true
			}
		}
	}
	return false
}
