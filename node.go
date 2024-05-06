package keren

type Node struct {
	ID       string
	Root     *Root
	Element  *Element
	Parent   *Node
	Children []*Node
}

func NewNode(element *Element) *Node {
	node := &Node{
		ID:       Identifier(),
		Element:  element,
		Children: []*Node{},
	}
	if element.Children != nil {
		for _, child := range element.Children {
			node.Add(NewNode(child))
		}
	}
	return node
}
func (node *Node) Add(child *Node) {
	node.Children = append(node.Children, child)
}
func (node *Node) Adds(childs ...*Node) {
	for _, child := range childs {
		node.Add(child)
	}
}
func (node *Node) Remove(child *Node) {
	for i, n := range node.Children {
		if n == child {
			node.Children = append(node.Children[:i], node.Children[i+1:]...)
		}
	}
}
func (node *Node) Append(element ...*Element) *Node {
	for _, elem := range element {
		child := NewNode(elem)
		node.Add(child)
	}
	return node
}
