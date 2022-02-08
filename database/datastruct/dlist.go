package datastruct

import (
	stlList "container/list"
)

// DList
// length
// first
// last
// insertNode
// create
// prevNode
// nextNode
// nodeValue
// deleteNode
// addNodeHead
// addNodeTail
// searchKey
// index
// rotate
// dup
// release o(n)
type DList struct {
	stlList.List
}
