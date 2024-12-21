package main

import (
    "github.com/deosjr/adventofcode2024/lib"
)

func main() {
    codes := []string{}
    lib.ReadFileByLine(21, func(line string) {
        codes = append(codes, line)
    })
    handwritten := [][]string{
        {"^<<A","^A",">vvA",">A"}, // 140A
        {"^<<A","^^>A","vvvA",">A"}, // 180A
        {"^<<A","^^A","v>>A","vvA"}, // 176A
        {"<^^^A","vvvA","^^A","vv>A"}, // 805A
        {"^^A","vA","<^^A","vvv>A"}, // 638A
    }

    p1 := 0
    for i, c := range codes {
        numpad := handwritten[i]
        shortest := rewriteSlice(numpad, 2)
        p1 += shortest * int(lib.MustParseInt(c[:3]))
    }
    lib.WritePart1("%d", p1)
    p2 := 0
    for i, c := range codes {
        numpad := handwritten[i]
        shortest := rewriteSlice(numpad, 25)
        p2 += shortest * int(lib.MustParseInt(c[:3]))
    }
    lib.WritePart2("%d", p2)
}

type input struct {
    s string
    iterations int
}

var mem = map[input]int{}

func rewriteSlice(in []string, i int) int {
    sum := 0
    for _, s := range in {
        sum += rewriteMany(s, i)
    }
    return sum
}

func rewriteMany(in string, i int) int {
    if i == 0 {
        return len(in)
    }
    if v, ok := mem[input{in, i}]; ok {
        return v
    }
    out := []string{}
    prev := 'A'
    for _, c := range in {
        cur := rune(c)
        out = append(out, rewriteOnce(prev, cur))
        prev = cur
    }
    ans := rewriteSlice(out, i-1)
    mem[input{in, i}] = ans
    return ans
}

func rewriteOnce(prev, cur rune) string {
    s := ""
    switch {
    case prev == 'A' && cur == '^': s+= "<"
    case prev == 'A' && cur == '>': s+= "v"
    case prev == 'A' && cur == 'v': s+= "<v"
    case prev == 'A' && cur == '<': s+= "v<<"
    case prev == '^' && cur == 'A': s+= ">"
    case prev == '^' && cur == '>': s+= "v>"    //
    case prev == '^' && cur == 'v': s+= "v"
    case prev == '^' && cur == '<': s+= "v<"
    case prev == 'v' && cur == '^': s+= "^"
    case prev == 'v' && cur == '>': s+= ">"
    case prev == 'v' && cur == 'A': s+= "^>"    //
    case prev == 'v' && cur == '<': s+= "<"
    case prev == '<' && cur == '^': s+= ">^"
    case prev == '<' && cur == '>': s+= ">>"
    case prev == '<' && cur == 'v': s+= ">"
    case prev == '<' && cur == 'A': s+= ">>^"
    case prev == '>' && cur == '^': s+= "<^"    //
    case prev == '>' && cur == 'A': s+= "^"
    case prev == '>' && cur == 'v': s+= "<"
    case prev == '>' && cur == '<': s+= "<<"
    }
    s += "A"
    return s
}
