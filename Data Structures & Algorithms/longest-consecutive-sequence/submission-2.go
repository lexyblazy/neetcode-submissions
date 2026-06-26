func longestConsecutive(nums []int) int {
  numSet := make(map[int]bool)

  for _,num := range nums {
    numSet[num] = true
  }

  longest := 0

   for num := range numSet {

    // if the val-1 does not exist, it is the start of the new sequence.
     if  _, exists := numSet[num - 1]; !exists {
            length := 1
            
            // continue to increment the value to check if the next sequence 
            // exists in the array, if exists increase the length
            for numSet[num + length] {
                length++
            }

            if length > longest {
             longest = length
            }
     }

   

   }

   return longest;

}
