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

    //generateMermaid(fringe, gates)

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

    // part2 done by hand looking at mermaid output
    // swaps: ggn,grm,jcb,ndw,twr,z10,z32,z39
    lib.WritePart2("ggn,grm,jcb,ndw,twr,z10,z32,z39")
}

// clearly defining an adder, with each input x/y going to 
// a halfadder: carry = AND(x,y); sum = XOR(x,y)
// two halfadders make a full adder, with a carry bit as third input
// c1, s1 = halfadd(x, y); c2, s2 = halfadd(s1, carry); then
// add(x,y,c): carry = OR(c1, c2); sum = s2
// NOTES: all outputs (except last) should be coming from XOR
// ISSUES: z32 and z39 coming from an AND
// z10 is coming from an OR
// XOR[grm] needs to go to out but doesnt (its out z32)
// XOR[twr] needs to go to out but doesnt (its out z39)
// XOR[ggn] should output z10 but doesnt
// AND[ndw] goes to two but should only go to one
// XOR[jcb] is sus
func generateMermaid(inputs map[gate]struct{}, rest []gate) {
    // gvm OR smt -> z10    SWAP    mbv XOR hks -> ggn
    // rmn AND whq -> z32   SWAP    rmn XOR whq -> grm
    // x39 AND y39 -> z39   SWAP    pqv XOR bnv -> twr
    fmt.Println("flowchart TD")
    for g := range inputs {
        if g.x == "x39" && g.op == "AND" {
            g.z = "twr"
        }
        fmt.Printf("%s --> |%s| %s[%s]\n", g.x, g.x, g.z, g.op)
        fmt.Printf("%s --> |%s| %s[%s]\n", g.y, g.y, g.z, g.op)
    }
    for _, g := range rest {
        switch {
        case g.x == "gvm" && g.op == "OR": g.z = "ggn"
        case g.x == "mbv" && g.op == "XOR": g.z = "z10"
        case g.x == "rmn" && g.op == "AND": g.z = "grm"
        case g.x == "rmn" && g.op == "XOR": g.z = "z32"
        case g.x == "pqv" && g.op == "XOR": g.z = "z39"
        }
        fmt.Printf("%s --> |%s| %s[%s]\n", g.x, g.x, g.z, g.op)
        fmt.Printf("%s --> |%s| %s[%s]\n", g.y, g.y, g.z, g.op)
    }
    for i:=0; i<46; i++ {
        fmt.Printf("z%02d --> |z%02d| output\n", i, i)
    }
}
