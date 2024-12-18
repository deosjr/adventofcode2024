package main

import (
    "fmt"
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

type maze struct {
    grid map[coord]bool
}

func (m maze) Neighbours(p coord) []coord {
    neighbours := []coord{}
    for _, n := range []coord{
        p.add(coord{1,0}),
        p.add(coord{-1,0}),
        p.add(coord{0,1}),
        p.add(coord{0,-1}),
    } {
        if n.x < minX || n.x > maxX || n.y < minY || n.y > maxY {
            continue
        }
        if !m.grid[n] {
            neighbours = append(neighbours, n)
        }
    }
    return neighbours
}

func (m maze) G(p, q coord) float64 {
    return 1
}

func (m maze) H(p, q coord) float64 {
    dx := float64(q.x - p.x)
	dy := float64(q.y - p.y)
	return math.Sqrt(dx*dx + dy*dy)
}

var (
    minX, minY int = 0, 0
    maxX, maxY int = 70, 70
)

func main() {
    coords := []coord{}
    lib.ReadFileByLine(18, func(line string) {
        var x, y int
        fmt.Sscanf(line, "%d,%d", &x, &y)
        coords = append(coords, coord{x, y})
    })
    m := maze{grid: map[coord]bool{}}
    for i:=0; i<1024; i++ {
        m.grid[coords[i]] = true    
    }
    start := coord{minX, minY}
    end := coord{maxX, maxY}
    route, err := path.FindRoute[coord](m, start, end)
    if err != nil {
        panic(err)
    }
    lib.WritePart1("%d", len(route)-2)

    for i:=1024; i<len(coords); i++ {
        m.grid[coords[i]] = true
        _, err := path.FindRoute[coord](m, start, end)
        if err != nil {
            lib.WritePart2("%d,%d", coords[i].x, coords[i].y)
            break
        }
    }
}
