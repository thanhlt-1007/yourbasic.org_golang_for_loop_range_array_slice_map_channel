package main

import "fmt"

func main() {
    // Basic for-each loop
    a := []string{"Foo", "Bar"}
    for i, s := range a {
        fmt.Println(i, s)
    }

    // String iteration: runes or bytes
    for i, ch := range "日本語" {
        fmt.Printf("%#U starts at byte position %d\n", ch, i)
    }

    const s = "日本語"
    for i := 0; i < len(s); i ++ {
        fmt.Printf("%x ", s[i])
    }

    // Map iteration: keys and values
    m := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
    }
    for k, v := range m {
        fmt.Println(k, v)
    }
}
