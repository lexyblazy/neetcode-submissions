type Word struct {
    freq [26]int
    length int
}

func NewWord(str string) Word {
    freq := [26]int{}
    for i := 0; i < len(str); i++{
        freq[str[i] - 'a']++
    }
    return Word{
        freq: freq,
        length: len(str),
    }
}

func isAnagram(str1 string, str2 string) bool {
    word1 := NewWord(str1)
    word2 := NewWord(str2)

    if word1.length != word2.length {
        return false
    }

    for i := 0; i < 26; i++{
        if word1.freq[i] != word2.freq[i] {
            return false
        }
    }
    return true
}

func groupAnagrams(strs []string) [][]string {

    group := make(map[Word][]string)
    result := [][]string{}

    for _, str := range strs {
        word := NewWord(str)

        if _,ok := group[word];ok {
            group[word] = append(group[word], str)
        }else {
            group[word] = []string{str}
        }
    }

    for _,value :=range group{
        result = append(result,value)
    }

    fmt.Println(result)

    return result;

  

}
