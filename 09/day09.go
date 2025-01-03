package main

import (
    "github.com/deosjr/adventofcode2024/lib"
)

type block struct {
    id, size, pos int
}

func main() {
    input := lib.ReadFile(9)
    files := []block{}
    free := []block{}
    index := 0
    for i, c := range input[:len(input)-1] {
        n := int(c) - 48
        if n == 0 {
            continue
        }
        if i % 2 == 0 {
            files = append(files, block{
                id: len(files),
                size: n,
                pos: index,
            })
        } else {
            free = append(free, block{
                size: n,
                pos: index,
            })
        }
        index += n
    }
    fileCopy, freeCopy := make([]block, len(files)), make([]block, len(free))
    copy(fileCopy, files)
    copy(freeCopy, free)
    ans1 := p1(files, free, nil)
    lib.WritePart1("%d", ans1)
    ans2 := p2(fileCopy, freeCopy, nil)
    lib.WritePart2("%d", ans2)
}

func checksum(files []block) int {
    sum := 0
    for _, file := range files {
        for i:=file.pos; i<file.pos+file.size; i++ {
            sum += i * file.id
        }
    }
    return sum
}

func p1(files, free, newfiles []block) int {
    if len(free) == 0 {
        return checksum(append(files, newfiles...))
    }
    leftmostfree := free[0]
    rightmostfile := files[len(files)-1]
    if leftmostfree.pos > rightmostfile.pos {
        return checksum(append(files, newfiles...))
    }
    if rightmostfile.size <= leftmostfree.size {
        rightmostfile.pos = leftmostfree.pos
        newfiles = append(newfiles, rightmostfile)
        files = files[:len(files)-1]
        if rightmostfile.size == leftmostfree.size {
            free = free[1:]
        } else {
            leftmostfree.size = leftmostfree.size - rightmostfile.size
            leftmostfree.pos = leftmostfree.pos + rightmostfile.size
            free[0] = leftmostfree
        }
    } else {
        rightmostfile.size = rightmostfile.size - leftmostfree.size
        files[len(files)-1] = rightmostfile
        free = free[1:]
        newfiles = append(newfiles, block{
            id: rightmostfile.id, 
            size: leftmostfree.size,
            pos: leftmostfree.pos,
        })
    }
    return p1(files, free, newfiles)
}

func p2(files, free, newfiles []block) int {
    if len(free) == 0 {
        return checksum(append(files, newfiles...))
    }
    rightmostfile := files[len(files)-1]
    if rightmostfile.pos < free[0].pos {
        return checksum(append(files, newfiles...))
    }
    files = files[:len(files)-1]
    for i, space := range free {
        if space.pos > rightmostfile.pos {
            break
        }
        if space.size >= rightmostfile.size {
            rightmostfile.pos = space.pos
            if space.size == rightmostfile.size {
                if i == len(free)-1 {
                    free = free[:i]
                } else {
                    free = append(free[:i], free[i+1:]...)
                }
            } else {
                space.size = space.size - rightmostfile.size
                space.pos = space.pos + rightmostfile.size
                free[i] = space
            }
            break
        }
    }
    newfiles = append(newfiles, rightmostfile)
    return p2(files, free, newfiles)
}
