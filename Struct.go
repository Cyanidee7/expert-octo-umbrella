package main

import (
	"fmt"
	"strings"
)

type agent struct {
	name  string
	kills int
}

func main() {
	var s1 agent
	s1.name = "jett"
	s1.kills = 27

	var s2 agent
	s2.name = "chamber"
	s2.kills = 14

	fmt.Println("name	:", s1.name)
	fmt.Println("kills	:", s1.kills)
}

func (s agent) Confirmation() {
	fmt.Println("choosen", s.name)
}

func (s agent) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}
