package container

import (
	"github.com/znkisoft/zedisDB/database/datastruct"
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

type List struct {
	/*
		AMEND: ziplist and dblinkedlist
		- under what circumstances will the ziplist be converted to a dblinkedlist?
		- - when list length is greater than 512
		- - when list size is greater than 64bytes
	*/
	data datastruct.List
}

type direction int

const (
	Head direction = iota
	Tail
)
