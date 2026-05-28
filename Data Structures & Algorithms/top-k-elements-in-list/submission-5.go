func topKFrequent(nums []int, k int) []int {
    // 1. convert to nums to a map
    freqMap := make(map[int]int)
    for _, num := range nums {
        freqMap[num]++
    }

    // 2. organize frequency in buckets
    buckets := make([][]int, len(nums) + 1)

    for num, freq := range freqMap {
        buckets[freq] = append(buckets[freq], num)
    }

    // 3. loop from the highest index till len(results) == k
   result := make([]int,0,k)
    for freq:= len(buckets) -1;freq > 0; freq--{
        for _, num := range buckets[freq] {
            result = append(result,num)

            if len(result) == k {
                return result
            }
        }
    }

    return result
}
