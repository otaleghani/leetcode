package main

// Definition for a Node.
type Node struct {
    Val int
    Neighbors []*Node
}


func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    
    nodes := make(map[int]*Node)
    var create func(int, []*Node) *Node

    create = func(val int, neighbors []*Node) *Node {
        curr := &Node{
            Val: val,
        }
        nodes[val] = curr
        
        for _, neighbor := range neighbors {
            if node, exists := nodes[neighbor.Val]; exists {
                curr.Neighbors = append(curr.Neighbors, node)
            } else {
                node := create(neighbor.Val, neighbor.Neighbors)
                curr.Neighbors = append(curr.Neighbors, node)
            }
        }
        
        return curr
    }

    return create(node.Val, node.Neighbors)
}
