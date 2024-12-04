package main

import (
	"github.com/deosjr/adventofcode2024/lib"
)

type coord struct { x, y int }

func (p coord) add(q coord) coord {
    return coord{p.x+q.x, p.y+q.y}
}

var m = map[coord]rune{}

func main() {
    xcoords := []coord{}
    acoords := []coord{}
    y := 0
    lib.ReadFileByLine(4, func(line string) {
        for x, c := range line {
            m[coord{x,y}] = c
            if c == 'X' {
                xcoords = append(xcoords, coord{x,y})
            }
            if c == 'A' {
                acoords = append(acoords, coord{x,y})
            }
        }
        y++
    })
    p1 := 0
    for _, xc := range xcoords {
        p1 += xmas(xc)
    }
    lib.WritePart1("%d", p1)
    p2 := 0
    for _, ac := range acoords {
        if x_mas(ac) { p2++ }
    }
    lib.WritePart2("%d", p2)
}

func xmas(start coord) int {
    sum := 0
    for _, d := range [][]coord{
        {{1,0}, {2,0}, {3,0}},
        {{-1,0}, {-2,0}, {-3,0}},
        {{0,1}, {0,2}, {0,3}},
        {{0,-1}, {0,-2}, {0,-3}},
        {{1,1}, {2,2}, {3,3}},
        {{1,-1}, {2,-2}, {3,-3}},
        {{-1,1}, {-2,2}, {-3,3}},
        {{-1,-1}, {-2,-2}, {-3,-3}},
    }{
        if checkXmas(start, d[0], d[1], d[2]) {
            sum++
        }
    }
    return sum
}

func checkXmas(start, one, two, three coord) bool {
    return m[start.add(one)] == 'M' && m[start.add(two)] == 'A' && m[start.add(three)] == 'S'
}

func x_mas(start coord) bool {
    sum := 0
    for _, d := range [][]coord{
        {{1,1}, {-1,-1}},
        {{1,-1}, {-1,1}},
        {{-1,1}, {1,-1}},
        {{-1,-1}, {1,1}},
    }{
        if checkMas(start, d[0], d[1]) {
            sum++
        }
    }
    return sum == 2
}

func checkMas(start, one, two coord) bool {
    return m[start.add(one)] == 'M' && m[start.add(two)] == 'S'
}
