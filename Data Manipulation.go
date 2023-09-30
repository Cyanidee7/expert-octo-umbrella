package main

import (
	"fmt"
)

type agent struct {
	name  string
	kills int
}

func (s agent) ChangeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *agent) ChangeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}

func main() {
	var s1 agent
	s1.name = "jett"
	s1.kills = 27

	var s2 agent
	s2.name = "killjoy"
	s2.kills = 27

	var s3 agent
	s3.name = "astra"
	s3.kills = 27

	fmt.Println("s1 before", s1.name)

	s1.ChangeName1(s2.name)
	fmt.Println("s1 after changeName1", s2.name)

	s1.ChangeName2("s3.name")
	fmt.Println("s2 after changeName2", s3.name)

}
