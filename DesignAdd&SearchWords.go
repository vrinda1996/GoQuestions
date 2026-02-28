type Trie struct{
    children [26]*Trie
    isWord bool
}
func NewTrie() *Trie{
    return &Trie{
        children : [26]*Trie{},
        isWord : false,
    }
}
type WordDictionary struct {
    root *Trie
}


func Constructor() WordDictionary {
    return WordDictionary{
        root : NewTrie(),
    }
}


func (this *WordDictionary) AddWord(word string)  {
    curr := this.root
    for _, c := range word {
        index := c-'a'
        if curr.children[index] == nil {
            curr.children[index] = NewTrie()
        }
        curr = curr.children[index]
    }
    curr.isWord = true
}


func (this *WordDictionary) Search(word string) bool {
    return dfs(word, this.root, 0)
}

func dfs(word string, root *Trie, index int) bool {
    if index == len(word){
        return root.isWord
    }
    c := word[index]
    if c == '.'{
        for _, child := range root.children{
            if child!= nil && dfs(word, child, index+1){
                return true
            }
        }
        return false
    }
    next := root.children[c - 'a']
    if next == nil {
        return false
    }

    return dfs(word, next, index+1)
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
