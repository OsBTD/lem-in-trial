package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type farm struct {
	ants_number int
	rooms       string
	start       string
	end         string
	links       []string
}

func Read() {
	bytes, err := os.ReadFile("test.txt")
	if err != nil {
		log.Println("error reading", err)
	}
	content := strings.Split(string(bytes), "\n")
	var ants, start, end, room, links farm
	var st, en int
	number, err := strconv.Atoi(content[0])
	if err != nil {
		log.Println("couldn't convert", err)
	}
	ants.ants_number = number

	for index := range content {
		if strings.TrimSpace(content[index]) == "##start" {
			start.start = strings.TrimSpace(content[index+1])
			st++

		} else if strings.TrimSpace(content[index]) == "##end" {
			end.end = strings.TrimSpace(content[index+1])
			en++
		} else if strings.Contains(content[index], "-") {
			links.links = strings.Split(strings.TrimSpace(content[index]), "-")
		} else if (strings.HasPrefix(strings.TrimSpace(content[index]), "#") || strings.HasPrefix(strings.TrimSpace(content[index]), "L")) && (strings.TrimSpace(content[index]) != "##start" && strings.TrimSpace(content[index]) != "##end") {
			continue
		}
	}
}
