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
    ans1, ans2 := 0, 0
    for _, t := range trailheads {
        nines := reachable(t, 0)
        dedup := map[coord]struct{}{}
        for _, k := range nines {
            dedup[k] = struct{}{}
        }
        ans1 += len(dedup)
        ans2 += len(nines)
    }
    lib.WritePart1("%d", ans1)
    lib.WritePart2("%d", ans2)
}

var mem = map[coord][]coord{}

func reachable(c coord, value int) []coord {
    if v, ok := mem[c]; ok {
        return v
    }
    list := []coord{}
    for _, n := range []coord{
        c.add(coord{-1,0}), c.add(coord{1,0}), c.add(coord{0,-1}), c.add(coord{0,1}),
    }{
        if !exists(n, value+1) {
            continue
        }
        if value == 8 {
            list = append(list, n)
        } else {
            list = append(list, reachable(n, value+1)...)
        }
    }
    mem[c] = list
    return list
}

func exists(c coord, value int) bool {
    v, ok := m[c]
    return ok && v == value
}
