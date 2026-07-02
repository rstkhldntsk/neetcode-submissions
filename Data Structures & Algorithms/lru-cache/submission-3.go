type LRUCache struct {
	capacity int
	m        map[int]*Node
	sentinel *Node
}

type Node struct {
	key, val   int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	sentinel := &Node{}
	sentinel.next, sentinel.prev = sentinel, sentinel
	c := LRUCache{
		sentinel: sentinel,
		m:        make(map[int]*Node),
		capacity: capacity,
	}
	
	return c
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.m[key]; ok {
		c.remove(node)
		c.placeAsMRU(node)
		return node.val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.m[key]; ok {
		node.val = value
		c.remove(node)
		c.placeAsMRU(node)
		return
	}
	if len(c.m) == c.capacity {
		k := c.sentinel.prev.key
		c.remove(c.sentinel.prev)
		delete(c.m, k)
	}
	n := &Node{key: key, val: value}
	c.m[key] = n
	c.placeAsMRU(n)
}

func (c *LRUCache) placeAsMRU(node *Node) {
	node.prev = c.sentinel
	node.next = c.sentinel.next
	c.sentinel.next.prev = node
	c.sentinel.next = node
}

func (c *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
