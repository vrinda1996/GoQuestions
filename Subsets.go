func subsets(nums []int) [][]int {
    var res [][]int
    backtracking(nums, []int{}, 0, &res)
    return res
}

func backtracking(nums, temp []int, index int, res *[][]int){
    temp2 := make([]int, len(temp))
    copy(temp2, temp)
    *res = append(*res, temp2)
    
    for i:= index; i< len(nums); i++{
        temp = append(temp, nums[i])
        backtracking(nums, temp, i +1, res)
        temp = temp[:len(temp)-1]
    }
}
