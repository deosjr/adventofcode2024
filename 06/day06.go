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

func (p coord) turnRight() coord {
    return coord{-p.y, p.x}
}

var m = map[coord]bool{}

func main() {
    var start coord
    y := 0
    lib.ReadFileByLine(6, func(line string) {
        for x, c := range line {
            m[coord{x,y}] = c == '#'
            if c == '^' {
                start = coord{x,y}
            }
        }
        y++
    })
    lib.WritePart1("%d", p1(start))
    lib.WritePart2("%d", p2(start))
}

func p1(start coord) int {
    ans, _ := loopOrOut(start)
    return ans
}

func p2(start coord) int {
    sum := 0
    for y:=0; y < 130; y++ {
        for x :=0; x < 130; x++ {
            c := coord{x,y}
            if c == start {
                continue
            }
            if m[c] {
                continue
            }
            m[c] = true
            if _, loop := loopOrOut(start); loop {
                sum++
            }
            m[c] = false
        }
    }
    return sum
}

type state struct {
    pos, heading coord
}

// bool argument is true if we are stuck in a loop
// returns length of path if out of bounds
func loopOrOut(start coord) (int, bool) {
    pos := start
    heading := coord{0, -1}
    mem := map[state]struct{}{ {pos, heading}: struct{}{} }
    visited := map[coord]struct{}{ start: struct{}{} }
    for {
        ahead := pos.add(heading)
        blocked, inArea := m[ahead]
        if !inArea {
            break
        }
        if blocked {
            heading = heading.turnRight()
        } else {
            pos = ahead
            visited[pos] = struct{}{}
        }
        newstate := state{pos, heading}
        if _, ok := mem[newstate]; ok {
            return 0, true
        }
        mem[newstate] = struct{}{}
    }
    return len(visited), false
}
