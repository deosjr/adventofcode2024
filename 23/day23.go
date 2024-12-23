package main


import (
    "slices"
    "strings"

    "github.com/deosjr/adventofcode2024/lib"
)

func main() {
    graph := map[string][]string{}
    vertices := map[string]struct{}{}
    lib.ReadFileByLine(23, func(line string) {
        x, y := line[:2], line[3:]
        graph[x] = append(graph[x], y)
        graph[y] = append(graph[y], x)
        vertices[x+y] = struct{}{}
        vertices[y+x] = struct{}{}
    })
    connected := map[string]struct{}{}
    for k, v := range graph {
        if k[0] != 't' {
            continue
        }
        for i:=0; i<len(v); i++ {
            for j:=i; j<len(v); j++ {
                from, to := v[i], v[j]
                if _, ok := vertices[from+to]; ok {
                    s := []string{k, from, to}
                    slices.Sort(s)
                    connected[strings.Join(s, "")] = struct{}{}
                }
            }
        }
    }
    lib.WritePart1("%d", len(connected))
}
