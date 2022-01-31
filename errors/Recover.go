// Go provides a way to capture a panic to provide a more graceful shutdown or to prevent shutdown at all. The built-in recover function is called from within a defer to check if a panic happened. If there was a panic, the value assigned to the panic is returned. Once a recover happens, execution continues normally.

package main

import ("fmt")

func main() {
    for _, val := range []int{1, 2, 0, 6} {
        div60(val)
    }
}

func div60(i int) {
    defer func() {
        if v := recover(); v != nil {
            fmt.Println(v)
        }
    }()
    fmt.Println(60 / i)
}
