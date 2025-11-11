func intersection(nums1 []int, nums2 []int) []int {
    m := make(map[int]int, len(nums1))
    var res []int
    for _, num := range nums1{
        m[num] =1
    }
    
    for _, num := range nums2{
        if _,ok := m[num]; ok && !slices.Contains(res, num){
            res = append(res, num)
        }
    }
    return res
}
