func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	l, r := 0, len(people)-1
	cnt := 0
	for l <= r {
		if l != r && people[l] + people[r] <= limit {
			l++
			r--
		} else {
			r--
		}
		cnt++
	}
	return cnt
}
