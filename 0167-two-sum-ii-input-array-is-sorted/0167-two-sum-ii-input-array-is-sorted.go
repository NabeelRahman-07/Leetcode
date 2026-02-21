func twoSum(numbers []int, target int) []int {
    p1 := 0
    p2 := len(numbers)-1

    for p1 < p2 {
        sum := numbers[p1]+numbers[p2]
        if sum < target {
            p1++
        }else if sum > target {
            p2--
        }else{
            return []int{p1+1, p2+1}
        }
    }
    return []int{}
}