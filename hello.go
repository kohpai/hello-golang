package main

import (
    "fmt"

    "github.com/kohpai/hello-golang/morestrings"
    "github.com/google/go-cmp/cmp"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(cmp.Diff("Hello, world", "Hello Go"))
}
