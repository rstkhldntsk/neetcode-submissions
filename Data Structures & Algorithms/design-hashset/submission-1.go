type MyHashSet struct {
	set []int
}

func Constructor() MyHashSet {
	// key is in the range [0, 1000000]
    // 31251 * 32 = 1000032
    return MyHashSet{set: make([]int, 31251)}
}

func (this *MyHashSet) getMask(key int) int {
    return 1 << (key % 32)
}

func (this *MyHashSet) Add(key int) {
    this.set[key/32] |= this.getMask(key)
}

func (this *MyHashSet) Remove(key int) {
	if this.Contains(key) {
        this.set[key/32] ^= this.getMask(key)
    }
}

func (this *MyHashSet) Contains(key int) bool {
    return (this.set[key/32] & this.getMask(key)) != 0
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 