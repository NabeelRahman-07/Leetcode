func maxArea(height []int) int {
    p1 := 0
    p2 := len(height)-1
    var result int
    for p1<p2 {
        var water int
        if height[p1] <= height[p2] {
            water = height[p1] * (p2 - p1)
            p1++
        }else{
            water = height[p2] * (p2 - p1)
            p2--
        }
        if result < water {
            result = water
        }
    }
    return result
}