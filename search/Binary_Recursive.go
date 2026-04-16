package search

func BinarySearchRecursive(arr []int, target int, left int, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return BinarySearchRecursive(arr, target, mid+1, right)
	} else {
		return BinarySearchRecursive(arr, target, left, mid-1)
	}
}
