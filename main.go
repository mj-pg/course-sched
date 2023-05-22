package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Graph map[interface{}]*node

type node struct {
	Val interface{}
	Adj []*node
}

func (g Graph) Get(val interface{}) *node {
	return g[val]
}

func (g Graph) Add(val interface{}) *node {

	// val exists
	if n := g.Get(val); n != nil {
		return n
	}

	// create node for new val
	n := &node{
		Val: val,
	}
	g[val] = n
	return n
}

func (g Graph) AddAdj(val, adj interface{}) *node {

	// get val node
	n := g.Get(val)

	// val is new
	if n == nil {
		n = &node{
			Val: val,
		}
		g[val] = n
	}

	// get adj node
	nAdj := g.Get(adj)

	// adj is new
	if nAdj == nil {
		nAdj = &node{
			Val: adj,
		}
		g[adj] = nAdj
	}
	// add adjacent node
	n.Adj = append(n.Adj, nAdj)
	return n
}

type Course struct {
	ID   string
	desc string
}

func main() {
	fmt.Println("vim-go")

	// construct graph of courses from input
	//
	courses := make(Graph)
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
		fmt.Println(course, prereqs)

		// save course
		courses.Add(course)
		for _, prereq := range prereqs {
			courses.AddAdj(course, prereq)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// debug
	for k, v := range courses {
		fmt.Println(k, v)
	}

}

func parseCourse(str string) (Course, []Course) {
	// split COURSE: PREREQ LIST
	//
	tokens := strings.Split(str, ":")
	if len(tokens) < 1 {
		return Course{}, nil
	}

	course := Course{
		ID: strings.TrimSpace(tokens[0]),
	}

	// without prereq
	//
	if len(tokens) < 2 {
		return course, nil
	}
	tokens[1] = strings.TrimSpace(tokens[1])
	if tokens[1] == "" {
		return course, nil
	}

	// with prereq
	//
	tokens = strings.Split(tokens[1], ",")
	var prereqs []Course
	for _, courseID := range tokens {
		prereqs = append(prereqs, Course{
			ID: strings.TrimSpace(courseID),
		})
	}
	return course, prereqs
}
