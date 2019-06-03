package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomizeAndSort(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice

}

func BubbleSort(items []int) []int {

	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	return items
}

func BinarySearch(sortedList []int, lookingFor int) int {

	var lo int = 0

	var hi int = len(sortedList) - 1

	for lo <= hi {
		var mid int = lo + (hi-lo)/2
		var midValue int = sortedList[mid]
		fmt.Println("Middle value is:", midValue)

		if midValue == lookingFor {
			return mid
		} else if midValue > lookingFor {
			// We want to use the left half of our list
			hi = mid - 1
		} else {
			// We want to use the right half of our list
			lo = mid + 1
		}
	}

	// If we get here we tried to look at an invalid sub-list
	// which means the number isn't in our list.
	fmt.Print("Number not found")
	return -1

}

func ExecuteBinarySearch(size int, element int) {

	//Get a random array which is sorted in nature

	array := RandomizeAndSort(size)

	//Sort the Array

	sortedArray := BubbleSort(array)

	fmt.Println("Sorted Array is %v", sortedArray)

	//Check if the number exists or not
	//and get the printed result

	BinarySearch(array, element)

}
