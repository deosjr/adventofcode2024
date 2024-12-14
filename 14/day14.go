package main

import (
    "fmt"
    "math"

    "github.com/deosjr/adventofcode2024/lib"
)

type coord struct {
    x, y int
}

type robot struct {
    px, py, vx, vy int
}

var width = 101
var height = 103

func main() {
    robots := []robot{}
    lib.ReadFileByLine(14, func(line string) {
        var px, py, vx, vy int
        fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
        robots = append(robots, robot{px, py, vx, vy})
    })
    var ulhc, urhc, llhc, lrhc int64
    for k, v := range run(robots, 100) {
        switch {
        case k.x < width/2 && k.y < height/2: ulhc += v
        case k.x < width/2 && k.y > height/2: llhc += v
        case k.x > width/2 && k.y < height/2: urhc += v
        case k.x > width/2 && k.y > height/2: lrhc += v
        }
    }
    lib.WritePart1("%d", ulhc*urhc*llhc*lrhc)
    mid := coord{width/2, height/2}
    minScore := math.MaxFloat64
    minSeconds := -1
    for i:=1; i<10000; i++ {
        m := run(robots, i)
        var sum float64
        for k, v := range m {
            dx := float64(k.x - mid.x)
            dy := float64(k.y - mid.y)
            sum += float64(v) * math.Sqrt( dx*dx + dy*dy  )
        }
        if sum < minScore {
            minScore = sum
            minSeconds = i
            fmt.Println(minSeconds, minScore)
        }
    }
}

func plusmod(n, m int) int {
    x := n % m
    if x < 0 {
        x += m
    }
    return x
}

func run(robots []robot, seconds int) map[coord]int64 {
    m := map[coord]int64{}
    for _, r := range robots {
        nx := plusmod(r.px + seconds * r.vx, width)
        ny := plusmod(r.py + seconds * r.vy, height)
        m[coord{nx, ny}] += 1
    }
    return m
}
