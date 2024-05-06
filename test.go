package main

import (
    "fmt"
    "sort"
)

func main() {
    a := []int{1, 2, 3, 4, 5}

    // Reverse the slice using sort.Reverse
    sort.Sort(sort.Reverse(sort.IntSlice(a)))

    fmt.Println("Reversed slice:", a)
}
