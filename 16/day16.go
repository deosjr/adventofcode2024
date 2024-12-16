package main

import (
    "math"

    "github.com/deosjr/Pathfinding/path"
    "github.com/deosjr/adventofcode2024/lib"
)

type coord struct {
    x, y int
}

func (p coord) add(q coord) coord {
    return coord{p.x+q.x, p.y+q.y}
}

func (p coord) turnLeft() coord {
    return coord{p.y, -p.x}
}

func (p coord) turnRight() coord {
    return coord{-p.y, p.x}
}

type maze struct {
    grid map[coord]bool
}

type node struct {
    pos, heading coord
}

func (m maze) Neighbours(p node) []node {
    neighbours := []node{
        {pos: p.pos, heading: p.heading.turnLeft()},
        {pos: p.pos, heading: p.heading.turnRight()},
    }
    next := p.pos.add(p.heading)
    if !m.grid[next] {
        neighbours = append(neighbours, node{pos: next, heading: p.heading})
    }
    return neighbours
}

func (m maze) G(p, q node) float64 {
    if p.pos == q.pos {
        return 1000
    }
    return 1
}

func (m maze) H(p, q node) float64 {
    dx := float64(q.pos.x - p.pos.x)
	dy := float64(q.pos.y - p.pos.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
    m := maze{grid: map[coord]bool{}}
    var startpos, endpos coord
    y := 0
    lib.ReadFileByLine(16, func(line string) {
        for x, c := range line {
            m.grid[coord{x,y}] = c == '#'
            if c == 'S' {
                startpos = coord{x,y}
            }
            if c == 'E' {
                endpos = coord{x,y}
            }
        }
        y++
    })
    start := node{pos:startpos, heading: coord{1,0}}
    end := node{pos:endpos}
    route, err := path.FindRouteWithGoalFunc[node](m, start, end, func(c, g node) bool {
        return c.pos == g.pos
    })
    if err != nil {
        panic(err)
    }
    sum := 0
    curr := route[len(route)-1]
    for i:=len(route)-2; i>0; i-- {
        next := route[i]
        sum += int(m.G(curr, next))
        curr = next
    }
    lib.WritePart1("%d", sum)
}
