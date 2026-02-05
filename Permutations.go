func permute(nums []int) [][]int {
    res := [][]int{}
    backtrack(nums, &res, []int{})
    return res
}

func backtrack(nums []int,res *[][]int, temp []int){
    if len(temp) == len(nums){
        combo := make([]int, len(temp))
        copy(combo, temp)
        *res = append(*res, combo)
        return 
    }
    for i:=0; i< len(nums); i++{
        if slices.Contains(temp, nums[i]){
            continue
        }
        temp = append(temp, nums[i])
        backtrack(nums, res, temp)
        temp = temp[:len(temp)-1]
    }
}
