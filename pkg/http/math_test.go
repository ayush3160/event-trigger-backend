package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test generated using Keploy
func TestIsHappyNumber_VariousInputs_019(t *testing.T) {
	assert.True(t, isHappyNumber(1), "isHappyNumber(1)")     // Base case, loop doesn't run much
	assert.True(t, isHappyNumber(7), "isHappyNumber(7)")     // 7 -> 49 -> 97 -> 130 -> 10 -> 1
	assert.True(t, isHappyNumber(19), "isHappyNumber(19)")   // 19 -> 82 -> 68 -> 100 -> 1
	assert.True(t, isHappyNumber(100), "isHappyNumber(100)") // 100 -> 1

	assert.False(t, isHappyNumber(0), "isHappyNumber(0)") // Enters loop, sum is 0, seen[0]=true, loop terminates, n=0 != 1
	assert.False(t, isHappyNumber(2), "isHappyNumber(2)") // Enters cycle: 2 -> 4 -> 16 -> 37 -> 58 -> 89 -> 145 -> 42 -> 20 -> 4
	assert.False(t, isHappyNumber(4), "isHappyNumber(4)") // Enters cycle: 4 -> 16 -> ... -> 4 (covers seen[n] = true)
}

// Test generated using Keploy

func TestFibonacci_PositiveInteger_1010(t *testing.T) {
	result, err := fibonacci(5)
	expected := []int{0, 1, 1, 2, 3}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, but got %d", v, i, result[i])
		}
	}
}

// Test generated using Keploy

func TestIsAnagram_ValidAnagrams_1919(t *testing.T) {
	result := isAnagram("listen", "silent")
	if !result {
		t.Errorf("Expected true for valid anagrams, but got false")
	}
}

// Test generated using Keploy

func TestPrimeFactors_VariousInputs_013(t *testing.T) {
	assert.Empty(t, primeFactors(0), "primeFactors(0)")
	assert.Empty(t, primeFactors(1), "primeFactors(1)")
	assert.Equal(t, []int{2}, primeFactors(2), "primeFactors(2)") // Covers line 108
	assert.Equal(t, []int{3}, primeFactors(3), "primeFactors(3)")
	assert.Equal(t, []int{2, 2}, primeFactors(4), "primeFactors(4)") // Covers inner loop line 102-105
	assert.Equal(t, []int{2, 3}, primeFactors(6), "primeFactors(6)")
	assert.Equal(t, []int{7}, primeFactors(7), "primeFactors(7)")
	assert.Equal(t, []int{2, 2, 3}, primeFactors(12), "primeFactors(12)")
	assert.Equal(t, []int{2, 2, 5, 5}, primeFactors(100), "primeFactors(100)")
	assert.Equal(t, []int{13}, primeFactors(13), "primeFactors(13)") // Covers line 108 with a prime > sqrt(n) initially
}

// Test generated using Keploy

func TestIsFibonacci_VariousInputs_020(t *testing.T) {
	assert.False(t, isFibonacci(-1), "isFibonacci(-1)") // Covers line 188
	assert.True(t, isFibonacci(0), "isFibonacci(0)")    // a == n initially
	assert.True(t, isFibonacci(1), "isFibonacci(1)")    // First iteration: a=1, b=1. Second: a=1, b=2. Loop terminates a==n.
	assert.True(t, isFibonacci(2), "isFibonacci(2)")
	assert.True(t, isFibonacci(3), "isFibonacci(3)")
	assert.True(t, isFibonacci(5), "isFibonacci(5)")
	assert.True(t, isFibonacci(8), "isFibonacci(8)")
	assert.True(t, isFibonacci(13), "isFibonacci(13)")

	assert.False(t, isFibonacci(4), "isFibonacci(4)") // Loop: (0,1)->(1,1)->(1,2)->(2,3)->(3,5). a becomes 5 (>4), returns false.
	assert.False(t, isFibonacci(6), "isFibonacci(6)")
	assert.False(t, isFibonacci(7), "isFibonacci(7)")
	assert.False(t, isFibonacci(10), "isFibonacci(10)")
}

// Test generated using Keploy

func TestFactorial_PositiveInteger_808(t *testing.T) {
	result, err := factorial(5)
	expected := 120
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsPalindrome_PalindromeString_1616(t *testing.T) {
	result := isPalindrome("racecar")
	if !result {
		t.Errorf("Expected true for palindrome string, but got false")
	}
}

// Test generated using Keploy

func TestIsPrime_PrimeNumber_1414(t *testing.T) {
	result := isPrime(7)
	if !result {
		t.Errorf("Expected true for prime number, but got false")
	}
}

// Test generated using Keploy

func TestReverseString_ValidString_1818(t *testing.T) {
	result := reverseString("hello")
	expected := "olleh"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Test generated using Keploy

func TestIsPerfectSquare_VariousInputs_021(t *testing.T) {
	assert.False(t, isPerfectSquare(-1), "isPerfectSquare(-1)") // Covers line 199
	assert.True(t, isPerfectSquare(0), "isPerfectSquare(0)")
	assert.True(t, isPerfectSquare(1), "isPerfectSquare(1)")
	assert.True(t, isPerfectSquare(4), "isPerfectSquare(4)")
	assert.True(t, isPerfectSquare(9), "isPerfectSquare(9)")
	assert.True(t, isPerfectSquare(16), "isPerfectSquare(16)")
	assert.True(t, isPerfectSquare(25), "isPerfectSquare(25)")

	assert.False(t, isPerfectSquare(2), "isPerfectSquare(2)")
	assert.False(t, isPerfectSquare(3), "isPerfectSquare(3)")
	assert.False(t, isPerfectSquare(5), "isPerfectSquare(5)")
	assert.False(t, isPerfectSquare(8), "isPerfectSquare(8)")
	assert.False(t, isPerfectSquare(10), "isPerfectSquare(10)")
	assert.False(t, isPerfectSquare(15), "isPerfectSquare(15)")
}

func TestGCD_PositiveIntegers_1212(t *testing.T) {
	result := gcd(48, 18)
	expected := 6
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestModulus_ValidInputs_303(t *testing.T) {
	result, err := modulus(10, 3)
	expected := 1
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestSquareRoot_ValidPositiveInteger_606(t *testing.T) {
	result, err := squareRoot(16)
	expected := 4
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestDivide_ValidInputs_101(t *testing.T) {
	result, err := divide(10, 2)
	expected := 5
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestLCM_PositiveIntegers_1313(t *testing.T) {
	result := lcm(4, 6)
	expected := 12
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsPalindrome_NonPalindromeString_1717(t *testing.T) {
	result := isPalindrome("hello")
	if result {
		t.Errorf("Expected false for non-palindrome string, but got true")
	}
}

// Test generated using Keploy

func TestIsAnagram_InvalidAnagrams_2020(t *testing.T) {
	result := isAnagram("hello", "world")
	if result {
		t.Errorf("Expected false for invalid anagrams, but got true")
	}
}

// Test generated using Keploy

func TestIsPrime_EdgeCases_012(t *testing.T) {
	assert.False(t, isPrime(0), "isPrime(0)") // Covers line 88
	assert.False(t, isPrime(1), "isPrime(1)") // Covers line 88
	assert.True(t, isPrime(2), "isPrime(2)")
	assert.True(t, isPrime(3), "isPrime(3)")
	assert.False(t, isPrime(4), "isPrime(4)") // Covers line 92-93
}

// Test generated using Keploy

func TestSum_PositiveIntegers_123(t *testing.T) {
	result := sum(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestSubtract_PositiveIntegers_456(t *testing.T) {
	result := subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestDivide_DivisionByZero_202(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Errorf("Expected error for division by zero, but got nil")
	}
	expectedError := "division by zero"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
	}
}

// Test generated using Keploy

func TestSquareRoot_NegativeInteger_707(t *testing.T) {
	_, err := squareRoot(-4)
	if err == nil {
		t.Errorf("Expected error for negative input, but got nil")
	}
	expectedError := "negative number"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
	}
}

// Test generated using Keploy

func TestMultiply_PositiveIntegers_789(t *testing.T) {
	result := multiply(3, 4)
	expected := 12
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestModulus_DivisionByZero_404(t *testing.T) {
	_, err := modulus(10, 0)
	if err == nil {
		t.Errorf("Expected error for division by zero, but got nil")
	}
	expectedError := "division by zero"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
	}
}

// Test generated using Keploy

func TestPower_PositiveIntegers_505(t *testing.T) {
	result := power(2, 3)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestFibonacci_NegativeInteger_1111(t *testing.T) {
	_, err := fibonacci(-5)
	if err == nil {
		t.Errorf("Expected error for negative input, but got nil")
	}
	expectedError := "negative number"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
	}
}

// Test generated using Keploy

func TestIsPrime_NonPrimeNumber_1515(t *testing.T) {
	result := isPrime(10)
	if result {
		t.Errorf("Expected false for non-prime number, but got true")
	}
}

// Test generated using Keploy

func TestFactorial_NegativeInteger_909(t *testing.T) {
	_, err := factorial(-3)
	if err == nil {
		t.Errorf("Expected error for negative input, but got nil")
	}
	expectedError := "negative number"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
	}
}

// Test generated using Keploy

func TestFactorial_Zero_008(t *testing.T) {
	result, err := factorial(0)
	require.NoError(t, err)
	assert.Equal(t, 1, result) // Factorial of 0 is 1
}

// Test generated using Keploy

func TestIsAnagram_EdgeCasesAndCoverage_016(t *testing.T) {
	assert.True(t, isAnagram("", ""), "isAnagram('', '')")
	assert.False(t, isAnagram("a", "b"), "isAnagram('a', 'b')")             // Same length, different chars
	assert.False(t, isAnagram("a", "aa"), "isAnagram('a', 'aa')")           // Covers line 135
	assert.False(t, isAnagram("aabb", "bbba"), "isAnagram('aabb', 'bbba')") // Covers line 144 (counts['b'] < 0)
	assert.True(t, isAnagram("anagram", "nagaram"), "isAnagram('anagram', 'nagaram')")
	assert.False(t, isAnagram("rat", "car"), "isAnagram('rat', 'car')")
}

// Test generated using Keploy

