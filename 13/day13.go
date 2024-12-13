package main

import (
    "fmt"
    "strings"

    "github.com/deosjr/adventofcode2024/lib"
)

type machine struct {
    xa, ya, xb, yb, xp, yp int64
}

func main() {
    machines := []machine{}
    for _, raw := range strings.Split(lib.ReadFile(13), "\n\n") {
        var xa, ya, xb, yb, xp, yp int64
        fmt.Sscanf(raw, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d", &xa, &ya, &xb, &yb, &xp, &yp)
        machines = append(machines, machine{xa, ya, xb, yb, xp, yp})
    }
    var sum int64
Loop:
    for _, m := range machines {
        for b:=100; b >=0; b-- {
            cost := int64(b)
            x := cost * m.xb
            y := cost * m.yb
            dx := m.xp - x
            dy := m.yp - y
            if dx % m.xa != 0 || dy % m.ya != 0 {
                continue
            }
            d := dx / m.xa
            if d != dy / m.ya {
                continue
            }
            cost += 3*d
            sum += cost
            continue Loop
        }
    }
    lib.WritePart1("%d", sum)
}
