func maxArea(heights []int) int {
  var maxArea int

  left := 0
  right := len(heights) -1

  for left < right {
	 width := right - left
	 height := min(heights[right], heights[left])

	 area := width * height

	 if area > maxArea {
		maxArea = area
	 }

	 if heights[left] < heights[right]{
		left++
	 }else{
		right--
	 }
	

  }


  return maxArea
}
