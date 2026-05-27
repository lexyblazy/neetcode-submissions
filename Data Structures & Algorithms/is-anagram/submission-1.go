func toHashMap(word string) map[rune]int{
    hashmap := make(map[rune]int)
	for _, letter := range word {
		_, exists := hashmap[letter]
		if exists {
			hashmap[letter] = hashmap[letter] + 1
		}else {
			hashmap[letter] = 1
		}
	}
	return hashmap
}

func isAnagram(s string, t string) bool {

// both strings must be of the same length
	if len(s) != len(t){
		return false
	}
	map1 := toHashMap(s)
	map2 := toHashMap(t)

	for key,value := range map1 {
		if value != map2[key]{
			return false
		}
	}

	return true
}
