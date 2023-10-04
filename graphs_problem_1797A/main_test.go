package graphsproblem1797a

import "testing"

func TestFindLeastObstacles_442233(t *testing.T) {
    got := FindLeastObstacles(4, 4, 2, 2, 3, 3)

    if got != 4 {
        t.Errorf("got %d, wanted 4", got)
    }
}

func TestFindLeastObstacles_671123(t *testing.T) {
    got := FindLeastObstacles(6, 7, 1, 1, 2, 3)

    if got != 2 {
        t.Errorf("got %d, wanted 2", got)
    }
}

func TestFindLeastObstacles_995136(t *testing.T) {
    got := FindLeastObstacles(9, 9, 5, 1, 3, 6)

    if got != 3 {
        t.Errorf("got %d, wanted 3", got)
    }
}
