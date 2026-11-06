type LRUCache struct {
    capacity int
	keyToNode map[int]*Node
	sentinel *Node
}

type Node struct {
	k int
	v int
	next *Node
	prev *Node
}

func Constructor(capacity int) LRUCache {
	sentinel := &Node{}
	sentinel.next, sentinel.prev = sentinel, sentinel
    return LRUCache{
		capacity: capacity,
		keyToNode: make(map[int]*Node),
		sentinel: sentinel,
	}
}

func (c *LRUCache) Get(key int) int {
    if n, exists := c.keyToNode[key]; exists {
		remove(n)
		add(n, c.sentinel)
		return n.v
	}

	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if n, exists := c.keyToNode[key]; exists {
		n.v = value
		remove(n)
		add(n, c.sentinel)
		return
	}
    if len(c.keyToNode) == c.capacity {
		delete(c.keyToNode, c.sentinel.prev.k)
		remove(c.sentinel.prev)
	}
	n := &Node{k: key, v: value}
	add(n, c.sentinel)
	c.keyToNode[key] = n
}

func add(n *Node, sentinel *Node) {
	n.next, n.prev, sentinel.next, sentinel.next.prev = sentinel.next, sentinel, n, n
}

func remove(n *Node) {
	n.prev.next, n.next.prev = n.next, n.prev
}
