package main

import (
    "github.com/deosjr/adventofcode2024/lib"
)

type buyer struct {
    secret int64
    prices []int
    changes []int
}

type sequence struct {
    a, b, c, d int
}

func newBuyer(initial int64) buyer {
    s := initial
    prices := []int{}
    changes := []int{}
    prev := int(s % 10)
    for i:=0; i<2000; i++ {
        s = genSecret(s)
        p := int(s % 10)
        prices = append(prices, p)
        changes = append(changes, p-prev)
        prev = p
    }
    return buyer{
        secret: s,
        prices: prices,
        changes: changes,
    }
}

func (b buyer) bananas(seq sequence) int {
    for n:=3; n+3 < 2000; n++ {
        d1, d2, d3, d4 := b.changes[n], b.changes[n+1], b.changes[n+2], b.changes[n+3]
        if d1 == seq.a && d2 == seq.b && d3 == seq.c && d4 == seq.d {
            return b.prices[n+3]
        }
    }
    return 0
}

func main() {
    buyers := []buyer{}
    lib.ReadFileByLine(22, func(line string) {
        n := lib.MustParseInt(line)
        buyers = append(buyers, newBuyer(n))
    })
    var p1 int64
    for _, b := range buyers {
        p1 += b.secret
    }
    lib.WritePart1("%d", p1)

    seqsPerBuyer := map[sequence]int{}
    for _, b := range buyers {
        m := map[sequence]struct{}{}
        for n:=3; n+3 < 2000; n++ {
            d1, d2, d3, d4 := b.changes[n], b.changes[n+1], b.changes[n+2], b.changes[n+3]
            m[sequence{d1,d2,d3,d4}] = struct{}{}
        }
        for seq := range m {
            seqsPerBuyer[seq] += 1
        }
    }
    sequences := []sequence{}
    for k, v := range seqsPerBuyer {
        //if v <= len(buyers)/2 {
        if v < 300 {
            continue
        }
        sequences = append(sequences, k)
    }
    //lib.WritePart2("Seqs considered: %d", len(sequences))
    //lib.WritePart2("Seqs considered: %v", sequences)

    max := 0
    for _, seq := range sequences {
        bananas := 0
        for _, b := range buyers {
            bananas += b.bananas(seq)
        }
        if bananas > max {
            max = bananas
        }
    }
    lib.WritePart2("%d", max)
}

func genSecret(seed int64) int64 {
    seed = ((seed * 64) ^ seed) % 16777216
    seed = ((seed / 32) ^ seed) % 16777216
    seed = ((seed * 2048) ^ seed) % 16777216
    return seed
}
