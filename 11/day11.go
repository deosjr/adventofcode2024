package main

import (
    "fmt"
    "strings"

	"github.com/deosjr/adventofcode2024/lib"
)

func main() {
    raw := strings.TrimSpace(lib.ReadFile(11))
    split := strings.Split(raw, " ")
    input := make([]int64, len(split))
    for i, s := range split {
        input[i] = lib.MustParseInt(s)
    }
    ans1 := len(p1(input))
    lib.WritePart1("%v", ans1)
}

func blink(stone int64) []int64 {
    if stone == 0 {
        return []int64{1}
    }
    list := fmt.Sprintf("%d", stone)
    if len(list) % 2 == 0 {
        pivot := len(list) / 2
        left := lib.MustParseInt(list[:pivot])
        right := lib.MustParseInt(list[pivot:])
        return []int64{left, right}
   }
   return []int64{stone * 2024}
}

func p1(list []int64) []int64 {
    newlist := []int64{}
    for i:=0; i<25; i++ {
        for _, s := range list {
            newlist = append(newlist, blink(s)...)
        }
        list = newlist
        newlist = []int64{}
    }
    return list
}
