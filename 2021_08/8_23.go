package _2021_08

import (
	"fmt"
	"time"
)

// Implement a job scheduler which takes in a
//function f and an integer n, and calls f after
//n milliseconds
const count = 4

type f func()

func F() {
	fmt.Println("DONE")
}

func Scheduler(function f, i int) {
	time.Sleep(time.Duration(i)*time.Millisecond)
	function()
}