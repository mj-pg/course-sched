package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Graph map[interface{}]*Node

type Node struct {
	val interface{}
	adj []interface{}
}

type Course struct {
	ID   string
	desc string
}

func main() {
	fmt.Println("vim-go")

	// parse input
	// course: prereq1, prereq2
	//
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "DONE" {
			break
		}
		if line == "" {
			continue
		}
		course, prereqs := toCourse(line)
		fmt.Println(course, prereqs)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func toCourse(str string) (Course, []Course) {
	tokens := strings.Split(str, ":")
	if len(tokens) < 1 {
		return Course{}, nil
	}

	course := Course{
		ID: strings.TrimSpace(tokens[0]),
	}

	if len(tokens) < 2 {
		return course, nil
	}
	tokens[1] = strings.TrimSpace(tokens[1])
	if tokens[1] == "" {
		return course, nil
	}

	tokens = strings.Split(tokens[1], ",")
	var prereqs []Course
	for _, courseID := range tokens {
		prereqs = append(prereqs, Course{
			ID: strings.TrimSpace(courseID),
		})
	}
	return course, prereqs
}
