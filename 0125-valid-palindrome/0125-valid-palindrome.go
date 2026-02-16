func isPalindrome(s string) bool {
	x := []rune(s)
    p1 := 0
	p2 := len(x) - 1
	for p1 < p2 {
		if unicode.IsLetter(x[p1]) || unicode.IsDigit(x[p1]) {
			if unicode.IsLetter(x[p2]) || unicode.IsDigit(x[p2]) {
				if unicode.ToLower(x[p1]) != unicode.ToLower(x[p2]) {
					return false
				}else{
                    p1++
                    p2--
                }
			}else{
                p2--
            }
		}else{
            p1++
        }  
	}
    return true
}