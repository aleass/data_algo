package tree

//Node binary tree for linked list
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewLinkedTree(value int) *Node {
	return &Node{Value: value}
}

func (n *Node) PreOrder() {
	if n == nil {
		return
	}
	println(n.Value)
	n.Left.PreOrder()
	n.Right.PreOrder()
}

func (n *Node) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	println(n.Value)
	n.Right.InOrder()
}

func (n *Node) PostOrder() {
	if n == nil {
		return
	}
	n.Left.PostOrder()
	n.Right.PostOrder()
	println(n.Value)
}

func (n *Node) InsertOne2(val int) {
	var node *Node
	if val <= n.Value {
		if n.Left == nil {
			n.Left = NewLinkedTree(val)
			return
		}
		node = n.Left
	} else if val > n.Value {
		if n.Right == nil {
			n.Right = NewLinkedTree(val)
			return
		}
		node = n.Right
	}
	node.InsertOne(val)
}

func (n *Node) InsertOne(val int) {
	var node = n
	for true {
		if val <= node.Value {
			if node.Left == nil {
				node.Left = NewLinkedTree(val)
				return
			}
			node = node.Left
		} else if val > node.Value {
			if node.Right == nil {
				node.Right = NewLinkedTree(val)
				return
			}
			node = node.Right
		}
	}
}

// Get return value is exist
func (n *Node) Get(val int) *Node {
	if n == nil {
		return nil
	}
	var node = n
	for true {
		if node.Value == val {
			return node
		}
		if val <= node.Value && node.Left != nil {
			node = node.Left
		} else if val > node.Value && node.Right != nil {
			node = node.Right
		} else {
			return nil
		}
	}
	return nil
}

//del node
func (n *Node) Del(val int) {
	if n == nil {
		return
	}
	var node, preNode = n, n
	var which string
	for true {
		if node.Value == val {
			which = "r"
			if preNode.Left.Value == val {
				which = "l"
			}
			break
		}
		preNode = node
		if val <= node.Value && node.Left != nil {
			node = node.Left
		} else if val > node.Value && node.Right != nil {
			node = node.Right
		} else {
			break
		}
	}

	var newNode *Node
	if node.Right == nil && node.Left == nil {
		newNode = nil
	} else if node.Left != nil && node.Right == nil {
		newNode = node.Left
	} else if node.Left == nil && node.Right != nil {
		newNode = node.Right
	} else if node.Left != nil && node.Right != nil {
		newNode = node.Right
		var _jreNode *Node
		for newNode.Left != nil {
			_jreNode = newNode
			newNode = newNode.Left
		}
		if _jreNode != nil {
			_jreNode.Left = nil
			newNode.Right = node.Right
		}
		newNode.Left = node.Left
	}

	if which == "l" {
		preNode.Left = newNode
	} else {
		preNode.Right = newNode
	}
}
