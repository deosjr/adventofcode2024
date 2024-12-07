package main

import (
    "fmt"
    "strings"

    "github.com/deosjr/adventofcode2024/lib"
)

func main() {
    equations := [][]int64{}
    lib.ReadFileByLine(7, func(line string) {
        split := strings.Split(strings.Replace(line, ":", "", -1), " ")
        nums := make([]int64, len(split))
        for i, s := range split {
            nums[i] = lib.MustParseInt(s)
        }
        equations = append(equations, nums)
    })
    var p1 int64 = 0
    for _, eq := range equations {
        test, first, rest := eq[0], eq[1], eq[2:]
        if calibrate(test, rest, first) {
            p1 += test
        }
    }
    lib.WritePart1("%d", p1)
    var p2 int64 = 0
    for _, eq := range equations {
        test, first, rest := eq[0], eq[1], eq[2:]
        if calibrate2(test, rest, first) {
            p2 += test
        }
    }
    lib.WritePart2("%d", p2)
}

func calibrate(test int64, list []int64, acc int64) bool {
    if len(list) == 0 {
        return test == acc
    }
    if acc > test {
        return false
    }
    next, rest := list[0], list[1:]
    plus := acc + next
    times := acc * next
    return calibrate(test, rest, plus) || calibrate(test, rest, times)
}

func calibrate2(test int64, list []int64, acc int64) bool {
    if len(list) == 0 {
        return test == acc
    }
    if acc > test {
        return false
    }
    next, rest := list[0], list[1:]
    plus := acc + next
    times := acc * next
    concat := lib.MustParseInt(fmt.Sprintf("%d%d", acc, next))
    return calibrate2(test, rest, plus) || calibrate2(test, rest, times) || calibrate2(test, rest, concat)
}
