package main

import "fmt"

func main() {
    // Basic for-each loop
    a := []string{"Foo", "Bar"}
    for i, s := range a {
        fmt.Println(i, s)
    }
}
