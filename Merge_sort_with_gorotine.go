package main

import "fmt"

func main() {
	s := []int{9, 7, 12, 14, 8, 2, 5, 30, 13, 11, 4, 6, 18, 19, 56, 34}
	result := make(chan []int)
	go Mergesort(s, result)
        r:=<-result
	for _, v := range r {
		fmt.Print(v , " ")
	}
	close(result)
}
func Mergesort(data []int, r chan []int) {
	if len(data) == 1 {
		r <- data
		return
	}
	leftchain := make(chan []int)
	rightchain := make(chan []int)
	middle := len(data) / 2
	go Mergesort(data[:middle], leftchain)
	go Mergesort(data[middle:], rightchain)
        leftdata:=<-leftchain
	rightdata:=<- rightchain
	close(leftchain)
	close(rightchain)
	r <- Merge(leftdata, rightdata)
	return

}
func Merge(leftdata []int, rightdata []int) (result []int) {
	result = make([]int, len(leftdata)+len(rightdata))
	leftindex, rightindex := 0, 0
	for i := 0; i < cap(result); i++ {
		if leftindex >= len(leftdata) {
			result[i] = rightdata[rightindex]
			rightindex++
		} else if rightindex >= len(rightdata) {
			result[i] = leftdata[leftindex]
			leftindex++
		} else if leftdata[leftindex] < rightdata[rightindex] {
			result[i] = leftdata[leftindex]
			leftindex++
		} else {
			result[i] = rightdata[rightindex]
			rightindex++
		}
	}
	return
}