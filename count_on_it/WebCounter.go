/***************************************
Count On It
Author: Lawjay Lee
Date Completed: 2023-09-28
Description: Counts the number of times the site has been visited
***************************************/

package main

import "fmt"

type WebCounter struct {
    count int
}

func NewWebCounter(count int) WebCounter {
    var counter WebCounter
    counter.count = count
    return counter
}

func (w *WebCounter) Get() int {
    return w.count
}

func (w *WebCounter) Hit() {
    w.count += 1
}

func (w *WebCounter) Reset() {
    w.count = 0
}

func main() {

	var wc WebCounter = NewWebCounter(42)

	fmt.Println("Web Counter Tester:")
	fmt.Println("\nStart:", wc.Get())

	wc.Hit()
	wc.Hit()

	fmt.Println("Two Hits:", wc.Get())

	wc.Hit()

	fmt.Println("Three Hits:", wc.Get())

	wc.Reset()

	fmt.Println("Reset:", wc.Get())

	wc.Hit()
	wc.Hit()
	wc.Hit()
	wc.Hit()

	fmt.Println("Four Hits:", wc.Get())
}
