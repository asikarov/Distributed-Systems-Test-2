package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	question1()
	question2()
}

func sumHelper(s []int) int {
	// sums slice
	sum := 0
	for _, e := range s {
		sum += e
	}
	return sum
}

func sumHelperCh(s []int, c chan int) {
	// sourced from Go tour channel page: https://go.dev/tour/concurrency/2
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func question1() {
	// get input from command line
	x := os.Args[1]
	num, _ := strconv.Atoi(x)

	// append random ints
	var s []int
	for i := 0; i <= num; i++ {
		s = append(s, rand.Int())
	}

	// sum slice
	firstTimeSlice := time.Now()
	sum := sumHelper(s)
	lastTimeSlice := time.Now()
	fmt.Println("standard sum: ", sum)
	diffSlice := lastTimeSlice.Sub(firstTimeSlice)
	fmt.Println("time taken: ", diffSlice)

	// sum slice in parallel with 2 go-routines, sourced from Go tour channel page: https://go.dev/tour/concurrency/2
	c := make(chan int)
	firstTimeSlice = time.Now()
	go sumHelperCh(s[:len(s)/2], c)
	go sumHelperCh(s[len(s)/2:], c)
	s1, s2 := <-c, <-c // receive from c
	sum = s1 + s2
	lastTimeSlice = time.Now()
	fmt.Println("concurrent sum: ", sum)
	diffSlice = lastTimeSlice.Sub(firstTimeSlice)
	fmt.Println("time taken: ", diffSlice)

}

func question2() {
	// get input from command line
	x := os.Args[1]
	num, _ := strconv.Atoi(x)

	// append random ints
	var s []int
	for i := 0; i <= num; i++ {
		s = append(s, rand.Int())
	}
	var sliceCopy []int
	copy(s, sliceCopy)

	// sort slice
	firstTimeSlice := time.Now()
	sort.Slice(s, func(i int, j int) bool { return i < j })
	lastTimeSlice := time.Now()
	diffSlice := lastTimeSlice.Sub(firstTimeSlice)
	fmt.Println("standard sort time: ", diffSlice)

	// stable sort sliceCopy
	firstTimeSlice = time.Now()
	sort.SliceStable(sliceCopy, func(i int, j int) bool { return i < j })
	lastTimeSlice = time.Now()
	diffSlice = lastTimeSlice.Sub(firstTimeSlice)
	fmt.Println("stable sort time: ", diffSlice)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
