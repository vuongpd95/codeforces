package graphsproblem1549b

import "testing"

func TestMaxPawnsReachRow1_111_000(t *testing.T) {
    got := MaxPawnsReachRow1([]int{1, 1, 1}, []int{0, 0, 0})
    want := 3

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}


func TestMaxPawnsReachRow1_1111_1111(t *testing.T) {
    got := MaxPawnsReachRow1([]int{1, 1, 1, 1}, []int{1, 1, 1, 1})
    want := 4

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestMaxPawnsReachRow1_010_010(t *testing.T) {
    got := MaxPawnsReachRow1([]int{0, 1, 0}, []int{0, 1, 0})
    want := 0

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}


func TestMaxPawnsReachRow1_00000_11001(t *testing.T) {
    got := MaxPawnsReachRow1([]int{0, 0, 0, 0, 0}, []int{1, 1, 0, 0, 1})
    want := 0

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
