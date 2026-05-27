func toMap(str string) map[rune]int{
    m := make(map[rune]int)
    for _,s := range str {
       _, ok := m[s]
       if ok {
        m[s] = m[s] + 1
       }else{
        m[s] = 1
       }
    }
    return m
}

func isAnagram(str1 string, str2 string) bool {

   if len(str1) != len(str2){
    return false
   }

   map1 := toMap(str1)
   map2 := toMap(str2)

   for val, count := range map1 {
     if count != map2[val]{
        return false
     }
   }

   return true
}

func groupAnagrams(strs []string) [][]string {

   groups := make(map[string][]string)
   results := [][]string{}


  for i:=0; i < len(strs); i++ {
    ang := []string{}
    for j:= 0; j < len(strs); j++ {
        if isAnagram(strs[i], strs[j]) {
            ang = append(ang,strs[j])
        }
    }
    if len(ang) > 0 {
        groups[strs[i]] = ang
    }else{
        groups[strs[i]] = []string{strs[i]}
    }
  }
  

    seen := make(map[string]bool)

    for word, anagrams := range groups {
        if _, ok := seen[word]; ok {
            continue
        }
        seen[word] = true

        for _,anagram := range anagrams {
           if _, ok := seen[anagram]; ok {
            continue
          }
            seen[anagram] = true   
        }

        results = append(results,anagrams)
        
    }
  

  return results

}
