package easy

func TwoSum(nums []int, target int) []int {
	indexes := make(map[int]int)
	for i := range nums {
		remainder := target - nums[i]
		//check if remainder is already in indexes
		j, ok := indexes[remainder]
		if ok {
			//found the other half already, return
			return []int{i, j}
		}
		indexes[nums[i]] = i
	}
	return nil
}
