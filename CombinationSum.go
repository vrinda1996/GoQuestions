func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    return calculateSum(candidates, 0 , target, [][]int{}, []int{})
}

func calculateSum(nums []int, index, target int, res [][]int, temp []int) [][]int{
    if target == 0 {
        combo := make([]int, len(temp)) // copy is needed here inorder to store temp else it will be overwritten with backtracking approach
        copy(combo, temp)  
        return append(res, combo)
    }
    if target < 0{
        return res
    }
    
    for ;index< len(nums); index++ {
        if nums[index] > target{
            break
        }
        temp = append(temp, nums[index])
        res = calculateSum(nums, index, target - nums[index], res, temp)
        temp = temp[:len(temp)-1]
    } 
    return res
    
}
