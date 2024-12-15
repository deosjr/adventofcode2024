package main

import (
    "slices"
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
    grid, instrs := input[0], input[1]
    m := map[coord]rune{}
    m2 := map[coord]rune{}
    var start coord
    for y, line := range strings.Split(grid, "\n") {
        for x, c := range line {
            switch c {
            case '.':
                m[coord{x,y}] = c
                m2[coord{2*x,y}] = c
                m2[coord{2*x+1,y}] = c
            case '@':
                m[coord{x,y}] = c
                m2[coord{2*x,y}] = c
                m2[coord{2*x+1,y}] = '.'
                start = coord{x,y}
            case '#':
                m[coord{x,y}] = c
                m2[coord{2*x,y}] = c
                m2[coord{2*x+1,y}] = c
            case 'O':
                m[coord{x,y}] = c
                m2[coord{2*x,y}] = '['
                m2[coord{2*x+1,y}] = ']'
            }
        }
    }
    instructions := []coord{}
    for _, ins := range strings.Replace(instrs, "\n", "", -1) {
        switch ins {
        case '^':
            instructions = append(instructions, coord{0,-1})
        case 'v':
            instructions = append(instructions, coord{0,1})
        case '<':
            instructions = append(instructions, coord{-1,0})
        case '>':
            instructions = append(instructions, coord{1,0})
        }
    }
    lib.WritePart1("%d", p1(m, instructions, start))
    start.x = start.x*2
    lib.WritePart2("%d", p2(m2, instructions, start))
}

func score(m map[coord]rune) int {
    sum := 0
    for k, v := range m {
        if v != 'O' && v != '[' {
            continue
        }
        sum += k.x + 100*k.y
    }
    return sum
}

func p1(m map[coord]rune, instructions []coord, start coord) int {
    pos := start
    for _, heading := range instructions {
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
    return score(m)
}

func p2(m map[coord]rune, instructions []coord, start coord) int {
    pos := start
    for _, heading := range instructions {
        next := pos.add(heading)
        switch m[next] {
        case '.':
            m[pos] = '.'
            m[next] = '@'
            pos = next
            continue
        case '#':
            continue
        }
        if heading.x == -1 || heading.x == 1 {
            pos = moveHorizontal(m, pos, heading)
        } else {
            pos = moveVertical(m, pos, heading)
        }
    }
    return score(m)
}

func moveHorizontal(m map[coord]rune, pos, heading coord) coord {
    next := pos
    numRocks := 0
    for {
        next = next.add(heading)
        if m[next] == '#' {
            return pos
        }
        if m[next] == '[' {
            numRocks++
        }
        if m[next] == '.' {
            break
        }
    }
    m[pos] = '.'
    m[pos.add(heading)] = '@'
    newnext := pos.add(heading)
    for i:=0; i<numRocks; i++ {
        newnext = newnext.add(heading)
        if heading.x == -1 {
            m[newnext] = ']'
        } else {
            m[newnext] = '['
        }
        newnext = newnext.add(heading)
        if heading.x == -1 {
            m[newnext] = '['
        } else {
            m[newnext] = ']'
        }
    }
    return pos.add(heading)
}

func moveVertical(m map[coord]rune, pos, heading coord) coord {
    rocks := map[coord]struct{}{}
    toCheck := []coord{pos.add(heading)}
    for len(toCheck) > 0 {
        check, pop := toCheck[0], toCheck[1:]    
        toCheck = pop
        switch m[check] {
        case '#':
            return pos
        case '.':
            continue
        case '[':
            if _, ok := rocks[check]; ok {
                continue
            }
            rocks[check] = struct{}{}
            next := check.add(heading)
            toCheck = append(toCheck, next, next.add(coord{1,0}))
        case ']':
            left := check.add(coord{-1,0})
            if _, ok := rocks[left]; ok {
                continue
            }
            rocks[left] = struct{}{}
            toCheck = append(toCheck, left.add(heading), check.add(heading))
        }
    }
    toUpdate := []coord{}
    for r := range rocks {
        toUpdate = append(toUpdate, r)
    }
    if heading.y == 1 {
        slices.SortFunc(toUpdate, func(p, q coord) int {
            switch {
            case p.y > q.y:
                return -1
            case p.y < q.y:
                return 1
            }
            return 0
        })
    } else {
        slices.SortFunc(toUpdate, func(p, q coord) int {
            switch {
            case p.y < q.y:
                return -1
            case p.y > q.y:
                return 1
            }
            return 0
        })
    }
    for _, r := range toUpdate {
        right := r.add(coord{1,0})
        m[r] = '.'
        m[right] = '.'
        m[r.add(heading)] = '['
        m[right.add(heading)] = ']'
    }
    m[pos] = '.'
    m[pos.add(heading)] = '@'
    return pos.add(heading)
}
