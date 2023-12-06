package main

import (
    "bufio"
    "bytes"
    "fmt"
    "log"
    "os"
    "strconv"
    "unicode"
)

var (
    validChars = []byte("1234567890.")
)

func main() {
    ptrFile, err := os.Open("input.txt")
    if err != nil {
        log.Fatalf("while opening the file %q: %s", ptrFile.Name(), err)
    }
    defer ptrFile.Close()

    lines := make([][]byte, 0)
    scnr := bufio.NewScanner(ptrFile)
    for scnr.Scan() {
        lines = append(lines, []byte(scnr.Text()))
    }

    sum := int64(0)
    for i := 0; i < len(lines); i++ {
        switch {
        case i == 0:
            sum += fl(lines[i], lines[i+1])
        case i == len(lines)-1:
            sum += ll(lines[i], lines[i-1])
        case i > 0 && i < len(lines)-1:
            sum += ml(lines[i], lines[i+1], lines[i-1])
        }
    }

    fmt.Println(sum)

}

func iD(c byte) bool {
    return unicode.IsDigit(rune(c))
}

func iS(c byte) bool {
    return !bytes.ContainsAny(validChars, string(c))
}

func fl(cl, bl []byte) int64 {
    sum := int64(0)
    for i := 0; i < len(cl); i++ {
        if cl[i] == '.' || iS(cl[i]) {
            continue
        }

        fPN := false
        j := i

        switch {
        case i == 0 && iD(cl[i]):
            for iD(cl[j]) {
                if iS(cl[j+1]) || iS(bl[j]) || iS(bl[j+1]) {
                    fPN = true
                }
                j++
            }
        case i > 0 && iD(cl[i]):
            for j < len(cl) && iD(cl[j]) {
                if j < len(cl)-1 && (iS(cl[j-1]) || iS(cl[j+1]) || iS(bl[j-1]) || iS(bl[j]) || iS(bl[j+1])) {
                    fPN = true
                }
                if j < len(cl) && (iS(cl[j-1]) || iS(bl[j-1]) || iS(bl[j])) {
                    fPN = true
                }
                j++
            }

        }

        switch {
        case fPN && j-i > 0:
            n, _ := strconv.Atoi(string(cl[i:j]))
            sum += int64(n)
            i = j - 1
        case j-i > 0:
            i = j - 1
        }
    }
    return sum
}

func ll(cl, al []byte) int64 {
    sum := int64(0)
    for i := 0; i < len(cl); i++ {
        if cl[i] == '.' || iS(cl[i]) {
            continue
        }

        fPN := false
        j := i

        switch {
        case i == 0 && iD(cl[i]):
            for iD(cl[j]) {
                if iS(cl[j+1]) || iS(al[j]) || iS(al[j+1]) {
                    fPN = true
                }
                j++
            }
        case i > 0 && iD(cl[i]):
            for j < len(cl) && iD(cl[j]) {
                if j < len(cl)-1 && (iS(cl[j-1]) || iS(cl[j+1]) || iS(al[j-1]) || iS(al[j]) || iS(al[j+1])) {
                    fPN = true
                }
                if j < len(cl) && (iS(cl[j-1]) || iS(al[j-1]) || iS(al[j])) {
                    fPN = true
                }
                j++
            }

        }

        switch {
        case fPN && j-i > 0:
            n, _ := strconv.Atoi(string(cl[i:j]))
            sum += int64(n)
            i = j - 1
        case j-i > 0:
            i = j - 1
        }
    }
    return sum
}

func ml(cl, bl, al []byte) int64 {
    sum := int64(0)
    for i := 0; i < len(cl); i++ {
        if cl[i] == '.' || iS(cl[i]) {
            continue
        }

        fPN := false
        j := i

        switch {
        case i == 0 && iD(cl[i]):
            for iD(cl[j]) {
                if iS(al[j]) || iS(al[j+1]) || iS(cl[j+1]) || iS(bl[j]) || iS(bl[j+1]) {
                    fPN = true
                }
                j++
            }
        case i > 0 && iD(cl[i]):
            for j < len(cl) && iD(cl[j]) {
                if j < len(cl)-1 && (iS(cl[j-1]) || iS(cl[j+1]) ||
                    iS(al[j-1]) || iS(al[j]) || iS(al[j+1]) ||
                    iS(bl[j-1]) || iS(bl[j]) || iS(bl[j+1])) {
                    fPN = true
                }
                if j < len(cl) && (iS(cl[j-1]) ||
                    iS(al[j-1]) || iS(al[j]) ||
                    iS(bl[j-1]) || iS(bl[j])) {
                    fPN = true
                }
                j++
            }
        }

        switch {
        case fPN && j-i > 0:
            n, _ := strconv.Atoi(string(cl[i:j]))
            sum += int64(n)
            i = j - 1
        case j-i > 0:
            i = j - 1
        }
    }
    return sum
}
