func isPalindrome(x int) bool {
    if x < 0 { return false }

    s := strconv.Itoa(x)

    runes := []rune(s)

    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        if runes[i] != runes[j] {
            return false
        }
    }
    return true
}
