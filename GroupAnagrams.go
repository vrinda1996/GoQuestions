func groupAnagrams(strs []string) [][]string {
    sortedMap := make(map[string][]string)
    var result [][]string
    for _, str := range strs {
        key := sorting(str)
        sortedMap[key] = append(sortedMap[key], str)
    }

    for _, val := range sortedMap{
        result = append(result, val)
    }
    return result

}

func sorting(str string) string{
    runes := []rune(str)
    sort.Slice(runes, func(i, j int) bool {
        return runes[i] < runes[j] // Sort in ascending order
    })
    return string(runes)
}
