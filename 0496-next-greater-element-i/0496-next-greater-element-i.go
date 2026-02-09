func nextGreaterElement(nums1 []int, nums2 []int) []int {
    var stack []int
    for _,val1 := range nums1{
        elem := -1
        pos := false
        for j :=0 ; j<len(nums2); j++{
            if nums2[j]==val1{
                pos = true
                continue
            }
            if nums2[j]>val1 && elem==-1 && pos{
                elem = nums2[j]
                break
            }
        }
        stack = append(stack, elem)
    }
    return stack
}