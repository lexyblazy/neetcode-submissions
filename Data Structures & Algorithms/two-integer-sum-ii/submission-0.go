func twoSum(numbers []int, target int) []int {
  left := 0
  right := len(numbers)-1

  result := []int{}

  for left < right {
	   sum := numbers[left] + numbers[right]
	 
	  if sum == target {
		  result = append(result, left+1, right+1)
	     break;
	  }

	  if sum > target {
		 right--
	  } else if sum < target {
		left++
	  }
  }
  
  return result;

}
