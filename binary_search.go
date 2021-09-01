package main

import "fmt"

func binarSearchWithRecusion(arr []int, toSearch int, low int, high int) int {

	mid:=(high+low)/2
	if (low>high) {
		return -1
	}
	if arr[mid]==toSearch{
		return mid 
	}else if( arr[mid] > toSearch){
		return binarSearchWithRecusion(arr,toSearch,low,mid-1)
	}else{
		return binarSearchWithRecusion(arr,toSearch,mid+1,high)
	}
	
}

func binary_search(arr []int, x int) int {
	low := 0
	high := len(arr) - 1
	
	for low<=high{
		mid := (low + high) / 2
		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	var arr = []int{1, 4, 7, 9, 16, 56, 70}
	fmt.Println(binary_search(arr, 0))
	fmt.Println(binarSearchWithRecusion(arr,0,0,len(arr)-1))
}
