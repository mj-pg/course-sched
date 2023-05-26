package main

import (
	"fmt"
	"strings"
)

type CourseID string
type Courses map[CourseID]*Course

type Course struct {
	ID      CourseID
	Desc    string
	Prereqs []*Course
}

func (cc Courses) Add(c Course) *Course {

	// course already exists
	if exist, ok := cc[c.ID]; ok {
		return exist
	}

	// save new course
	cc[c.ID] = &c
	return cc[c.ID]
}

func (cc Courses) AddPrereqs(c Course, prereqs ...Course) {

	// course is new
	if cc[c.ID] == nil {
		cc[c.ID] = &c
	}

	course := cc[c.ID]
	for _, prereq := range prereqs {
		added := cc.Add(prereq)
		course.addPrereq(added)
	}
}

func (c *Course) addPrereq(prereq *Course) {
	for _, exist := range c.Prereqs {
		if prereq == exist {
			return
		}
	}

	c.Prereqs = append(c.Prereqs, prereq)
}

func displayPath(c *Course) {

	if c == nil {
		return
	}
	cc := path(c)
	for _, c := range cc {
		fmt.Printf("%s ", c.ID)
	}
}

func path(c *Course) []*Course {
	var ret []*Course
	for _, p := range c.Prereqs {
		prereqs := path(p)
		ret = append(ret, prereqs...)
	}
	ret = append(ret, c)
	return ret
}

func parseCourse(str string) (Course, []Course) {
	// split course from prereq list
	//
	tokens := strings.Split(str, ":")
	if len(tokens) < 1 {
		return Course{}, nil
	}

	course := Course{
		ID: CourseID(strings.TrimSpace(tokens[0])),
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
			ID: CourseID(strings.TrimSpace(courseID)),
		})
	}
	return course, prereqs
}
