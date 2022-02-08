package datastruct

type List struct {
	head *Node
	size int
}

type Node struct {
	value interface{}
	next  *Node
}
