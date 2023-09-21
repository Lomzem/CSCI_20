package main

import "fmt"

var gnumber int = 1

func main() {

   var number int = 2

   if number < number {
      var number int = 3
      fmt.Println(number)
   }

   fmt.Println(number)

   {
      var number int = 4

      if 1 <= 1 {
         var number int = 5
         fmt.Println(number)
      }

      fmt.Println(number)
   }
}
