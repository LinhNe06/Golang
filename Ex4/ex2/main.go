package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}
	if value < root.Value {
		root.Left = insert(root.Left, value)
	} else {
		root.Right = insert(root.Right, value)
	}
	return root
}

func inOrderDESC(root *Node, result *[]int) {
	if root != nil {
		inOrderDESC(root.Right, result)
		*result = append(*result, root.Value)
		inOrderDESC(root.Left, result)
	}
}

func main() {
	data1 := []int{15, 10, 20, 8, 12, 17, 25}
	data2 := []int{50, 30, 70, 20, 40, 60, 80}
	data3 := []int{10, 5, 15, 3, 7, 12, 18}
	data4 := []int{100, 20, 150, 10, 30, 110, 200}

	var tree1, tree2, tree3, tree4 *Node

	for _, v := range data1 {
		tree1 = insert(tree1, v)
	}
	for _, v := range data2 {
		tree2 = insert(tree2, v)
	}
	for _, v := range data3 {
		tree3 = insert(tree3, v)
	}
	for _, v := range data4 {
		tree4 = insert(tree4, v)
	}

	var res1, res2, res3, res4 []int
	inOrderDESC(tree1, &res1)
	inOrderDESC(tree2, &res2)
	inOrderDESC(tree3, &res3)
	inOrderDESC(tree4, &res4)

	fmt.Println("Cây 1:", res1)
	fmt.Println("Cây 2:", res2)
	fmt.Println("Cây 3:", res3)
	fmt.Println("Cây 4:", res4)

	fmt.Println("\nSo sánh")
	fmt.Printf("Gốc lớn nhất: Cây 1 (%d) | Cây 2 (%d) | Cây 3 (%d) | Cây 4 (%d)\n",
		res1[0], res2[0], res3[0], res4[0])
}
