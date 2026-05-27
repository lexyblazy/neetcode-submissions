func toMap(nums []int) map[int]int {
	seen := make(map[int]int)
	for index, num := range nums{
		seen[num] = index
	}
	return seen
}

func twoSum(nums []int, target int) []int {
    hashMap := toMap(nums)
	result := []int{}

	for i, num := range nums {
		difference := target - num
		j, exists := hashMap[difference]

		if i != j && exists {
			result = append(result,i,j)
			break
		}

	}
	return  result
}
