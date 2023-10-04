package graphsproblem1797a

// https://codeforces.com/problemset/problem/1797/A
func FindLeastObstacles(n int, m int, x1 int, y1 int, x2 int, y2 int) int {
    startCellObs := findLeastObstaclesOfCell(n, m, x1, y1)
    endCellObs := findLeastObstaclesOfCell(n, m, x2, y2)

    if startCellObs >= endCellObs {
        return endCellObs
    } else {
        return startCellObs
    }
}

func findLeastObstaclesOfCell(n int, m int, x int, y int) int {
    top_x := x - 1
    top_y := y

    bot_x := x + 1
    bot_y := y

    left_x := x
    left_y := y - 1

    right_x := x
    right_y := y + 1

    result := 0

    if cellInsideMaze(n, m, top_x, top_y) {
        result = result + 1
    }

    if cellInsideMaze(n, m, bot_x, bot_y) {
        result = result + 1
    }

    if cellInsideMaze(n, m, left_x, left_y) {
        result = result + 1
    }

    if cellInsideMaze(n, m, right_x, right_y) {
        result = result + 1
    }

    return result
}

func cellInsideMaze(n int, m int, x int, y int) bool {
    return (x <= n && y <= m && x >= 1 && y >= 1)
}
