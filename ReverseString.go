func reverseString(s []byte)  {
    n := len(s)-1
    for i:=0 ; i<= n/2; i++ {
        s[i], s[n-i] = s[n-i], s[i]
    }
}

func main(){
  s:= []byte{"h","e","l","l","o"}
  fmt.Println(reverseString(s))
}
