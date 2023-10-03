package graphsproblem755a

import (
	"errors"
	"math"
)

// https://codeforces.com/problemset/problem/755/A
func FindCounterExamplePolandBall(n int) (int, error) {
    result := -1

    if n < 1 || n > 1000 {
        return result, errors.New(
            "Error! n should be >= 1 and <= 1000",
        )
    }

    for m := 0; m <= 1000; m++ {
        a := m * n + 1
        if IsPrime(a) == false {
            result = m
            break
        }
    }

    if result == -1 {
        return result, errors.New("Error! No counter example found")
    }

    return result, nil
}

// https://www.scaler.com/topics/data-structures/primality-test/
func IsPrime(a int) bool {
    if a <= 3 {
        return true
    }

    result := true

    for i := 2.0; i <= math.Sqrt(float64(a)); i++ {
       if math.Mod(float64(a), i) == 0.0 {
           result = false
       }
    }

    return result
}
