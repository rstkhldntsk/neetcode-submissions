func majorityElement(nums []int) []int {
	num1, num2 := -1, -1
	cnt1, cnt2 := 0, 0

	for _, num := range nums {
		if num == num1 {
			cnt1++
		} else if num == num2 {
			cnt2++
		} else if cnt1 == 0 {
			cnt1 = 1
			num1 = num
		} else if cnt2 == 0 {
			cnt2 = 1
			num2 = num
		} else {
			cnt1--
			cnt2--
		}
	}

	cnt1, cnt2 = 0, 0
	for _, num := range nums {
		if num == num1 {
			cnt1++
		} else if num == num2 {
			cnt2++
		}
	}

	res := make([]int, 0)
	if cnt1 > len(nums)/3 {
		res = append(res, num1)
	}
	if cnt2 > len(nums)/3 {
		res = append(res, num2)
	}

	return res
}
