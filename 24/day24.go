package main

import (
    "fmt"
    "strings" 

    "github.com/deosjr/adventofcode2024/lib"
)

type gate struct {
    x, op, y, z string
}

func (g gate) ready(m map[string]bool) bool {
    _, xok := m[g.x]
    _, yok := m[g.y]
    _, zok := m[g.z]
    return xok && yok && !zok
}

func f(op string, x, y bool) bool {
    switch op {
    case "AND": return x && y
    case "OR": return x || y
    case "XOR": return x != y
    }
    panic(op)
}

func main() {
    input := lib.ReadFile(24)
    split := strings.Split(input, "\n\n")
    m := map[string]bool{}
    gates := []gate{}
    fringe := map[gate]struct{}{}
    for _, s := range strings.Split(split[0], "\n") {
        ss := strings.Split(s, ": ")
        m[ss[0]] = ss[1][0] == '1'
    }
    for _, s := range strings.Split(strings.TrimSpace(split[1]), "\n") {
        ss := strings.Split(s, " ")
        g := gate{ss[0], ss[1], ss[2], ss[4]}
        if g.ready(m) {
            fringe[g] = struct{}{}
            continue
        }
        gates = append(gates, g)
    }
    for len(fringe) > 0 {
        var g gate
        for k := range fringe {
            g = k
            break
        }
        delete(fringe, g)
        m[g.z] = f(g.op, m[g.x], m[g.y])
        for _, ng := range gates {
            if !ng.ready(m) {
                continue
            }
            fringe[ng] = struct{}{}
        }
    }
    // z45 is highest
    var ans int64 = 0
    for i:=0; i<46; i++ {
        z := fmt.Sprintf("z%02d", i)
        var b int64 = 0
        if m[z] { b = 1 }
        b = b << i
        ans |= b
    }
    lib.WritePart1("%d", ans)
}
