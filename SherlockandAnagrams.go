package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}



func calculateSubsets(s string, lent int , m map[string]int) {
    for len(s) != 0 {
        a := SortString(s[0:lent])
        m[a] =  m[a]+1
        s = s[1:]
        if len(s) < lent {
            break
        }
    }
}

func combination(a int) int {
    return (a*(a-1))/2
}


// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
    var result int
    m := make(map[string]int)
    for i := 1 ;i<len(s);i++ {
        calculateSubsets(s,i,m)
    }

    for _,v := range m {
        if v >= 2 {
            result += combination(v)
        }
    }
    return int32(result)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        s := readLine(reader)

        result := sherlockAndAnagrams(s)

        fmt.Fprintf(writer, "%d\n", result)
    }

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
