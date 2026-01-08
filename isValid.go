/***
  Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
 

Example 1:

Input: s = "()"

Output: true

Example 2:

Input: s = "()[]{}"

Output: true

Example 3:

Input: s = "(]"

Output: false

Example 4:

Input: s = "([])"

Output: true

Example 5:

Input: s = "([)]"

Output: false

 
**/
func isValid(s string) bool {
    st := []rune{}
    match := map[rune]rune{
        ')':'(',
        '}':'{',
        ']':'[',
    }

    for _,char := range s{
        if(char == '(' || char == '{' || char == '['){
            st = append(st, char)
        } else if(char == ')' || char == '}' || char == ']'){
            if len(st) == 0 || st[len(st)-1] != match[char] {
                return false
            }
            st = st[:len(st)-1]
        }
    }
    return len(st) == 0
}
