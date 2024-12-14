package main

import (
    "fmt"

    "github.com/deosjr/adventofcode2024/lib"
)

type coord struct {
    x, y int
}

type robot struct {
    px, py, vx, vy int
}

func main() {
    robots := []robot{}
    lib.ReadFileByLine(14, func(line string) {
        var px, py, vx, vy int
        fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
        robots = append(robots, robot{px, py, vx, vy})
    })
    width, height := 101, 103
    m := map[coord]int64{}
    for _, r := range robots {
        nx := plusmod(r.px + 100 * r.vx, width)
        ny := plusmod(r.py + 100 * r.vy, height)
        m[coord{nx, ny}] += 1
    }
    var ulhc, urhc, llhc, lrhc int64
    for k, v := range m {
        switch {
        case k.x < width/2 && k.y < height/2: ulhc += v
        case k.x < width/2 && k.y > height/2: llhc += v
        case k.x > width/2 && k.y < height/2: urhc += v
        case k.x > width/2 && k.y > height/2: lrhc += v
        }
    }
    lib.WritePart1("%d", ulhc*urhc*llhc*lrhc)
}

func plusmod(n, m int) int {
    x := n % m
    if x < 0 {
        x += m
    }
    return x
}
