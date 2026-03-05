package main

// Definition for a QuadTree node.
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	var build func(r, c, length int) *Node
	build = func(r, c, length int) *Node {
		firstVal := grid[r][c]
		isUniform := true

		for i := r; i < r+length; i++ {
			for j := c; j < c+length; j++ {
				if grid[i][j] != firstVal {
					isUniform = false
					break
				}
			}
			if !isUniform {
				break
			}
		}

		if isUniform {
			return &Node{
				Val:    firstVal == 1,
				IsLeaf: true,
			}
		}

		half := length / 2
		return &Node{
			Val:         true,
			IsLeaf:      false,
			TopLeft:     build(r, c, half),
			BottomLeft:  build(r+half, c, half),
			TopRight:    build(r, c+half, half),
			BottomRight: build(r+half, c+half, half),
		}
	}

	return build(0, 0, len(grid))
}
