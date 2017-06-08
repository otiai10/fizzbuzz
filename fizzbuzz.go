package fizzbuzz

import "fmt"

func GetOutput(i int) string {
    fizz := i % 3
    buzz := i % 5
    switch {
        case fizz == 0 && buzz == 0:
            return "FizzBuzz"
        case fizz == 0:
            return "fizz"
        case buzz == 0:
            return "buzz"
        default:
            return fmt.Sprintf("%v", i)
    }
}
