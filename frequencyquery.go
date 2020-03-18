package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
    out := make([]int32,0)
    mapp := make(map[int32]int32)
    freq := make(map[int32]int32)
    for i := 0;i<len(queries);i++ {
        a,b := queries[i][0] , queries[i][1] 
        switch a {
            case 1:
                if freq[mapp[b]] != 0 {
                //update frequency table
                  freq[mapp[b]] = freq[mapp[b]]-1  
                }
                mapp[b] = mapp[b]+1
                freq[mapp[b]] = freq[mapp[b]]+1
            case 2:
                if mapp[b] != 0 {
                    freq[mapp[b]] = freq[mapp[b]]-1
                    mapp[b] = mapp[b]-1
                    freq[mapp[b]] = freq[mapp[b]]+1 
                }
            case 3:
                if freq[b] != 0 {
                    out = append(out, 1 )
                } else {
                    out = append(out, 0 )
                } 
        }
    }
    return out
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(qTemp)

    var queries [][]int32
    for i := 0; i < int(q); i++ {
        queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var queriesRow []int32
        for _, queriesRowItem := range queriesRowTemp {
            queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
            checkError(err)
            queriesItem := int32(queriesItemTemp)
            queriesRow = append(queriesRow, queriesItem)
        }

        if len(queriesRow) != 2 {
            panic("Bad input")
        }

        queries = append(queries, queriesRow)
    }

    ans := freqQuery(queries)

    for i, ansItem := range ans {
        fmt.Fprintf(writer, "%d", ansItem)

        if i != len(ans) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
