package main

import (
    "strings"

    "github.com/deosjr/adventofcode2024/lib"
)

func main() {
    input := lib.ReadFile(25)
    locks := [][5]int{}
    keys := [][5]int{}
    for _, raw := range strings.Split(input, "\n\n") {
        if len(raw) == 0 {
            continue
        }
        if rune(raw[0]) == '#' {
            locks = append(locks, parseLock(raw))
        } else {
            keys = append(keys, parseKey(raw))
        }
    }
    ans := 0
    for _, key := range keys {
        for _, lock := range locks {
            if fits(key, lock) {
                ans++
            }
        }
    }
    lib.WritePart1("%d", ans)
}

func fits(key, lock [5]int) bool {
    for i:=0; i<5; i++ {
        if key[i] + lock[i] > 5 {
            return false
        }
    }
    return true
}

func parseLock(grid string) [5]int {
    split := strings.Split(grid, "\n")
    out := [5]int{}
    for i:=0; i<5; i++ {
        for y:=1; y<7; y++ {
            if split[y][i] == '#' {
                out[i] += 1
            }
        }
    }
    return out
}

func parseKey(grid string) [5]int {
    split := strings.Split(grid, "\n")
    out := [5]int{}
    for i:=0; i<5; i++ {
        for y:=0; y<6; y++ {
            if split[y][i] == '#' {
                out[i] += 1
            }
        }
    }
    return out
}
