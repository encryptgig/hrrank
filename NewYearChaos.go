package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
    var out int32
    // check if difference is 2
    for i := int32(1);i<= int32(len(q));i++ {
        if q[i-1] > i {
            if q[i-1] - i > int32(2) {
                fmt.Println("Too chaotic")
                return
            }
        }
    }

    var i int32

    // since the differnce is only 2 positions we can do an insertion sort
    for i =0;i< int32(len(q))-1;i++ {
        j := i+1
        p := q[j]
        for j > 0 && q[j-1] > p  {
            q[j] = q[j-1]
            j--
            out++
        }
        q[j] = p
    }

    
    fmt.Println(out)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        qTemp := strings.Split(readLine(reader), " ")

        var q []int32

        for i := 0; i < int(n); i++ {
            qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
            checkError(err)
            qItem := int32(qItemTemp)
            q = append(q, qItem)
        }

        minimumBribes(q)
    }
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
