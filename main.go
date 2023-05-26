package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("vim-go")

	// construct graph of courses from input
	//
	courses := make(Courses)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "DONE" {
			break
		}
		if line == "" {
			continue
		}
		course, prereqs := parseCourse(line)
		//fmt.Println(course, prereqs)

		// save course
		courses.Add(course)
		courses.AddPrereqs(course, prereqs...)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// debug, check graph
	/*for k, v := range courses {
		fmt.Println(k, v)
	}*/

	// print sequence of prereqs for each course
	for _, course := range courses {
		fmt.Printf("Path of %s: ", course.ID)
		displayPath(course)
		fmt.Println()

	}

}
