package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

// func factorial(n int) int {
//     var product int = 1
//     for i := 1; i <= n; i++ {
//         product *= i
//     }
//     return product
// }

func main() {
    fmt.Println("Recursive 5!:", factorial(5))
}
