func merge(nums1 []int, m int, nums2 []int, n int)  {
    j := n-1
    i := m-1
    p := (m+n)-1
    for j>=0{
        if i>=0 && nums1[i]>nums2[j]{
            nums1[p]=nums1[i]
            i--
        }else{
            nums1[p]=nums2[j]
            j--
        }
        p--
    }
}