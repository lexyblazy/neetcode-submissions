
func isAnagram(s string, t string) bool {

	// both must be of the length
	if len(s) != len(t){
		return false
	}
 // since we are dealing with alphabets, we know the boundary for letters are 26
	letters := make([]int,26)

	
	for i:=0; i < len(s); i++ {
		// increment and decrement for each word. if they are truly anagrams,
		// the resultant will be zero
		letters[s[i]- 'a']++
		letters[t[i]- 'a']--
	}

	for _, val := range letters {
		if val !=0{
			return false
		}
	}
	return true
 
}
