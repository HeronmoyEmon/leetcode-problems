// Problem 69: Sqrt(x)
// https://leetcode.com/problems/sqrtx/
//
// Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
// The returned integer should be non-negative as well.
//
// You must not use any built-in exponent function or operator.
// For example, do not use pow(x, 0.5) in code.
//
// Example 1:
// Input: x = 4
// Output: 2
//
// Example 2:
// Input: x = 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since we want the floor of the square root, we return 2.

package main

import "fmt"

func findSqrt(x int) int {
    if x < 2 {
        return x
    }

    min := 1
    max := x / 2
    var result int

    for min <= max {
        mid := (min + max) / 2
        square := mid * mid

        if square == x {
            return mid
        } else if square < x {
            result = mid
            min = mid + 1
        } else {
            max = mid - 1
        }
    }

    return result
}

func main() {
    fmt.Println(findSqrt(4))  // Output: 2
    fmt.Println(findSqrt(8))  // Output: 2
}
