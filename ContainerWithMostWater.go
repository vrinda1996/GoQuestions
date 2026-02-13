func maxArea(height []int) int {
    left , right := 0, len(height) - 1
    area := 0
    for left < right {
        ht := right - left
        minHeight := min(height[right], height[left])
        if area < ht * minHeight {
            area = ht * minHeight
        }
        if height[right] < height[left] {
            right--
        } else {
            left++
        }
    }
    return area
}

func min(a, b int) int {
    if a< b{
        return a
    }
    return b
}
