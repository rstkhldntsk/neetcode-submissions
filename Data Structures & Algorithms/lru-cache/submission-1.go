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
		c.add(node)
		return node.val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.m[key]; ok {
		c.remove(node)
		c.add(node)
		node.val = value
		return
	}
	if len(c.m) == c.capacity {
		delete(c.m, c.sentinel.prev.key)
		c.remove(c.sentinel.prev)
	}
	n := &Node{key: key, val: value}
	c.add(n)
	c.m[key] = n
}

func (c *LRUCache) remove(n *Node) {
	n.prev.next, n.next.prev = n.next, n.prev
}

func (c *LRUCache) add(n *Node) {
	c.sentinel.next, n.next, n.prev, c.sentinel.next.prev = n, c.sentinel.next, c.sentinel, n
}
