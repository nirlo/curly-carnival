/*
	Checking out concurrency with Golang!

	Authored by: Nicholas Lockhart
	Course: CST 8333 Program Language Reseach Project
	 Professor: Stan Pieda
*/package main

import (
	"fmt"
	"time"
)

// Initalializing the slices for use throughout the program
var (
	data    = []string{"one", "two", "three", "why", "am", "I", "trying", "to", "count", "these", "test1", "test2", "test3", "test4", "5", "6", "7", "8", "9", "final!"}
	output  = []string{}
	output1 = []string{}
)

// Main function
// Prints out the inital values, runs the goroutines
// and prints out the end values
func main() {

	fmt.Println("----Nicholas Lockhart----\noriginal array before: ", data)
	fmt.Println("output array before: ", output)

	var c = make(chan string, 20)

	start := time.Now()
	processData(c)
	// I'm adding a second waiting time here so the timing is easier to read
	time.Sleep(time.Second * 1)
	t := time.Now()
	elapsedCon := t.Sub(start)
	fmt.Println("time elapsed concurrently: ", elapsedCon)

	fmt.Println("original array after: ", data)
	fmt.Println("output array after: ", output)

	start1 := time.Now()
	plainProcess()
	// Without the extra second, this will read as completing in 0s
	time.Sleep(time.Second * 1)
	t1 := time.Now()
	elapsedNorm := t1.Sub(start1)
	fmt.Println("time elapsed normal process: ", elapsedNorm)
	fmt.Println(output1)
}

func processData(c chan string) {
	go get(c)
	go put(c)
}

// Gets the values from the original structure and places them
// in the channel for processing
func get(c chan string) {
	for i, val := range data {
		c <- val
		if i == len(data) {
			close(c)
		}
	}
}

// Takes the values in the channel and puts them in the output
func put(c chan string) {
	for {
		val := <-c
		output = append(output, val)
	}
}

// A plain function for reading the data from one
// array and placing it in another
func plainProcess() {
	for _, val := range data {
		output1 = append(output1, val)
	}
}
