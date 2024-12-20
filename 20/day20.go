package main

import (
    "math"

    "github.com/deosjr/adventofcode2024/lib"
)

type coord struct {
    x, y int
}

func (p coord) add(q coord) coord {
    return coord{p.x+q.x, p.y+q.y}
}

func neighbours(c coord) []coord {
    return []coord{
        c.add(coord{1,0}), c.add(coord{-1,0}), c.add(coord{0,1}), c.add(coord{0,-1}),
    }
}

func main() {
    walls := map[coord][]coord{}    // wall coord -> empty neighbours
    var start, end coord
    y := 0
    lib.ReadFileByLine(20, func(line string) {
        for x, c := range line {
            if c == 'S' {
                start = coord{x,y}
            }
            if c == 'E' {
                end = coord{x,y}
            }
            if c != '#' {
                continue
            }
            walls[coord{x,y}] = nil
        }
        y++
    })
    fill := map[coord]int{}         // empty -> distance from end
    fill[end] = 0
    pos := end
    for _, n := range neighbours(end) {
        if _, ok := walls[n]; ok {
            walls[n] = append(walls[n], end)
        }
    }
    for pos != start {
        for _, n := range neighbours(pos) {
            if _, ok := walls[n]; ok {
                continue
            }
            if _, ok := fill[n]; ok {
                continue
            }
            fill[n] = len(fill)
            pos = n
            for _, n := range neighbours(pos) {
                if _, ok := walls[n]; ok {
                    walls[n] = append(walls[n], pos)
                }
            }
            break
        }
    }
    p1 := 0
    cheats := map[int]int{}         // seconds -> number of cheats
    for _, v := range walls {
        if len(v) < 2 {
            continue
        }
        for i:=0; i<len(v); i++ {
            for j:=i+1; j<len(v); j++ {
                score1, score2 := fill[v[i]], fill[v[j]]
                abs := int(math.Abs(float64(score1)-float64(score2))) - 2
                if abs == 0 {
                    continue
                }
                if abs >= 100 {
                    p1++
                }
                cheats[abs] += 1
            }
        }
    }
    lib.WritePart1("%d", p1)
}
