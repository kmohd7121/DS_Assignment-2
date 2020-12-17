package main

import "fmt"
type Node struct {
	Data int 
	left *Node
	right *Node
	height int
}
func main(){
	rootnode:=new(Node)
	rootnode=nil
	rootnode=insert(rootnode,10)
	rootnode=insert(rootnode,20)
	rootnode=insert(rootnode,30)
	rootnode=insert(rootnode,40)
	rootnode=insert(rootnode,50)
	rootnode=insert(rootnode,25)

	fmt.Println("Preorder traversal of the constructed AVL tree is \n")
	preOrder(rootnode)


}
func insert(n *Node ,data int) *Node{
	if n==nil{
		newNode:=Node{data,nil,nil,1}
		return &newNode
	}
	if data < n.Data{
		n.left=insert(n.left,data)
	}else if data >n.Data{
		n.right=insert(n.right,data)
	}else{
		return n
	}
	n.height=1+max(height(n.left),height(n.right))
	balance:=getBalance(n)
	if balance > 1 && n.left!=nill && data < n.left.Data{
		fmt.Println("balance 1")
		return rightRotate(n)
	}
	if balance < -1 && n.right!=nil && data > n.right.Data{
		fmt.Println("balance 2")
		return leftRotate(n)

	}
	if balance >1 && n.left!=nil && data >n.right.Data{
		n.left=leftRotate(n.left)
		fmt.Println("balance 3")
		return rightRotate(n)
	}
	if balance < -1 && n.right!=nil && data < n.right.Data{
		n.right=rightRotate(n.right)
		fmt.Println("balance 4")
		return leftRotate(n)
	}
	return n
}
func rightRotate(y *Node) *Node{
	x:=y.left
	t2:=x.right
	x.right=y
	y.left=t2
	y.height=max(height(y.left),height(y.right))+1
	x.height=max(height(x.left),height(x.right))+1
	return x
}
func leftRotate(x *Node) *Node{
	y:=x.right
	t2:=y.left
	y.left=x
	x.right=t2
	x.height=max(height(x.left),height(x.right))+1
	y.height=max(height(y.left),height(y.right))+1
	return y
}
func max(a,b int)int{
	if a>b{
		return a
	}else {
		return b
	}
}
func height(n *Node) int{
	if n==nil{
		return 0
	}else{
		return n.height
	}
}
func preOrder(n *Node){
	if n!=nil{
		fmt.Println(n.Data," ")
		preOrder(n.left)
		preOrder(n.right)

	}
}