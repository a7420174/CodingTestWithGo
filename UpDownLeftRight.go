package main

import (
	"fmt"
	"strings"
)

type Vertex struct{
	x, y int
}

func main() {
	var input string
	N := 5
	input = "R R R L L U U D D D D D D D"
	moves := strings.Split(input, " ")
	V := Vertex{1, 1}
	
	mapping := map[string]Vertex{
		"L" : {-1, 0},
		"R" : {1, 0},
		"U" : {0, -1},
		"D" : {0, 1},
	}
	
	for _, move := range moves {
		nV := Vertex{
			V.x + mapping[move].x,
			V.y + mapping[move].y,
		}
		if nV.x < 1 || nV.x > N || nV.y < 1 || nV.y > N {
			continue
		}
		V = nV
	}
	
	
	fmt.Println(V)
}
