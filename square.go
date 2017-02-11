package main

import "strconv"

type LatinSquare struct {
	body [][]int
}

func NewLatinSquare(size int) LatinSquare {
	var body [][]int

	for i := 0; i < size; i++ {
		row := make([]int, size)
		for j := 0; j < size; j++ {
			row[j] = -1
		}
		body = append(body, row)
	}

	return LatinSquare{body}
}

func (square LatinSquare) get(x, y int) int {
	return square.body[x][y]
}

func (square LatinSquare) getPossibilities(x, y int) []int {
	size := len(square.body)
	var invalid, possibilities []int

	for i := x; i >= 0; i-- {
		invalid = append(invalid, square.get(i, y))
	}

	for j := y; j >= 0; j-- {
		invalid = append(invalid, square.get(x, j))
	}

	for i := 0; i < size; i++ {
		isPossibility := true

		for j := 0; j < len(invalid); j++ {
			if invalid[j] == i {
				isPossibility = false
				break
			}
		}

		if isPossibility {
			possibilities = append(possibilities, i)
		}
	}

	return possibilities
}

func (square *LatinSquare) set(x, y, value int) {
	square.body[x][y] = value
}

func (square LatinSquare) copy() LatinSquare {
	size := len(square.body)
	body := make([][]int, size)

	for i := range square.body {
		body[i] = make([]int, size)
		copy(body[i], square.body[i])
	}

	return LatinSquare{body}
}

func (square LatinSquare) String() string {
	s := ""
	for _, row := range square.body {
		for _, value := range row {
			s += strconv.Itoa(value)
		}
		s += "\n"
	}

	s += "\n"

	return s
}
