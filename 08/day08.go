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

func (p coord) sub(q coord) coord {
    return coord{p.x-q.x, p.y-q.y}
}

var m = map[coord]rune{}
var antennae = map[rune][]coord{}

func main() {
    y := 0
    lib.ReadFileByLine(8, func(line string) {
        for x, c := range line {
            m[coord{x,y}] = c
            if c != '.' {
                antennae[c] = append(antennae[c], coord{x,y})
            }
        }
        y++
    })
    antinodes := map[coord]struct{}{}
    for _, pos := range antennae {
        for i:=0; i<len(pos); i++ {
            for j:=i+1; j<len(pos); j++ {
                x, y := pos[i], pos[j]
                c1 := x.add(x.sub(y))
                if _, ok := m[c1]; ok {
                    antinodes[c1] = struct{}{}
                }
                c2 := y.add(y.sub(x))
                if _, ok := m[c2]; ok {
                    antinodes[c2] = struct{}{}
                }
            }
        }
    }
    lib.WritePart1("%d", len(antinodes))
    for _, pos := range antennae {
        for i:=0; i<len(pos); i++ {
            for j:=i+1; j<len(pos); j++ {
                x, y := pos[i], pos[j]
                antinodes[x] = struct{}{}
                antinodes[y] = struct{}{}
                c1 := x.sub(y)
                c2 := y.sub(x)
                for {
                    x = x.add(c1)
                    if _, ok := m[x]; !ok {
                        break
                    }
                    antinodes[x] = struct{}{}
                }
                for {
                    y = y.add(c2)
                    if _, ok := m[y]; !ok {
                        break
                    }
                    antinodes[y] = struct{}{}
                }
            }
        }
    }
    lib.WritePart2("%d", len(antinodes))
}
