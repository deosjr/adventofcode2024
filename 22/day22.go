package main

import (
    "github.com/deosjr/adventofcode2024/lib"
)

func main() {
    secrets := []int64{}
    lib.ReadFileByLine(22, func(line string) {
        secrets = append(secrets, lib.MustParseInt(line))
    })
    //secrets = []int64{1, 10, 100, 2024}
    var p1 int64
    for _, s := range secrets {
        for i:=0; i<2000; i++ {
            s = genSecret(s)
        }
        p1 += s
    }
    lib.WritePart1("%d", p1)
}

func genSecret(seed int64) int64 {
    seed = ((seed * 64) ^ seed) % 16777216
    seed = ((seed / 32) ^ seed) % 16777216
    seed = ((seed * 2048) ^ seed) % 16777216
    return seed
}
