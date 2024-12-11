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
    lib.WritePart1("%v", p1(input))
    lib.WritePart2("%v", p2(input))
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

type in struct {
    stone int64
    n int
}

var mem = map[in]int64{}

func blinkRec(stone int64, n int) int64 {
    if v, ok := mem[in{stone, n}]; ok {
        return v
    }
    rec := blink(stone)
    if n == 1 {
        mem[in{stone, n}] = int64(len(rec))
        return int64(len(rec))
    }
    var sum int64
    for _, s := range rec {
        sum += blinkRec(s, n-1)
    }
    mem[in{stone, n}] = sum
    return sum
}

func p1(list []int64) int64 {
    var sum int64
    for _, s := range list {
        sum += blinkRec(s, 25)
    }
    return sum
}

func p2(list []int64) int64 {
    var sum int64
    for _, s := range list {
        sum += blinkRec(s, 75)
    }
    return sum
}
