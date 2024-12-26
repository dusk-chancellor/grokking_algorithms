package main

import "fmt"

func checkBin(list []int, i int) interface{} {
	// low and high keep track of which part of the list you'll search in.
	low := 0
	high := len(list) - 1
	// While you haven't narrowed it down to one element ...
	for low <= high {
		// ... check the middle element
		mid := (low + high) / 2
		// Found the item.
		if list[mid] == i {
			return mid
		}
		// The guess was too high.
		if list[mid] > i {
			high = mid - 1
		// The guess was too low.
		} else {
			low = mid + 1
		}
	}
	// Item doesn't exist
	return nil
}

func recursiveCheckBin(list []int, item int, high, low int) interface{} {
	// Check base case
	if high >= low {
		mid := (high + low) / 2
		// If element is present at the middle itself
		if list[mid] == item {
			return mid
		// If element is smaller than mid, then it can only 
        // be present in left subarray
		} else if list[mid] > item {
			return recursiveCheckBin(list, item, mid - 1, low)
		// Else the element can only be present in right subarray 
		} else {
			return recursiveCheckBin(list, item, high, mid + 1)
		}

	}
	// Element is not present in the array 
	return nil
}
// you can run this code on: https://go.dev/play/
func main() {
	list := []int{1, 3, 5, 7, 9}
	fmt.Println(checkBin(list, 3))							// => 1
	fmt.Println(checkBin(list, -1))							// => <nil>
	fmt.Println(recursiveCheckBin(list, 3, len(list)-1, 0)) // => 1
	fmt.Println(recursiveCheckBin(list, -1, len(list)-1, 0)) // => <nil>
}
