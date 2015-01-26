package main

import "fmt"
import "sort"
import "time"
import "math/rand"
import "example/sorting"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	n := rand.Int31n(150)
	n = 10000
	s := make([]int, n)
	for i, _ := range s {
		s[i] = rand.Int() % int(n)
//		s[i] = i
	}
	fmt.Println("Quicksorting", s)
	sorting.Quicksort(sort.IntSlice(s))
	fmt.Println(s)
}
