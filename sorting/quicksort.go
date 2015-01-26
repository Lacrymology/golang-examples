// Package sorting provides a number of sorting algorithms for user-defined
// collections that satisfy sort.Interface
package sorting

import "sync"
import "sort"

// Quicksort runs the quicksort algorithm on a sortable collection in-place.
// It is a parallelized implementation of the algorithm
func Quicksort(s sort.Interface) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go quicksort(s, 0, s.Len()-1, wg)
	wg.Wait()
}

// quicksort runs the actual quicksort algorithm on the sort.Interface object.
// from and to are the indices of the sub-collection to sort (to is inclusive)
func quicksort(s sort.Interface, from int, to int, done *sync.WaitGroup) {

	defer done.Done()

	if from >= to {
		return
	}

//	if sort.IsSorted(s) {
//		return
//	}

	p := partition(s, from, to)
	if p > 0 {
		done.Add(1)
		go quicksort(s, from, p-1, done)
	}
	if p+1 < s.Len() {
		done.Add(1)
		go quicksort(s, p+1, to, done)
	}
}

// partition finds a good pivot, partitions the elements in bigger and smaller
// than the pivot, and returns the pivot's new index
func partition(s sort.Interface, from int, to int) int {
	// move pivot to the end
	s.Swap(from, to)

	// last sorted element
	walk := from
	for i := from; i < to; i++ {
		if s.Less(i, to) {
			s.Swap(i, walk)
			walk += 1
		}
	}
	s.Swap(walk, to)
	return walk
}
