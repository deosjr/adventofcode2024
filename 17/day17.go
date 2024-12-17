package main

import (
    "fmt"
    "strings"

    "github.com/deosjr/adventofcode2024/lib"
)

type computer struct {
    a, b, c int64
    program []uint8
    ip int
    res []uint8
}

func (c *computer) run() []uint8 {
    for c.ip < len(c.program) {
        ins := c.program[c.ip]
        switch ins {
        case 0: c.adv()
        case 1: c.bxl()
        case 2: c.bst()
        case 3: c.jnz()
        case 4: c.bxc()
        case 5: c.out()
        case 6: c.bdv()
        case 7: c.cdv()
        }
        if ins != 3 || c.a == 0 {
            c.ip += 2
        }
    }
    return c.res
}

func (c *computer) combo(x uint8) int64 {
    switch {
    case x < 4: return int64(x)
    case x == 4: return c.a
    case x == 5: return c.b
    case x == 6: return c.c
    }
    panic("incorrect combo operand")
}

func (c *computer) adv() {
    denom := c.combo(c.program[c.ip+1])
    c.a = c.a >> denom
}

func (c *computer) bxl() {
    c.b = c.b ^ int64(c.program[c.ip+1])
}

func (c *computer) bst() {
    c.b = c.combo(c.program[c.ip+1]) % 8
}

func (c *computer) jnz() {
    if c.a == 0 {
        return
    }
    c.ip = int(c.program[c.ip+1])
}

func (c *computer) bxc() {
    c.b = c.b ^ c.c
}

func (c *computer) out() {
    c.res = append(c.res, uint8(c.combo(c.program[c.ip+1]) % 8))
}

func (c *computer) bdv() {
    denom := c.combo(c.program[c.ip+1])
    c.b = c.a >> denom
}

func (c *computer) cdv() {
    denom := c.combo(c.program[c.ip+1])
    c.c = c.a >> denom
}

func main() {
    input := lib.ReadFile(17)
    split := strings.Split(input, "\n")
    var a, b, c int64
    fmt.Sscanf(split[0], "Register A: %d", &a)
    fmt.Sscanf(split[1], "Register B: %d", &b)
    fmt.Sscanf(split[2], "Register C: %d", &c)
    program := []uint8{}
    for _, c := range strings.Split(strings.Split(split[4], ": ")[1], ",") {
        program = append(program, c[0] - 48)
    }
    comp := &computer{a:a, b:b, c:c, program:program}
    p1 := comp.run()
    s := make([]string, len(p1))
    for i, n := range p1 {
        s[i] = fmt.Sprintf("%d", n)
    }
    lib.WritePart1("%s", strings.Join(s, ",")) 

    candidates := [][]uint8{{}}
    for len(candidates) > 0 {
        newc := candidates[0]
        if len(newc) == 16 {
            break
        }
        candidates = candidates[1:]
        var sofar int64
        for i, n := range newc {
            sofar += int64(n) << ((len(newc)-i)*3)
        }
        for i:=0; i<8; i++ {
            newa := sofar + int64(i)
            comp := &computer{a:newa, b:b, c:c, program:program}
            out := comp.run()
            if len(out) != len(newc)+1 {
                continue
            }
            if out[0] == program[len(program)-(len(newc)+1)] {
                newcc := make([]uint8, len(newc)+1)
                for i, n := range newc {
                    newcc[i] = n
                }
                newcc[len(newc)] = uint8(i)
                candidates = append(candidates, newcc)
            }
        }
    }
    var sum int64
    for i, n := range candidates[0] {
        sum += int64(n) << ((16-i-1)*3)
    }
    lib.WritePart2("%d", sum)
}
