func myFunc(a []int) []int {
    b := make([]int, 0)
    for _, v := range a {
        if v != 0 {
            b = append(b, v)
        }
    }
    return b
}

//Alternative solution using a for loop
func myFuncAlt(a []int) []int {
    result := []int{}
    for i := 0; i < len(a); i++ {
        if a[i] != 0 {
            result = append(result, a[i])
        }
    }
    return result
} 