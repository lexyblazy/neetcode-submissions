func topKFrequent(nums []int, k int) []int {
  freqMap := make(map[int]int)
  result := []int{}

  for _, num := range nums{
    if _, ok := freqMap[num];ok {
        freqMap[num]++
    }else {
        freqMap[num] = 1
    }
  }

  fmt.Println(freqMap)

  seen := make(map[int]struct{})
  for i:=0; i < k; i++ {  
    mostFreq := 0
    value := 0
    for num, freq := range freqMap{
        // do a check here if the number has been seen, we skip
        if _, exists := seen[num]; exists {
            continue
        }
        if freq > mostFreq {
            mostFreq = freq
            value = num
        }
    }

    
    if _,ok := seen[value]; !ok {
        result = append(result,value)
    }
    seen[value] = struct{}{}
    
  }

  return result
}
