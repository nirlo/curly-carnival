/*
	Testing concurrency!

	Authored by: Nicholas Lockhart
	Course: CST 8333 Program Language Reseach Project
	 Professor: Stan Pieda
*/package main

import (
	"testing"
)

// Testing to ensure that the data from the input is the same and in
// the same order as the first. testing concurrent
func testProcessData(t *testing.T) {
	var c = make(chan string)
	processData(c)

	t.Error("We have a failure!")
	for i, val := range data {
		if val != output[i] {
			t.Error("We have a failure!")
		}
	}
}

// Testing to ensure that the data from the input is the same and in
// the same order as the first. testing a normal function
func testPlainProcess(t *testing.T) {
	plainProcess()
	for i, val := range data {
		if val != output[i] {
			t.Error("We have a failure!")
		}
	}
}
