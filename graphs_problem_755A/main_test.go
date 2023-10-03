package graphsproblem755a

import "testing"

func TestFindCounterExamplePolandBall_3_4(t *testing.T) {
    got, err := FindCounterExamplePolandBall(3)
    got1, err1 := FindCounterExamplePolandBall(4)


    if got != 1 || got1 != 2 {
        t.Errorf(
            "got %d, %d, wanted 1, 2. Errors: err: %t, err1: %t",
            got, got1, err, err1,
        )
    }
}

func TestIsPrime_0123(t *testing.T) {
    got := IsPrime(0)
    got1 := IsPrime(1)
    got2 := IsPrime(2)
    got3 := IsPrime(3)

    if got != true || got1 != true || got2 != true || got3 != true {
        t.Errorf(
            "got %t, %t, %t, %t, wanted true, true, true. true",
            got, got1, got2, got3,
        )
    }
}

func TestIsPrime_17(t *testing.T) {
    got := IsPrime(17)

    if got != true {
        t.Errorf(
            "got %t, wanted true",
            got,
        )
    }
}
