package main

import (
	"exercises/ds"
	"fmt"
)

type Cache interface {
	Set(key string, val int)
	Get(key string) (int, bool)
	Delete(key string)
}

type cache struct {
	cacheRecords map[string]*ds.LinkedListNode[int]
	limit        int
	size         int
	dll          ds.DoubleLL[int]
}

func (c *cache) Set(key string, val int) {
	if c.size < c.limit {
		c.dll.AddNodeAtHead(val)
		c.cacheRecords[key] = c.dll.Head()
		c.size++
	} else {
		c.dll.RemoveNodeAtTail()
		// todo: remove map key also
		c.dll.AddNodeAtHead(val)
		c.cacheRecords[key] = c.dll.Head()
	}
}

func (c *cache) Get(key string) (int, bool) {
	node, ok := c.cacheRecords[key]
	if ok {
		c.dll.MoveNodeToHead(node)
		val := node.Value
		return val, ok
	}
	return 0, false
}

func (c *cache) Delete(key string) {
	d, ok := c.cacheRecords[key]
	if ok {
		c.dll.RemoveNode(d)
	}
	delete(c.cacheRecords, key)
	c.size--
}

func NewCache(limit int) Cache {
	dll := ds.NewDoubleLL[int]()
	return &cache{cacheRecords: make(map[string]*ds.LinkedListNode[int]), limit: limit, dll: dll}
}

//insert - (set)
//access - (get, ispresent)
//delete - remove/del

func main() {
	testCache := NewCache(5)
	testCache.Set("1", 1)
	testCache.Set("2", 2)
	testCache.Set("3", 3)
	testCache.Set("4", 4)
	testCache.Set("5", 5)

	value, isPresent := testCache.Get("3")
	fmt.Println(value, isPresent)

	testCache.Delete("3")
	testCache.Delete("3")
	value, isPresent = testCache.Get("3")
	fmt.Println(value, isPresent)

	testCache.Set("3", 3)
	testCache.Set("6", 6)
	//value, isPresent = testCache.Get("1")
	//fmt.Println(value, isPresent)
}
