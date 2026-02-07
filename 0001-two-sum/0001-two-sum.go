func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for ind, val := range nums{
        complement := target-val
        if id, ok := m[complement] ; ok{
            return []int{id,ind}
        }
        m[val] = ind
    }
    return nil
}