func sortColors(nums []int)  {
    if len(nums) == 1 {
        return
    }
    zero, one := 0, 0
    for _, val := range nums {
        if val == 1 {
            one++
        }
        if val == 0 {
            zero++
        }
    }
    for i := 0 ; i < len(nums) ; i++ {
        if zero > 0 {
            nums[i] = 0
            zero--
            continue
        }
        if one > 0 {
            nums[i] = 1
            one--
            continue
        }
        nums[i] = 2
    }


}