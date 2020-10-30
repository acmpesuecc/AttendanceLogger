package main

// basically iostream
import "fmt"

//Map (same as dict from python)
func main() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

}