package main

import (
	"fmt"
)

type node struct{

	data string
	left *node
	right *node

}
func buildNode( value string) *node{

	newNode := node{ value , nil , nil }
	return &newNode

}

func preOrder( root *node)(){
	if root == nil {
		return
	}
	fmt.Print(root.data + " ")
	preOrder( root.left )
	preOrder ( root.right )
}

func postOrder( root *node)(){
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Print(root.data + " ")
}


func main(){

	root := buildNode("+")
	root.left = buildNode( "a")
	root.right = buildNode( "-")
	root.right.left = buildNode( "b")
	root.right.right = buildNode( "c")

	fmt.Print(" Preorder traversal : ")
	preOrder( root )
	fmt.Print("\n Postorder traversal : " )
	postOrder( root)

}


