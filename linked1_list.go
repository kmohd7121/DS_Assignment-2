package main

import "fmt"

func main() {
	list := new(node)
	list.pushfront(18)
	list.pushfront(19)
	list.pushback(9)
	list.popfront()
	list.popback()
	for head != nil {
		fmt.Printf("%d ", head.data)
		head = head.next
	}
}

type node struct {
	data int
	next *node
}

var head *node = nil
//push element in the front of linked list
func (f *node) pushfront(val int) *node {
	if head == nil {
		f.data = val
		f.next = nil
		head = f
		return f
	} else {
		x := new(node)
		x = head
		y:=&node{data: val,next: x,}
		head = y
		return head
	}
}
// push element back to the linked list
func (f *node) pushback(val int) *node{
	if head ==nil{
		f.data=val
		f.next=nil
		head=f
		return f
	}else{
		for f.next!=nil{
			f=f.next
		}
		f.next=new(node)
		f.next.data=val
		f.next.next=nil
		return f
	}
}
// pop element front of the linked list
func (f *node) popfront() *node{
	if head==nil{
		return head
	}
	popnode:=new(node)
	popnode=head.next
	head=popnode
	return head
}
// pop element back to the linked list
func (f *node) popback() *node{
	if head ==nil{
		return head
	}
	popnode:=new(node)
	popnode=head
	for popnode.next.next!=nil{
		popnode=popnode.next
	}
	popnode.next=nil
	return head
}