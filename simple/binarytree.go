package yougo

import "fmt"

func tryit() {
	root := &node{
		v:     1,
		left:  nil,
		right: nil,
	}
	root.insertValue(0)
	root.insertValue(1)
	root.insertValue(2)
	root.insertValue(3)
	root.insertValue(4)
	//print output
	root.printByOrder()
	fmt.Println()
	printByLevel(root)
	// search
	fmt.Println(search(root, 2))
	fmt.Println(search(root, 1))
	fmt.Println(search(root, 5))
	//search again
	fmt.Println(root.search(2))
	fmt.Println(root.search(1))
	fmt.Println(root.search(5))
}

type node struct {
	v     int
	left  *node
	right *node
}

func (n *node) insertValue(v int) {
	if v < n.v {
		if n.left == nil {
			n.left = &node{v: v}
		} else {
			n.left.insertValue(v)
		}
	} else if v > n.v {
		if n.right == nil {
			n.right = &node{v: v}
		} else {
			n.right.insertValue(v)
		}
	} else {
		//same value
		return
	}

}

func (n *node) printByOrder() {
	if n.left != nil {
		n.left.printByOrder()
	}
	fmt.Printf("%v ", n.v)
	if n.right != nil {
		n.right.printByOrder()
	}
}

func printByLevel(root *node) {
	q := []*node{root}

	for len(q) > 0 {
		numberAtLevel := len(q)
		for i := 0; i < numberAtLevel; i++ {
			node := q[i]
			fmt.Printf("%v ", node.v)
			// add chidlren to the queue
			if node.left != nil {
				q = append(q, node.left)
			}
			if node.right != nil {
				q = append(q, node.right)
			}
		}
		// queue now starts from the children level
		fmt.Println()
		q = q[numberAtLevel:]
	}
}

func search(root *node, value int) *node {
	if root.v == value {
		return root
	} else if root.v > value {
		if root.left != nil {
			return search(root.left, value)
		} else {
			return nil
		}
	} else if root.v < value {
		if root.right != nil {
			return search(root.right, value)
		} else {
			return nil
		}
	}
	return nil
}

func (n *node) search(value int) bool {
	if n == nil {
		return false
	}
	if n.v == value {
		return true
	} else if value < n.v {
		return n.left.search(value)
	} else {
		return n.right.search(value)
	}
}
