package main

import (
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

var (
    m = map[coord]rune{}
    seen = map[coord]*region{}
    regions []*region
)

type region struct {
    plant rune
    coords []coord
    area int
    perimeter int
}

func main() {
    y := 0
    lib.ReadFileByLine(12, func(line string) {
        for x, c := range line {
            m[coord{x,y}] = c
        }
        y++
    })
    for k, v := range m {
        createRegion(k, v)
    }
    var p1, p2 int
    for _, r := range regions {
        p1 += r.area * r.perimeter
        p2 += r.area * r.numberOfSides()
    }
    lib.WritePart1("%d", p1)
    lib.WritePart2("%d", p2)
}

func createRegion(c coord, k rune) {
    if _, ok := seen[c]; ok {
        return
    }
    r := &region{
        plant: k,
        coords: []coord{c},
        area: 1,
    }
    seen[c] = r
    r.floodFill(c)
    regions = append(regions, r)
}

func (r *region) same(c coord) bool {
    v, ok := m[c]
    return ok && v == r.plant
}

func (r *region) floodFill(c coord) {
    for _, n := range []coord{
        c.add(coord{-1,0}),
        c.add(coord{1,0}),
        c.add(coord{0,-1}),
        c.add(coord{0,1}),
    }{
        if !r.same(n) {
            r.perimeter += 1
            continue
        }
        if _, ok := seen[n]; ok {
            continue
        }
        seen[n] = r
        r.coords = append(r.coords, n)
        r.area += 1
        r.floodFill(n)
    }
}

var mem = map[coord]struct{}{}

func (r *region) numberOfSides() int {
    sum := 0
    for _, start := range r.coords {
        for {
            oneRight := start.add(coord{1,0})
            if !r.same(oneRight) {
                break
            }
            start = oneRight
        }
        sum += r.walk(start, start, coord{0,-1}, coord{1,0}, 0)
    }
    return sum
}

func (r *region) walk(start, pos, heading, right coord, acc int) int {
    if start == pos && heading.x == 0 && heading.y == -1 && acc > 0 {
        return acc
    }
    if right.x == 1 && right.y == 0 {
        if _, ok := mem[pos]; ok {
            return 0
        }
        mem[pos] = struct{}{}
    }
    next := pos.add(heading)
    if !r.same(next) {
        return r.walk(start, pos, heading.turnLeft(), right.turnLeft(), acc+1)
    }
    nextRight := next.add(right)
    if r.same(nextRight) {
        return r.walk(start, next, heading.turnRight(), right.turnRight(), acc+1)
    }
    return r.walk(start, next, heading, right, acc)
}
