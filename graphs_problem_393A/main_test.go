package graphsproblem393a

import "testing"

func TestFindLoveTriangle_24513(t *testing.T) {
    got := FindLoveTriangle([]int{2, 4, 5, 1, 3})
    want := true

    if got != want {
        t.Errorf("got %t, wanted %t", got, want)
    }
}


func TestFindLoveTriangle_5551(t *testing.T) {
    got := FindLoveTriangle([]int{5, 5, 5, 1})
    want := false

    if got != want {
        t.Errorf("got %t, wanted %t", got, want)
    }
}
