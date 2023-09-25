package graphsproblem1549b

// https://codeforces.com/problemset/problem/1549/B
func MaxPawnsReachRow1(rown []int, row1 []int) int {
    result := 0
    i := 0
    l := len(rown)

    connected_triangle_hole_count := 0
    last_connected_right_i := -1

    for i <= (l - 1) {
        hole := row1[i]
        center := rown[i]
        left := -1
        right := -1
        
        if i - 1 >= 0 {
            left = rown[i - 1]
        }

        if i + 1 <= (l - 1) {
            right = rown[i + 1]
        }

        if (hole == 0 && center == 0) || (hole == 1 && left <= 0 && right <= 0) {
            i += 1
            continue
        }

        if hole == 0 && center == 1 {
            row1[i] = 1
            rown[i] = 0
            result += 1
            i += 1
            continue
        }

        if hole == 1 && left == 1 && right <= 0 {
            row1[i] = 1
            rown[i - 1] = 0
            result += 1
            i -= 2
            if i < 0 {
                i = 0
            }
            continue
        }

        if hole == 1 && left <= 0 && right == 1 {
            row1[i] = 1
            rown[i + 1] = 0
            result += 1
            i += 1
            continue
        }        

        if last_connected_right_i == -1 {
            last_connected_right_i = i + 1
            connected_triangle_hole_count += 1
            i += 1
            continue
        }

        if last_connected_right_i == (i - 1) {
            last_connected_right_i = i + 1
            connected_triangle_hole_count += 1
            i += 1
            continue
        } else {
            result += connected_triangle_hole_count
            connected_triangle_hole_count = 1
            last_connected_right_i = i + 1
            i += 1
            continue
        }
    }

    return result
}

