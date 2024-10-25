package main
import "fmt"

type Node struct{
	Next *Node
	Value interface{}
}

func detectCycle(first *Node){
	visited := make(map[*Node]bool) //Create a map to track the visited nodes

	for n := first ; n!= nil;n = n.Next{
		if visited[n]{
			fmt.Println("Cycle detected")
			break
		}
		visited[n] = true
		fmt.Println(n.Value)
	}
}

func main(){
	node1 := &Node{Value: "A"}
	node2 := &Node{Value : "B"}
	node3 := &Node{Value :"C"}

	//Create a linked list node1-> node2 ->node3 ->nil
	node1.Next = node2
	node2.Next = node3

	//Detect a cycle in the linked list(No cycle in this case)
	detectCycle(node1)

	//Introduce a cycle: node3 -> node1
	node3.Next = node1

	//Detect a cycle again(Cycle will be detected this time)
	detectCycle(node1)
}