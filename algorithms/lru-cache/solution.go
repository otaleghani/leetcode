package main

type LRUCache struct {
	hash map[int]*Node
	cap  int
	head *Node
	tail *Node
}

type Node struct {
	val  int
	key  int
	next *Node
	prev *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		hash: make(map[int]*Node),
		cap:  capacity,
		head: &Node{},
		tail: &Node{},
	}

	cache.head.next = cache.tail
	cache.tail.prev = cache.head

	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.hash[key]; exists {
		this.removeNode(node)
		this.addToFront(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.hash[key]; exists {
		node.val = value
		this.removeNode(node)
		this.addToFront(node)
		return
	}

	if len(this.hash) == this.cap {
		lruNode := this.tail.prev
		this.removeNode(lruNode)
		delete(this.hash, lruNode.key)
	}

	newNode := &Node{key: key, val: value}
	this.addToFront(newNode)
	this.hash[key] = newNode
}

func (this *LRUCache) addToFront(node *Node) {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
