package main

import (
	"fmt"
	"math/rand"
	"time"

	"../sortlib/mysort"
)

func main() {
	rand.Seed(time.Now().Unix()) //Seed setting
	n := 10000 //data size

	data := make([]int, n) // create slice
	for i := 0; i < n; i++ {
		data[i] = rand.Intn(100)
	}

	fmt.Println(data)
	start := time.Now() 
	
	// Choose what you like ==================================
	fmt.Println(mysort.Insertionsort(data))
	//fmt.Println(mysort.Mergesort(data))
	//fmt.Println(mysort.Quicksort(data))
	//====================================Uai================
	end := time.Now()
	fmt.Printf("%fs\n", (end.Sub(start)).Seconds()) //Execution time
}
