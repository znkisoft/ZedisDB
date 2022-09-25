package container

import (
	"container/list"
)

// lpush key value1 [value2] 将一个或多个值插入到列表头部
// lpushx 将一个或多个值插入到已存在的列表头部
// rpush 在列表中添加一个或多个值
// rpushx 为已存在的列表添加值
// lpop 移出并获取列表的第一个元素
// rpop 移除并获取列表最后一个元素
// rpoplpush source destination 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
// lrem  key count value 移除列表元素
// llen 获取列表长度
// lindex 通过索引获取列表中的元素
// lset key index value 通过索引设置列表元素的值
// lrange 获取列表指定范围内的元素
// LINSERT key BEFORE|AFTER pivot value 在列表的元素前或者后插入元素

// List
//
// # Redis prototypes reference:
// ```c
// list *listCreate(void);
// void listRelease(list *list);
// list *listAddNodeHead(list *list, void *value);
// list *listAddNodeTail(list *list, void *value);
// list *listInsertNode(list *list, listNode *old_node, void *value, int after);
// void listDelNode(list *list, listNode *node);
// listIter *listGetIterator(list *list, int direction);
// listNode *listNext(listIter *iter);
// void listReleaseIterator(listIter *iter);
// list *listDup(list *orig);
// listNode *listSearchKey(list *list, void *key);
// listNode *listIndex(list *list, long index);
// void listRewind(list *list, listIter *li);
// void listRewindTail(list *list, listIter *li);
// void listRotate(list *list);
// ```
type List struct {
	/*
		AMEND: ziplist and dblinkedlist
		- under what circumstances will the ziplist be converted to a dblinkedlist?
		- - when list length is greater than 512
		- - when list size is greater than 64bytes
	*/
	li *list.List // TODO add ziplist implementation
}

type direction int

const (
	Head direction = iota
	Tail
)

func NewList() *List {
	return &List{
		list.New(),
	}
}

// void listRelease(list *list);
// list *listAddNodeHead(list *list, void *value);
// list *listAddNodeTail(list *list, void *value);
// list *listInsertNode(list *list, listNode *old_node, void *value, int after);
// void listDelNode(list *list, listNode *node);
// listNode *listNext(listIter *iter);
// listNode *listSearchKey(list *list, void *key);
// listNode *listIndex(list *list, long index);
// void listRotate(list *list);

// listIter *listGetIterator(list *list, int direction);
// void listReleaseIterator(listIter *iter);
// list *listDup(list *orig);
// void listRewind(list *list, listIter *li);
// void listRewindTail(list *list, listIter *li);
