package main

import "fmt"

func main(){
	data:=[3]int{1,2,3}
	Permutation(data[:],0,3)
}
func Permutation(data []int ,i int,length int){
	if length==i{
		fmt.Print("\n",data)
		return
	}
	for j:=i;j<length;j++{
		swap(data,i,j)
		Permutation(data,i+1,length)
		swap(data,i,j)
	}
}
func swap(data  []int,x int ,y int){
	data[x],data[y]=data[y],data[x]
}