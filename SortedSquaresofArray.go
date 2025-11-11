func sortedSquares(nums []int) []int {
    res := make([]int, 0, len(nums))
    start, end := 0, len(nums)-1
    for start<=end{
        if(nums[start] * nums[start] < nums[end] * nums[end]){
            res = append(res, nums[end] * nums[end])
            end--
        } else{
            res = append(res, nums[start] * nums[start])
            start++
        }
    }
    reverse(res)
    return res
}

func reverse(nums []int){
    for i,j:= 0, len(nums)-1; i<j; i, j = i+1, j-1{
        nums[i], nums[j] = nums[j], nums[i]
    }
}
