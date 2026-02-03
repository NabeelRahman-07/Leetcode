func restoreString(s string, indices []int) string {
    n:=len(s)
    result:=make([]rune,n)
    char:=[]rune(s)
    for i:=0;i<n;i++{
        result[indices[i]]=char[i]
    }
    return string(result)
}