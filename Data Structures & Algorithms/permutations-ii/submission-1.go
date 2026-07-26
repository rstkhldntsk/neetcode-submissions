import "slices"

func permuteUnique(nums []int) [][]int {
    counter := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        counter[nums[i]]++
    }

    result := make([][]int, 0)
    var backtrack func(prefix []int)
    backtrack = func(prefix []int) {
        if len(prefix) == len(nums) {
            result = append(result, slices.Clone(prefix))
            return
        }
        for num, count := range counter {
            if count == 0 {
                continue
            }
            counter[num]--
            backtrack(append(prefix, num))
            counter[num]++
        }
    }
    backtrack([]int{})
    return result
}