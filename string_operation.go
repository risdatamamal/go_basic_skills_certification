package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)



/*
 * Complete the 'ModifyString' function below and add imports if needed.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING str as parameter.
 */

var (
    digits = map[rune]struct{}{
        '0': struct{}{},
        '1': struct{}{},
        '2': struct{}{},
        '3': struct{}{},
        '4': struct{}{},
        '5': struct{}{},
        '6': struct{}{},
        '7': struct{}{},
        '8': struct{}{},
        '9': struct{}{},
    }
)

func ModifyString(str string) string {
    var validChars []rune
    
    for _, s := range strings.TrimSpace(str) {
        if _, ok := digits[s]; ok {
            continue
        }
        
        validChars = append(validChars, s)
    }
    
    N := len(validChars)
    
    var sb strings.Builder
    
    for i := N - 1; i >= 0; i-- {
        sb.WriteRune(validChars[i])
    }
    
    return sb.String()
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    str := readLine(reader)

    result := ModifyString(str)

    fmt.Fprintf(writer, "%s\n", result)

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