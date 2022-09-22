package main

import (
    "fmt"
    "strconv"
)

/*
 * Complete the 'countBits' function below.
 *
 * The function is expected to return an int32.
 * The function accepts unit32 num as parameter.
 */

func countBits(n int) int {
    count := 0
    for n != 0 {
        count += n&1
        n >>= 1
    }
    return count
}

func main() {
    var n int
    fmt.Scanln(&n)
    strconv.FormatInt(int64(n), 2)
    fmt.Printf("%d", countBits(n))
}
