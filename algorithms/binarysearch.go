package main

import (
	"fmt"
)

func binarySearch(arr []int, start, end, elem int) string {
	if start > end {
		return "element not present"
	}

	mid := int((end + start) / 2)

	if arr[mid] > elem {
		return binarySearch(arr, start, mid, elem)
	} else if arr[mid] < elem {
		return binarySearch(arr, mid+1, end, elem)
	} else {
		return fmt.Sprintf("element found in %d", mid)
	}
}

func iterBinarySearch(arr []int, start, end, elem int) string {
	for start < end {
		mid := int((end + start) / 2)

		if arr[mid] > elem {
			end = mid
		} else if arr[mid] < elem {
			start = mid
		} else {
			return fmt.Sprintf("element found in %d", mid)
		}
	}
	return "element not found in the array"
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(a, 0, len(a)-1, 10))
	fmt.Println(iterBinarySearch(a, 0, len(a)-1, 3))
}
