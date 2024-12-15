package main

import (
    "strings"

    "github.com/deosjr/adventofcode2024/lib"
)

type coord struct {
    x, y int
}

func (p coord) add(q coord) coord {
    return coord{p.x+q.x, p.y+q.y}
}

func main() {
    input := strings.Split(lib.ReadFile(15), "\n\n")
    grid, instructions := input[0], input[1]
    m := map[coord]rune{}
    var pos coord
    for y, line := range strings.Split(grid, "\n") {
        for x, c := range line {
            m[coord{x,y}] = c
            if c == '@' {
                pos = coord{x,y}
            }
        }
    }
    for _, ins := range strings.Replace(instructions, "\n", "", -1) {
        var heading coord
        switch ins {
        case '^':
            heading = coord{0,-1}
        case 'v':
            heading = coord{0,1}
        case '<':
            heading = coord{-1,0}
        case '>':
            heading = coord{1,0}
        }
        next := pos.add(heading)
        switch m[next] {
        case '.':
            m[pos] = '.'
            m[next] = '@'
            pos = next
            continue
        case '#':
            continue
        case 'O':
            newnext := next
            for {
                newnext = newnext.add(heading)
                if m[newnext] == '#' {
                    break
                }
                if m[newnext] == '.' {
                    m[pos] = '.'
                    m[next] = '@'
                    m[newnext] = 'O'
                    pos = next
                    break
                }
            }
        }
    }
    sum := 0
    for k, v := range m {
        if v != 'O' {
            continue
        }
        sum += k.x + 100*k.y
    }
    lib.WritePart1("%v", sum)
}
