package http

import (
	"fmt"
	"math"
)

func sum(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func modulus(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a % b, nil
}

func power(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func squareRoot(a int) (int, error) {
	if a < 0 {
		return 0, fmt.Errorf("negative number")
	}
	return int(math.Sqrt(float64(a))), nil
}

func factorial(a int) (int, error) {
	if a < 0 {
		return 0, fmt.Errorf("negative number")
	}
	if a == 0 {
		return 1, nil
	}
	result := 1
	for i := 1; i <= a; i++ {
		result *= i
	}
	return result, nil
}

func fibonacci(n int) ([]int, error) {
	if n < 0 {
		return nil, fmt.Errorf("negative number")
	}
	fib := make([]int, n)
	if n > 0 {
		fib[0] = 0
	}
	if n > 1 {
		fib[1] = 1
	}
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib, nil
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeFactors(n int) []int {
	factors := []int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	counts := make(map[rune]int)
	for _, r := range s1 {
		counts[r]++
	}
	for _, r := range s2 {
		counts[r]--
		if counts[r] < 0 {
			return false
		}
	}
	return true
}

func isArmstrong(n int) bool {
	sum := 0
	digits := len(fmt.Sprintf("%d", n))
	for n > 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		n /= 10
	}
	return sum == n
}

func isPerfectNumber(n int) bool {
	sum := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func isHappyNumber(n int) bool {
	seen := make(map[int]bool)
	for n != 1 && !seen[n] {
		seen[n] = true
		sum := 0
		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}
		n = sum
	}
	return n == 1
}

func isFibonacci(n int) bool {
	if n < 0 {
		return false
	}
	a, b := 0, 1
	for a < n {
		a, b = b, a+b
	}
	return a == n
}

func isPerfectSquare(n int) bool {
	if n < 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}
