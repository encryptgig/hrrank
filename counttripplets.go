package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func nc3( val int64 ) int64 {
    return int64((val * (val-1) * (val-2)))/6
}

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {

    var count int64
    m1 := map[int64]int64{}
    m2 := map[int64]int64{}

    if r == 1 {
        m1 := map[uint64]int{}
        for i := 0 ; i < len(arr) ; i++ {
            m1[uint64(arr[i])] = m1[uint64(arr[i])]+1
        }
        for _,v := range m1 {
            if v > 2 {
                count += nc3( int64(v) )
            }
        }
        return count
    }

    for i := len(arr) - 1; i >= 0; i-- {
        a := arr[i];
        v,ok := m1[a*r]
        if ok {
            m2[a] = m2[a]+v
        }

        v,ok = m2[a*r]
        if ok {
            count += v
        }
         m1[a] = m1[a]+1
    }

    return count
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(nr[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    r, err := strconv.ParseInt(nr[1], 10, 64)
    checkError(err)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int64

    for i := 0; i < int(n); i++ {
        arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arr = append(arr, arrItem)
    }

    ans := countTriplets(arr, r)

    fmt.Fprintf(writer, "%d\n", ans)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
