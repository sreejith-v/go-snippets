package main

import "fmt"

func main() {
	ex := []int{2, 1, 1, 1, 1, 99, 99, 1, 1, 1, 1, 1, 1, 1, 1, 2}

	leftPtr := 0
	rightPtr := len(ex) - 1
	volume := 0

	maxVolume := 0
	for leftPtr < rightPtr {
		volume = min(ex[leftPtr], ex[rightPtr]) * (rightPtr - leftPtr)
		maxVolume = max(maxVolume, volume)
		if ex[leftPtr] <= ex[rightPtr] {
			leftPtr += 1
		} else {
			rightPtr -= 1
		}
	}
	fmt.Println(maxVolume)
}
