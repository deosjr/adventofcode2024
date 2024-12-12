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
    p1 := 0
    for _, r := range regions {
        p1 += r.area * r.perimeter
    }
    lib.WritePart1("%d", p1)
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

func (r *region) floodFill(c coord) {
    for _, n := range []coord{
        c.add(coord{-1,0}),
        c.add(coord{1,0}),
        c.add(coord{0,-1}),
        c.add(coord{0,1}),
    }{
        v, ok := m[n]
        if !ok || v != r.plant {
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
