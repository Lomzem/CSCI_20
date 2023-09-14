/***************************************
Lotto Numbers
Author: Lawjay Lee
Date Completed: 2023-09-01
Description: Prints out 5 random ints between 1-47
***************************************/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
    rand.Seed(int64(time.Now().Nanosecond()))
}

func main() {
    fmt.Println("Lucky Lotto Numbers:")

    var num_1 int = rand.Intn(47) + 1
    var num_2 int = rand.Intn(47) + 1
    var num_3 int = rand.Intn(47) + 1
    var num_4 int = rand.Intn(47) + 1
    var num_5 int = rand.Intn(47) + 1

    fmt.Println(num_1, num_2, num_3, num_4, num_5)
}
