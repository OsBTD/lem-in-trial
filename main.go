package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type farm struct {
	ants_number int
	rooms       map[string][]int
	start       map[string][]int
	end         map[string][]int
	links       map[string][]string
}

func main() {
	var myFarm farm
	myFarm.Read("test.txt")
	BFS(myFarm)
	fmt.Println("number of ants is : ", myFarm.ants_number)
	fmt.Println("rooms are : ", myFarm.rooms)
	fmt.Println("start is : ", myFarm.start)
	fmt.Println("end is : ", myFarm.end)
	fmt.Println("links are : ", myFarm.links)
	fmt.Println("adjacent is : ", Graph(myFarm))
}

func (myFarm *farm) Read(filename string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Println("error reading", err)
	}
	content := strings.Split(string(bytes), "\n")

	myFarm.rooms = make(map[string][]int)
	myFarm.start = make(map[string][]int)
	myFarm.end = make(map[string][]int)
	myFarm.links = make(map[string][]string)

	var st, en int
	number, err := strconv.Atoi(content[0])
	if err != nil {
		log.Println("couldn't convert", err)
	}
	myFarm.ants_number = number

	for index := range content {
		if strings.TrimSpace(content[index]) == "##start" {
			st++
			if index+1 <= len(content)-1 {
				split := strings.Split(strings.TrimSpace(content[index+1]), " ")
				x, err := strconv.Atoi(split[1])
				y, err2 := strconv.Atoi(split[2])
				if err == nil && err2 == nil {
					myFarm.start[split[0]] = []int{x, y}
				}

			}

		} else if strings.TrimSpace(content[index]) == "##end" {
			en++
			if index+1 <= len(content)-1 {
				split := strings.Split(strings.TrimSpace(content[index+1]), " ")
				x, err := strconv.Atoi(split[1])
				y, err2 := strconv.Atoi(split[2])
				if err == nil && err2 == nil {
					myFarm.end[split[0]] = []int{x, y}
				}

			}
		} else if strings.Contains(content[index], "-") {
			split := strings.Split(strings.TrimSpace(content[index]), "-")
			if len(split) == 2 {
				myFarm.links[split[0]] = append(myFarm.links[split[0]], split[1])
			}
		} else if strings.Count(content[index], " ") == 2 {
			split := strings.Split(strings.TrimSpace(content[index]), " ")
			if len(split) == 3 {
				x, err := strconv.Atoi(split[1])
				y, err2 := strconv.Atoi(split[2])
				if err == nil || err2 == nil {
					myFarm.rooms[split[0]] = []int{x, y}
				}
			}
		} else if (strings.HasPrefix(strings.TrimSpace(content[index]), "#") || strings.HasPrefix(strings.TrimSpace(content[index]), "L")) && (strings.TrimSpace(content[index]) != "##start" && strings.TrimSpace(content[index]) != "##end") {
			continue
		}
	}
	if en != 1 || st != 1 {
		log.Println("rooms setup is incorrect", err)
	}
}

func Graph(farm farm) map[string][]string {
	adjacent := make(map[string][]string)
	for room := range farm.rooms {
		adjacent[room] = []string{}
	}
	for room, links := range farm.links {
		for _, link := range links {
			adjacent[room] = append(adjacent[room], link)
			adjacent[link] = append(adjacent[link], room)

		}
	}

	return adjacent
}

func BFS(myFarm farm) {
	adjacent := Graph(myFarm)
	var Queue []string
	var endd string
	start := myFarm.start
	end := myFarm.end
	Visited := make(map[string]bool)
	Parents := make(map[string]string)

	fmt.Println("\n=== Initialization ===")
	fmt.Println("Start room map:", start)
	fmt.Println("Adjacent list:", adjacent)

	// Initialize with start room
	for key := range start {
		Queue = append(Queue, key)
		Visited[key] = true
		fmt.Printf("\nAdding start room '%s' to queue\n", key)
		fmt.Printf("Current visited map: %v\n", Visited)
	}

	for key := range end {
		endd = key
		fmt.Printf("\nEnd room is: '%s'\n", endd)
	}

	fmt.Printf("\n=== Starting BFS Traversal ===\n")
	fmt.Printf("Initial queue: %v\n", Queue)

	stepCount := 1
	// Modified loop condition to handle both empty queue and end room discovery
	for len(Queue) > 0 {
		current := Queue[0]
		Queue = Queue[1:]

		fmt.Printf("\n--- Step %d ---\n", stepCount)
		fmt.Printf("Processing room: '%s'\n", current)
		fmt.Printf("Current parent map: %v\n", Parents)
		fmt.Printf("Connected to rooms: %v\n", adjacent[current])

		// Check if we've reached the end room
		if current == endd {
			fmt.Printf("\n!!! Found end room '%s' - Breaking BFS !!!\n", endd)
			break
		}

		for _, link := range adjacent[current] {
			if !Visited[link] {
				Queue = append(Queue, link)
				Visited[link] = true
				Parents[link] = current
				fmt.Printf("\nDiscovered new room: '%s'\n", link)
				fmt.Printf("Updated parent map - added: '%s' → '%s'\n", link, current)
				fmt.Printf("Current queue: %v\n", Queue)
				fmt.Printf("Current visited map: %v\n", Visited)
			} else {
				fmt.Printf("\nRoom '%s' already visited (parent: '%s')\n", link, Parents[link])
			}
		}
		stepCount++
	}

	// Check if we actually found a path
	if !Visited[endd] {
		fmt.Printf("\n=== No path found to end room ===\n")
		return
	}

	fmt.Printf("\n=== BFS Complete ===\n")
	fmt.Printf("Final parent map: %v\n", Parents)

	fmt.Printf("\n=== Path Reconstruction ===\n")
	path := []string{endd}
	current := endd
	fmt.Printf("Starting from end room: '%s'\n", current)

	for Parents[current] != "" {
		fmt.Printf("Parent of '%s' is '%s'\n", current, Parents[current])
		current = Parents[current]
		path = append([]string{current}, path...)
	}

	fmt.Printf("\nFinal path from start to end: %v\n", path)
}
