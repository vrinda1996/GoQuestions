func longestPalindrome(s string) string {
    maxStr := ""
    if len(s) <=1 {
        return s
    }
    for i:=0; i< len(s)-1;i++ {
        odd, even := expand(s, i, i), expand(s, i, i+1)
        if len(maxStr) < len(odd) {
            maxStr = odd
        }
        if len(maxStr) < len(even) {
            maxStr = even
        }
    }
    return maxStr
}

func expand(s string, left, right int) string {
    for left >=0 && right < len(s) && s[left] == s[right]{
        left--
        right++
    }
    return s[left+1:right]
}
