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

var m = map[coord]int{}

func main() {
    trailheads := []coord{}
    y := 0
    lib.ReadFileByLine(10, func(line string) {
        for x, c := range line {
            m[coord{x,y}] = int(c) - 48
            if m[coord{x,y}] == 0 {
                trailheads = append(trailheads, coord{x,y})
            }
        }
        y++
    })
    ans1 := 0
    for _, t := range trailheads {
        ans1 += len(reachable(t, 0))
    }
    lib.WritePart1("%d", ans1)
}

var mem = map[coord][]coord{}

func reachable(c coord, value int) []coord {
    if v, ok := mem[c]; ok {
        return v
    }
    next := map[coord]struct{}{}
    for _, n := range []coord{
        c.add(coord{-1,0}), c.add(coord{1,0}), c.add(coord{0,-1}), c.add(coord{0,1}),
    }{
        if !exists(n, value+1) {
            continue
        }
        if value == 8 {
            next[n] = struct{}{}
        } else {
            for _, r := range reachable(n, value+1) {
                next[r] = struct{}{}
            }
        }
    }
    list := []coord{}
    for k := range next {
        list = append(list, k)
    }
    mem[c] = list
    return list
}

func exists(c coord, value int) bool {
    v, ok := m[c]
    return ok && v == value
}
