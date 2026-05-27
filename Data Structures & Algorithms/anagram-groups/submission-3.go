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
