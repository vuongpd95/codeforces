package graphsproblem393a

// https://codeforces.com/problemset/problem/939/A
func FindLoveTriangle(arr []int) bool {
    l := len(arr)
    result := false

    for i := 0; i < l; i++ {
        edge1 := arr[i]

        if edge1 < 1 || edge1 > l { continue }

        edge2 := arr[edge1 - 1]

        if edge2 < 1 || edge2 > l { continue }

        if arr[edge2 - 1] == i + 1 {
            result = true
            break
        }
    }

    return result
}
