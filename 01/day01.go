package main

import (
    "slices"
    "strings"

	"github.com/deosjr/adventofcode2024/lib"
)

func main() {
    left := []int64{}
    right := []int64{}
    lib.ReadFileByLine(1, func(line string) {
        split := strings.Split(line, "   ")
        left = append(left, lib.MustParseInt(split[0]))
        right = append(right, lib.MustParseInt(split[1]))
    })
    slices.Sort(left)
    slices.Sort(right)
    var p1 int64 = 0
    for i, x := range left {
        y := right[i]
        div := x-y
        if div < 0 {
            div += -2 * div
        }
        p1 += div
    }
    lib.WritePart1("%d", p1)
    m := map[int64]int64{}
    for _, x := range right {
        m[x] += 1
    }
    var p2 int64 = 0
    for _, x := range left {
        p2 += x * m[x]
    }
    lib.WritePart2("%d", p2)
}
