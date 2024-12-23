package main


import (
    "slices"
    "strings"

    "github.com/deosjr/adventofcode2024/lib"
)

var graph = map[string][]string{}

func main() {
    neighbours := map[string]struct{}{}
    lib.ReadFileByLine(23, func(line string) {
        x, y := line[:2], line[3:]
        graph[x] = append(graph[x], y)
        graph[y] = append(graph[y], x)
        neighbours[x+y] = struct{}{}
        neighbours[y+x] = struct{}{}
    })
    connected := map[string]struct{}{}
    for k, v := range graph {
        if k[0] != 't' {
            continue
        }
        for i:=0; i<len(v); i++ {
            for j:=i; j<len(v); j++ {
                from, to := v[i], v[j]
                if _, ok := neighbours[from+to]; ok {
                    s := []string{k, from, to}
                    slices.Sort(s)
                    connected[strings.Join(s, "")] = struct{}{}
                }
            }
        }
    }
    lib.WritePart1("%d", len(connected))
    max := 0
    ans := ""
    vertexSet := map[string]struct{}{}
    for k := range graph {
        vertexSet[k] = struct{}{}
    }
    r := map[string]struct{}{}
    x := map[string]struct{}{}
    for _, clique := range BronKerbosch(r, vertexSet, x) {
        if len(clique) > max {
            max = len(clique)
            ans = strings.Join(clique, ",")
        }
    }
    lib.WritePart2("%s", ans)
}

type set map[string]struct{}

func toClique(s set) []string {
    out := []string{}
    for k := range s {
        out = append(out, k)
    }
    slices.Sort(out)
    return out
}

func BronKerbosch(r, p, x set) [][]string {
    if len(p) == 0 && len(x) == 0 {
        return [][]string{ toClique(r) }
    }
    out := [][]string{}
    for v := range p {
        neighbours := n(v)
        setv := map[string]struct{}{ v: struct{}{} }
        rv := union(r, setv)
        np := intersection(p, neighbours)
        nx := intersection(x, neighbours)
        res := BronKerbosch(rv, np, nx)
        out = append(out, res...)
        p = dif(p, setv)
        x = union(x, setv)
    }
    return out
}

func n(v string) set {
    out := map[string]struct{}{}
    for _, n := range graph[v] {
        out[n] = struct{}{}
    }
    return out
}

func union(x, y set) set {
    out := map[string]struct{}{}
    for k := range x {
        out[k] = struct{}{}
    }
    for k := range y {
        out[k] = struct{}{}
    }
    return out
}

func intersection(x, y set) set {
    out := map[string]struct{}{}
    for k := range x {
        if _, ok := y[k]; ok {
            out[k] = struct{}{}
        }
    }
    return out
}

func dif(x, y set) set {
    out := map[string]struct{}{}
    for k := range x {
        if _, ok := y[k]; !ok {
            out[k] = struct{}{}
        }
    }
    return out
}
