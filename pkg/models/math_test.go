package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestIsAnagram_VariousInputs_010(t *testing.T) {
	assert.True(t, isAnagram("", ""))
	assert.False(t, isAnagram("a", ""))
	assert.False(t, isAnagram("abc", "abcd"))
	assert.True(t, isAnagram("listen", "silent"))
	assert.False(t, isAnagram("rat", "car"))
	assert.False(t, isAnagram("hello", "holla")) // Covers counts[r] < 0 case
	assert.True(t, isAnagram("anagram", "nagaram"))
}

// Test generated using Keploy

func TestIsHappyNumber_VariousInputs_013(t *testing.T) {
	assert.True(t, isHappyNumber(1))
	assert.True(t, isHappyNumber(19)) // 1^2 + 9^2 = 82 -> 8^2 + 2^2 = 68 -> 6^2 + 8^2 = 100 -> 1^2 = 1
	assert.True(t, isHappyNumber(7))  // 7^2=49 -> 4^2+9^2=16+81=97 -> 9^2+7^2=81+49=130 -> 1^2+3^2+0^2=1+9=10 -> 1^2+0^2=1
	assert.False(t, isHappyNumber(2)) // 2 -> 4 -> 16 -> 37 -> 58 -> 89 -> 145 -> 42 -> 20 -> 4 (cycle)
	assert.False(t, isHappyNumber(4)) // Cycle detected
	assert.False(t, isHappyNumber(0)) // 0 is not 1, seen[0]=true, sum=0, n=0, loop again, returns false
}

// Test generated using Keploy

func TestFibonacci_PositiveInteger_1111(t *testing.T) {
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

func TestIsPalindrome_VariousInputs_008(t *testing.T) {
	assert.True(t, isPalindrome(""))
	assert.True(t, isPalindrome("a"))
	assert.True(t, isPalindrome("madam"))
	assert.False(t, isPalindrome("hello"))
	assert.False(t, isPalindrome("race a car")) // Tests space and case difference
}

// Test generated using Keploy

func TestIsFibonacci_VariousInputs_014(t *testing.T) {
	assert.False(t, isFibonacci(-1))
	assert.True(t, isFibonacci(0))
	assert.True(t, isFibonacci(1))
	assert.True(t, isFibonacci(2))
	assert.True(t, isFibonacci(3))
	assert.True(t, isFibonacci(5))
	assert.True(t, isFibonacci(8))
	assert.False(t, isFibonacci(4))
	assert.False(t, isFibonacci(7))
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

func TestIsPrime_PrimeNumber_1616(t *testing.T) {
	result := isPrime(7)
	if !result {
		t.Errorf("Expected true, but got false")
	}
}

// Test generated using Keploy

func TestPrimeFactors_LessThanTwo_005(t *testing.T) {
	assert.Equal(t, []int{}, primeFactors(1))
	assert.Equal(t, []int{}, primeFactors(0))
	assert.Equal(t, []int{}, primeFactors(-10))
}

// Test generated using Keploy

func TestReverseString_VariousInputs_009(t *testing.T) {
	assert.Equal(t, "", reverseString(""))
	assert.Equal(t, "a", reverseString("a"))
	assert.Equal(t, "olleh", reverseString("hello"))
	assert.Equal(t, "madam", reverseString("madam"))
	assert.Equal(t, "🚀olleH", reverseString("Hello🚀")) // Test Unicode
}

// Test generated using Keploy

func TestIsPerfectSquare_VariousInputs_015(t *testing.T) {
	assert.False(t, isPerfectSquare(-1))
	assert.False(t, isPerfectSquare(-4))
	assert.True(t, isPerfectSquare(0))
	assert.True(t, isPerfectSquare(1))
	assert.True(t, isPerfectSquare(4))
	assert.True(t, isPerfectSquare(9))
	assert.True(t, isPerfectSquare(16))
	assert.False(t, isPerfectSquare(2))
	assert.False(t, isPerfectSquare(10))
}

func TestGCD_PositiveIntegers_1414(t *testing.T) {
	result := gcd(48, 18)
	expected := 6
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestPrimeFactors_PrimeNumber_006(t *testing.T) {
	assert.Equal(t, []int{7}, primeFactors(7))
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

func TestSquareRoot_PositiveInteger_606(t *testing.T) {
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

func TestPrimeFactors_CompositeNumber_007(t *testing.T) {
	assert.Equal(t, []int{2, 2, 3}, primeFactors(12))
}

// Test generated using Keploy

func TestIsPrime_LessThanTwo_004(t *testing.T) {
	assert.False(t, isPrime(1))
	assert.False(t, isPrime(0))
	assert.False(t, isPrime(-5))
}

// Test generated using Keploy

func TestFactorial_NegativeInteger_1010(t *testing.T) {
	_, err := factorial(-5)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

// Test generated using Keploy

func TestFibonacci_NegativeInteger_1313(t *testing.T) {
	_, err := fibonacci(-5)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
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
		t.Errorf("Expected error, but got nil")
	}
}

// Test generated using Keploy

func TestSquareRoot_NegativeInteger_707(t *testing.T) {
	_, err := squareRoot(-16)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

// Test generated using Keploy

func TestMultiply_PositiveIntegers_789(t *testing.T) {
	result := multiply(3, 5)
	expected := 15
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestModulus_DivisionByZero_404(t *testing.T) {
	_, err := modulus(10, 0)
	if err == nil {
		t.Errorf("Expected error, but got nil")
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

func TestLCM_PositiveIntegers_1515(t *testing.T) {
	result := lcm(4, 6)
	expected := 12
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestFactorial_ZeroInput_909(t *testing.T) {
	result, err := factorial(0)
	expected := 1
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsPrime_NonPrimeNumber_1717(t *testing.T) {
	result := isPrime(4)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

// Test generated using Keploy

