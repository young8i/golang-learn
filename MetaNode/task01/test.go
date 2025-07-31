package main

import "fmt"

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

func (g *Gender) String() string {
	switch *g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Unknown"
	}
}

func (g *Gender) IsMale() bool {
	return *g == Male
}

const pre int = 1
const a int = iota
const (
	b int = iota
	c
	d
	e
)
const (
	f = 2
	g = iota
	h
	i
)

//func main() {
//	//g := Male
//	//s := g.String()
//	//fmt.Println(s)
//}

func main() {
	var a int = 10
	if b := 1; a > 10 {
		b = 2
		// c = 2
		fmt.Println("a > 10")
	} else if c := 3; b > 1 {
		b = 3
		fmt.Println("b > 1")
	} else {
		fmt.Println("其他")
		if c == 3 {
			fmt.Println("c == 3")
		}
		fmt.Println(b)
		fmt.Println(c)
	}
}
