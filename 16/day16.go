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
    route := findRoute(m, start, endpos)
    lib.WritePart1("%d", score(m, route))

    // this should be reusing way more of A* but I'm not editing the lib right now
    toCheck := []node{}
    best := map[node]struct{}{}
    for i:=len(route)-1; i>1; i-- {
        toCheck = append(toCheck, route[i])
        best[route[i]] = struct{}{}
    }
    checked := map[node]struct{}{}
    for len(toCheck) > 0 {
        check := toCheck[0]
        toCheck = toCheck[1:]
        canonical := findRoute(m, check, endpos)
        if len(canonical) == 0 {
            continue
        }
        canonicalsum := score(m, canonical)
        for _, n := range m.Neighbours(check) {
            if _, ok := best[n]; ok {
                continue
            }
            if _, ok := checked[n]; ok {
                continue
            }
            checked[n] = struct{}{}
            newroute := findRoute(m, n, endpos)
            if newroute == nil {
                continue
            }
            if score(m, newroute) + int(m.G(check, n)) > canonicalsum {
                continue
            }
            for _, newnode := range newroute {
                best[newnode] = struct{}{}
                if _, ok := checked[newnode]; ok {
                    continue
                }
                toCheck = append(toCheck, newnode)
            }
        }
    }
    unique := map[coord]struct{}{}
    for k := range best {
        unique[k.pos] = struct{}{}
    }
    lib.WritePart2("%d", len(unique))
}

func score(m maze, route []node) int {
    sum := 0
    curr := route[len(route)-1]
    for i:=len(route)-2; i>0; i-- {
        next := route[i]
        sum += int(m.G(curr, next))
        curr = next
    }
    return sum
}

func findRoute(m maze, start node, endpos coord) []node {
    end := node{pos:endpos}
    route, err := path.FindRouteWithGoalFunc[node](m, start, end, func(c, g node) bool {
        return c.pos == g.pos
    })
    if err != nil {
        return nil
    }
    return route
}
