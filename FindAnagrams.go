func findAnagrams(s string, p string) []int {
    if len(p) > len(s) {
        return nil
    }
    var window, pCount [26]int
    var res []int
    for i:=0; i<len(p) ; i++ {
        pCount[p[i] - 'a']++
    }
    left := 0
    for right :=0; right< len(s); right++ {
        window[s[right] - 'a']++

        if right - left + 1 > len(p){
            window[s[left] - 'a']--
            left++
        }
        if window == pCount{
            res = append(res, left)
        }
    }
    return res
}
