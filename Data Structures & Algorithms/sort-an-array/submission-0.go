func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l, r int) {
	for l <= r {
		lt, gt := partition(nums, l, r)
		if lt-l < r-gt {
			quickSort(nums, l, lt-1)
			l = gt+1
		} else {
			quickSort(nums, gt+1, r)
			r = lt-1
		}
	}
}

func partition(nums []int, l, r int) (int, int) {
	pivotIdx := l + rand.Intn(r-l+1)
	pivot := nums[pivotIdx]
	nums[l], nums[pivotIdx] = nums[pivotIdx], nums[l]

	lt := l
	i := l+1
	gt := r
	for i <= gt {
		if nums[i] < pivot {
			nums[lt], nums[i] = nums[i], nums[lt]
			i++
			lt++
		} else if nums[i] > pivot {
			nums[gt], nums[i] = nums[i], nums[gt]
			gt--
		} else {
			i++
		}
	}
	return lt, gt
}

