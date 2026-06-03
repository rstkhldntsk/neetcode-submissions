func numRescueBoats(people []int, limit int) int {
    sort.Ints(people)
    res, l, r := 0, 0, len(people)-1
    for l <= r {
        remain := limit - people[r]
        r--
        res++
        if remain >= people[l] {
            l++
        }
    }
    return res
}